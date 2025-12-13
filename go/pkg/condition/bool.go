package condition

import "fmt"

// Bool implements the IAM Bool operator.
type Bool struct{}

func (Bool) Name() string { return "Bool" }

func (Bool) Eval(actual Value, expected Value) (bool, error) {
    ab, okA := toBool(actual)
    eb, okE := toBool(expected)
    if !okA || !okE {
        return false, fmt.Errorf("Bool: non-bool operand")
    }
    return ab == eb, nil
}

func toBool(v any) (bool, bool) {
    switch t := v.(type) {
    case bool:
        return t, true
    case string:
        if t == "true" {
            return true, true
        }
        if t == "false" {
            return false, true
        }
        return false, false
    default:
        return false, false
    }
}
