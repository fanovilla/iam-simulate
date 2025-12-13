package simulate

// EvaluationResult represents the overall decision for a request.
type EvaluationResult int

const (
    EvaluationResultAllowed EvaluationResult = iota + 1
    EvaluationResultExplicitlyDenied
    EvaluationResultImplicitlyDenied
)

// Simulation describes a single IAM authorization check to evaluate.
// It mirrors the TypeScript model while following Go idioms.
type Simulation struct {
    // Action is the AWS action name, e.g., "s3:PutObject".
    Action string `json:"action"`

    // Resource is the ARN or resource spec to evaluate.
    Resource string `json:"resource"`

    // Principal ARN or identifier for the caller.
    Principal string `json:"principal"`

    // IdentityPolicies are identity-based policy documents attached to principal.
    IdentityPolicies [][]byte `json:"identityPolicies,omitempty"`

    // ResourcePolicies are policies attached to the resource (if applicable).
    ResourcePolicies [][]byte `json:"resourcePolicies,omitempty"`

    // SCPs are Service Control Policies affecting the principal's account.
    SCPs [][]byte `json:"scps,omitempty"`

    // RCPs are Resource Control Policies for the resource's account.
    RCPs [][]byte `json:"rcps,omitempty"`

    // PermissionBoundaries attached to the principal, if any.
    PermissionBoundaries [][]byte `json:"permissionBoundaries,omitempty"`

    // VPCEndpointPolicies applied when the request flows through a VPC endpoint.
    VPCEndpointPolicies [][]byte `json:"vpcEndpointPolicies,omitempty"`

    // Context carries request context key-value pairs.
    Context map[string]any `json:"context,omitempty"`
}

// SimulationResult is the outcome of a simulation.
type SimulationResult struct {
    Result           EvaluationResult `json:"result"`
    Analysis         *RequestAnalysis `json:"analysis,omitempty"`
    Errors           []ValidationError `json:"errors,omitempty"`
    SameAccount      bool              `json:"sameAccount"`
    IgnoredConditions []string         `json:"ignoredConditions,omitempty"`
    IgnoredRoleSessionName bool        `json:"ignoredRoleSessionName,omitempty"`
}

// ValidationError captures an input validation problem.
type ValidationError struct {
    Code    string `json:"code"`
    Message string `json:"message"`
    Field   string `json:"field,omitempty"`
}

// RequestAnalysis mirrors the detailed per-layer analysis model from TS.
// Initially a placeholder; will be expanded with layer details.
type RequestAnalysis struct {
    // TODO: fill with per-layer statements/explains
}

// Option configures Simulation behavior.
type Option func(*options)

type options struct {
    discoveryMode bool
}

// WithDiscoveryMode enables discovery mode to record ignored conditions/keys.
func WithDiscoveryMode(enabled bool) Option {
    return func(o *options) { o.discoveryMode = enabled }
}
