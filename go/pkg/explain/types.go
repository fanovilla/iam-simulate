package explain

// StatementExplain captures matching details for a single policy statement.
type StatementExplain struct {
    Sid             string   `json:"sid,omitempty"`
    Effect          string   `json:"effect"`
    ActionMatched   bool     `json:"actionMatched"`
    PrincipalMatched bool    `json:"principalMatched"`
    ResourceMatched bool     `json:"resourceMatched"`
    ConditionMatched bool    `json:"conditionMatched"`
    MatchedActions  []string `json:"matchedActions,omitempty"`
    MatchedResources []string `json:"matchedResources,omitempty"`
}

// LayerExplain aggregates allows/denies/unmatched statements for a layer.
type LayerExplain struct {
    Allows []StatementExplain `json:"allows,omitempty"`
    Denies []StatementExplain `json:"denies,omitempty"`
    Unmatched []StatementExplain `json:"unmatched,omitempty"`
}

// AnalysisExplain is the top-level explain model across layers.
type AnalysisExplain struct {
    Identity       LayerExplain `json:"identity"`
    Resource       LayerExplain `json:"resource"`
    SCP            LayerExplain `json:"scp"`
    RCP            LayerExplain `json:"rcp"`
    PermissionBoundary LayerExplain `json:"permissionBoundary"`
    VPCEndpoint    LayerExplain `json:"vpcEndpoint"`
}
