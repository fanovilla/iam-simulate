package condition

import (
	"fmt"
)

// ArnLike implements ARN pattern matching with wildcards.
type ArnLike struct{}

func (ArnLike) Name() string { return "ArnLike" }

func (ArnLike) Eval(actual Value, expected Value) (bool, error) {
	// Convert actual value to string
	actualStr, ok := actual.(string)
	if !ok {
		return false, fmt.Errorf("ArnLike: actual value must be string, got %T", actual)
	}

	// Expected can be a string or []string
	switch exp := expected.(type) {
	case string:
		return arnMatches(actualStr, exp), nil
	case []any:
		for _, e := range exp {
			if eStr, ok := e.(string); ok {
				if arnMatches(actualStr, eStr) {
					return true, nil
				}
			}
		}
		return false, nil
	case []string:
		for _, pattern := range exp {
			if arnMatches(actualStr, pattern) {
				return true, nil
			}
		}
		return false, nil
	default:
		return false, fmt.Errorf("ArnLike: expected must be string or []string, got %T", expected)
	}
}

// ArnNotLike implements negated ARN pattern matching.
type ArnNotLike struct{}

func (ArnNotLike) Name() string { return "ArnNotLike" }

func (ArnNotLike) Eval(actual Value, expected Value) (bool, error) {
	result, err := ArnLike{}.Eval(actual, expected)
	if err != nil {
		return false, err
	}
	return !result, nil
}

// arnMatches checks if an ARN matches a pattern with wildcards (* and ?).
// This is a simplified version - AWS has more complex ARN matching rules.
func arnMatches(arn, pattern string) bool {
	if pattern == "*" {
		return true
	}
	if pattern == arn {
		return true
	}

	// Convert pattern to a simple wildcard matcher
	// * matches any sequence of characters
	// ? matches a single character
	return wildcardMatch(arn, pattern)
}

// wildcardMatch performs simple wildcard matching.
func wildcardMatch(s, pattern string) bool {
	// Simple implementation - can be enhanced for better performance
	sIdx, pIdx := 0, 0
	sLen, pLen := len(s), len(pattern)
	starIdx, matchIdx := -1, 0

	for sIdx < sLen {
		if pIdx < pLen {
			if pattern[pIdx] == '*' {
				starIdx = pIdx
				matchIdx = sIdx
				pIdx++
				continue
			}
			if pattern[pIdx] == '?' || pattern[pIdx] == s[sIdx] {
				sIdx++
				pIdx++
				continue
			}
		}

		// No match, backtrack to last *
		if starIdx != -1 {
			pIdx = starIdx + 1
			matchIdx++
			sIdx = matchIdx
			continue
		}

		return false
	}

	// Skip remaining * in pattern
	for pIdx < pLen && pattern[pIdx] == '*' {
		pIdx++
	}

	return pIdx == pLen
}
