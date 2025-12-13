package condition

import (
    "fmt"
    "strings"
)

// StringEqualsIgnoreCase implements case-insensitive string equality.
type StringEqualsIgnoreCase struct{}

func (StringEqualsIgnoreCase) Name() string { return "StringEqualsIgnoreCase" }

func (StringEqualsIgnoreCase) Eval(actual Value, expected Value) (bool, error) {
    as, err := toStringSlice(actual)
    if err != nil { return false, fmt.Errorf("StringEqualsIgnoreCase actual: %w", err) }
    es, err := toStringSlice(expected)
    if err != nil { return false, fmt.Errorf("StringEqualsIgnoreCase expected: %w", err) }
    for _, a := range as {
        for _, e := range es {
            if strings.EqualFold(a, e) {
                return true, nil
            }
        }
    }
    return false, nil
}

// StringNotEqualsIgnoreCase is the negation of the case-insensitive equality.
type StringNotEqualsIgnoreCase struct{}

func (StringNotEqualsIgnoreCase) Name() string { return "StringNotEqualsIgnoreCase" }

func (StringNotEqualsIgnoreCase) Eval(actual Value, expected Value) (bool, error) {
    eq := StringEqualsIgnoreCase{}
    m, err := eq.Eval(actual, expected)
    if err != nil { return false, err }
    return !m, nil
}
