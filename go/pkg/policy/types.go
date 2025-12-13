package policy

import (
    "encoding/json"
    "fmt"
)

// Effect is the policy statement effect.
type Effect string

const (
    EffectAllow Effect = "Allow"
    EffectDeny  Effect = "Deny"
)

// Statement represents a single statement in a policy document.
// It intentionally models flexible encodings via StringOrSlice types.
type Statement struct {
    Sid          string        `json:"Sid,omitempty"`
    Effect       Effect        `json:"Effect"`
    Action       StringOrSlice `json:"Action,omitempty"`
    NotAction    StringOrSlice `json:"NotAction,omitempty"`
    Resource     StringOrSlice `json:"Resource,omitempty"`
    NotResource  StringOrSlice `json:"NotResource,omitempty"`
    // Principal and NotPrincipal support multiple encodings. We'll retain raw JSON
    // and parse to richer forms at evaluation time.
    Principal     json.RawMessage                         `json:"Principal,omitempty"`
    NotPrincipal  json.RawMessage                         `json:"NotPrincipal,omitempty"`
    Condition     map[string]map[string]any               `json:"Condition,omitempty"`
}

// PolicyDocument is the top-level IAM policy document.
type PolicyDocument struct {
    Version   string      `json:"Version,omitempty"`
    Statement []Statement `json:"Statement"`
}

// UnmarshalJSON supports Statement being either a single object or an array.
func (p *PolicyDocument) UnmarshalJSON(b []byte) error {
    // Alias to avoid infinite recursion
    type alias PolicyDocument
    var raw struct {
        Version   string          `json:"Version,omitempty"`
        Statement json.RawMessage `json:"Statement"`
    }
    if err := json.Unmarshal(b, &raw); err != nil {
        return err
    }
    p.Version = raw.Version
    if len(raw.Statement) == 0 {
        return fmt.Errorf("policy: missing Statement")
    }
    // Try array first
    var arr []Statement
    if err := json.Unmarshal(raw.Statement, &arr); err == nil {
        p.Statement = arr
        return nil
    }
    // Try single object
    var one Statement
    if err := json.Unmarshal(raw.Statement, &one); err == nil {
        p.Statement = []Statement{one}
        return nil
    }
    return fmt.Errorf("policy: invalid Statement encoding")
}
