package engine

import (
	"encoding/json"
	"fmt"

	"github.com/fanovilla/iam-simulate-go/pkg/action"
	"github.com/fanovilla/iam-simulate-go/pkg/arn"
	"github.com/fanovilla/iam-simulate-go/pkg/condition"
	"github.com/fanovilla/iam-simulate-go/pkg/policy"
)

// Result is the core engine decision enumeration.
type Result int

const (
	ResultAllowed Result = iota + 1
	ResultExplicitlyDenied
	ResultImplicitlyDenied
)

// Analysis is a placeholder for detailed per-layer analysis.
type Analysis struct{}

// Evaluate performs evaluation against identity, resource, and SCP policies.
// It matches action/resource with explicit deny precedence and evaluates
// policy conditions against the provided context.
// SCPs are evaluated as an additional layer: if SCPs don't allow the action, it's implicitly denied.
func Evaluate(actionStr, resourceStr, principal string, ctx map[string]any, idPolicies, resPolicies, scpPolicies []policy.PolicyDocument) (Result, *Analysis, error) {
	act, err := action.Parse(actionStr)
	if err != nil {
		return ResultImplicitlyDenied, nil, fmt.Errorf("parse action: %w", err)
	}
	var res arn.ARN
	if resourceStr != "*" {
		res, err = arn.Parse(resourceStr)
		if err != nil {
			return ResultImplicitlyDenied, nil, fmt.Errorf("parse resource: %w", err)
		}
	} else {
		res = arn.ARN{Partition: "*", Service: "*", Region: "*", AccountID: "*", Resource: "*"}
	}

	// Evaluate SCPs first - they act as a permission boundary
	scpRes := evalPolicies(act, res, ctx, scpPolicies)

	// Explicit deny in SCP trumps everything
	if scpRes == ResultExplicitlyDenied {
		return ResultExplicitlyDenied, &Analysis{}, nil
	}

	// If SCPs are present but don't allow, it's an implicit deny
	if len(scpPolicies) > 0 && scpRes != ResultAllowed {
		return ResultImplicitlyDenied, &Analysis{}, nil
	}

	idRes := evalPolicies(act, res, ctx, idPolicies)
	rpRes := evalPolicies(act, res, ctx, resPolicies)

	// Precedence: any explicit deny anywhere -> deny
	if idRes == ResultExplicitlyDenied || rpRes == ResultExplicitlyDenied {
		return ResultExplicitlyDenied, &Analysis{}, nil
	}
	// Allow requires identity allow AND (resource allow or no resource policies)
	idAllow := idRes == ResultAllowed
	rpAllow := rpRes == ResultAllowed || len(resPolicies) == 0
	if idAllow && rpAllow {
		return ResultAllowed, &Analysis{}, nil
	}
	return ResultImplicitlyDenied, &Analysis{}, nil
}

func evalPolicies(act action.Name, res arn.ARN, ctx map[string]any, docs []policy.PolicyDocument) Result {
	anyAllow := false
	for _, doc := range docs {
		for _, st := range doc.Statement {
			// Parse and match action
			if !matchesAction(act, st) {
				continue
			}
			// Parse and match resource
			if !matchesResource(res, st) {
				continue
			}
			// Match conditions (if any)
			if !matchesConditions(ctx, st) {
				continue
			}
			if st.Effect == policy.EffectDeny {
				return ResultExplicitlyDenied
			}
			if st.Effect == policy.EffectAllow {
				anyAllow = true
			}
		}
	}
	if anyAllow {
		return ResultAllowed
	}
	return ResultImplicitlyDenied
}

func matchesConditions(ctx map[string]any, st policy.Statement) bool {
	if len(st.Condition) == 0 {
		return true
	}
	reg := condition.Default
	for opName, conds := range st.Condition {
		op := reg.Get(opName)
		if op == nil {
			// Unknown operator -> do not match for safety
			return false
		}
		for key, expected := range conds {
			actual, ok := ctx[key]
			if !ok {
				return false
			}
			okEval, err := op.Eval(actual, expected)
			if err != nil || !okEval {
				return false
			}
		}
	}
	return true
}

func matchesAction(a action.Name, st policy.Statement) bool {
	// If NotAction present, match when parsed action does NOT match any NotAction entry.
	if len(st.NotAction.Items) > 0 {
		for _, pat := range st.NotAction.Items {
			p, err := action.Parse(pat)
			if err != nil {
				continue
			}
			if a.Matches(p) {
				return false
			}
		}
		return true
	}
	// Else use Action list
	if len(st.Action.Items) == 0 {
		// No Action means not applicable
		return false
	}
	for _, pat := range st.Action.Items {
		p, err := action.Parse(pat)
		if err != nil {
			continue
		}
		if a.Matches(p) {
			return true
		}
	}
	return false
}

func matchesResource(r arn.ARN, st policy.Statement) bool {
	// Resource can be non-ARN like "*"; handle generically
	if len(st.NotResource.Items) > 0 {
		for _, pat := range st.NotResource.Items {
			if pat == "*" {
				return false
			}
			if ra, err := arn.Parse(pat); err == nil {
				if r.Matches(ra) {
					return false
				}
			} else {
				// If not a valid ARN, do a star match against raw resource part
				if pat == r.Raw || pat == "*" {
					return false
				}
			}
		}
		return true
	}
	if len(st.Resource.Items) == 0 {
		// Absent Resource is not applicable
		return false
	}
	for _, pat := range st.Resource.Items {
		if pat == "*" {
			return true
		}
		if ra, err := arn.Parse(pat); err == nil {
			if r.Matches(ra) {
				return true
			}
		} else {
			// Non-ARN resource pattern: treat '*' as match-all only
			if pat == "*" {
				return true
			}
		}
	}
	return false
}

// DecodePolicies decodes a list of JSON policy documents.
func DecodePolicies(blobs [][]byte) ([]policy.PolicyDocument, error) {
	out := make([]policy.PolicyDocument, 0, len(blobs))
	for _, b := range blobs {
		if len(b) == 0 {
			continue
		}
		var doc policy.PolicyDocument
		if err := json.Unmarshal(b, &doc); err != nil {
			return nil, fmt.Errorf("decode policy: %w", err)
		}
		out = append(out, doc)
	}
	return out, nil
}
