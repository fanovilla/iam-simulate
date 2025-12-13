package policy

import (
    "encoding/json"
    "fmt"
)

// StringOrSlice decodes JSON values that may be a single string or an array of strings.
type StringOrSlice struct {
    Items []string
}

func (s *StringOrSlice) UnmarshalJSON(b []byte) error {
    // Try array of strings first
    var arr []string
    if err := json.Unmarshal(b, &arr); err == nil {
        s.Items = arr
        return nil
    }
    // Then try single string
    var one string
    if err := json.Unmarshal(b, &one); err == nil {
        s.Items = []string{one}
        return nil
    }
    // Finally, allow null to become empty
    var n any
    if err := json.Unmarshal(b, &n); err == nil && n == nil {
        s.Items = nil
        return nil
    }
    return fmt.Errorf("StringOrSlice: unsupported JSON: %s", string(b))
}

func (s StringOrSlice) MarshalJSON() ([]byte, error) {
    if len(s.Items) == 1 {
        return json.Marshal(s.Items[0])
    }
    return json.Marshal(s.Items)
}

// AsSlice returns the underlying slice (may be nil).
func (s StringOrSlice) AsSlice() []string { return s.Items }
