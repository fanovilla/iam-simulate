package condition

import (
    "fmt"
)

// StringEquals implements the IAM StringEquals operator.
// It performs case-sensitive equality. If either side is a list, it matches
// if any pair equals per IAM semantics.
type StringEquals struct{}

func (StringEquals) Name() string { return "StringEquals" }

func (StringEquals) Eval(actual Value, expected Value) (bool, error) {
    // Normalize to slices
    as, err := toStringSlice(actual)
    if err != nil { return false, fmt.Errorf("StringEquals actual: %w", err) }
    es, err := toStringSlice(expected)
    if err != nil { return false, fmt.Errorf("StringEquals expected: %w", err) }
    for _, a := range as {
        for _, e := range es {
            if a == e {
                return true, nil
            }
        }
    }
    return false, nil
}

func toStringSlice(v any) ([]string, error) {
    switch t := v.(type) {
    case string:
        return []string{t}, nil
    case []string:
        return t, nil
    case []any:
        out := make([]string, 0, len(t))
        for _, it := range t {
            s, ok := it.(string)
            if !ok { return nil, fmt.Errorf("non-string in slice") }
            out = append(out, s)
        }
        return out, nil
    case nil:
        return nil, nil
    default:
        return nil, fmt.Errorf("unsupported type %T", v)
    }
}
