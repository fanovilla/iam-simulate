package condition

import (
    "fmt"
    "strconv"
)

// Helpers
func toFloatSlice(v any) ([]float64, error) {
    switch t := v.(type) {
    case float64:
        return []float64{t}, nil
    case float32:
        return []float64{float64(t)}, nil
    case int:
        return []float64{float64(t)}, nil
    case int64:
        return []float64{float64(t)}, nil
    case int32:
        return []float64{float64(t)}, nil
    case string:
        f, err := strconv.ParseFloat(t, 64)
        if err != nil {
            return nil, fmt.Errorf("not a number: %v", v)
        }
        return []float64{f}, nil
    case []float64:
        return t, nil
    case []any:
        out := make([]float64, 0, len(t))
        for _, it := range t {
            fs, err := toFloatSlice(it)
            if err != nil {
                return nil, err
            }
            if len(fs) != 1 {
                // flatten one by one
                out = append(out, fs...)
            } else {
                out = append(out, fs[0])
            }
        }
        return out, nil
    default:
        return nil, fmt.Errorf("unsupported numeric type %T", v)
    }
}

// NumericEquals
type NumericEquals struct{}

func (NumericEquals) Name() string { return "NumericEquals" }

func (NumericEquals) Eval(actual Value, expected Value) (bool, error) {
    as, err := toFloatSlice(actual)
    if err != nil { return false, fmt.Errorf("NumericEquals actual: %w", err) }
    es, err := toFloatSlice(expected)
    if err != nil { return false, fmt.Errorf("NumericEquals expected: %w", err) }
    for _, a := range as {
        for _, e := range es {
            if a == e {
                return true, nil
            }
        }
    }
    return false, nil
}

// NumericNotEquals (negation of equals)
type NumericNotEquals struct{}

func (NumericNotEquals) Name() string { return "NumericNotEquals" }

func (NumericNotEquals) Eval(actual Value, expected Value) (bool, error) {
    eq := NumericEquals{}
    m, err := eq.Eval(actual, expected)
    if err != nil { return false, err }
    return !m, nil
}

// NumericLessThan
type NumericLessThan struct{}

func (NumericLessThan) Name() string { return "NumericLessThan" }

func (NumericLessThan) Eval(actual Value, expected Value) (bool, error) {
    as, err := toFloatSlice(actual)
    if err != nil { return false, fmt.Errorf("NumericLessThan actual: %w", err) }
    es, err := toFloatSlice(expected)
    if err != nil { return false, fmt.Errorf("NumericLessThan expected: %w", err) }
    for _, a := range as {
        for _, e := range es {
            if a < e {
                return true, nil
            }
        }
    }
    return false, nil
}

// NumericLessThanEquals
type NumericLessThanEquals struct{}

func (NumericLessThanEquals) Name() string { return "NumericLessThanEquals" }

func (NumericLessThanEquals) Eval(actual Value, expected Value) (bool, error) {
    as, err := toFloatSlice(actual)
    if err != nil { return false, fmt.Errorf("NumericLessThanEquals actual: %w", err) }
    es, err := toFloatSlice(expected)
    if err != nil { return false, fmt.Errorf("NumericLessThanEquals expected: %w", err) }
    for _, a := range as {
        for _, e := range es {
            if a <= e {
                return true, nil
            }
        }
    }
    return false, nil
}

// NumericGreaterThan
type NumericGreaterThan struct{}

func (NumericGreaterThan) Name() string { return "NumericGreaterThan" }

func (NumericGreaterThan) Eval(actual Value, expected Value) (bool, error) {
    as, err := toFloatSlice(actual)
    if err != nil { return false, fmt.Errorf("NumericGreaterThan actual: %w", err) }
    es, err := toFloatSlice(expected)
    if err != nil { return false, fmt.Errorf("NumericGreaterThan expected: %w", err) }
    for _, a := range as {
        for _, e := range es {
            if a > e {
                return true, nil
            }
        }
    }
    return false, nil
}

// NumericGreaterThanEquals
type NumericGreaterThanEquals struct{}

func (NumericGreaterThanEquals) Name() string { return "NumericGreaterThanEquals" }

func (NumericGreaterThanEquals) Eval(actual Value, expected Value) (bool, error) {
    as, err := toFloatSlice(actual)
    if err != nil { return false, fmt.Errorf("NumericGreaterThanEquals actual: %w", err) }
    es, err := toFloatSlice(expected)
    if err != nil { return false, fmt.Errorf("NumericGreaterThanEquals expected: %w", err) }
    for _, a := range as {
        for _, e := range es {
            if a >= e {
                return true, nil
            }
        }
    }
    return false, nil
}
