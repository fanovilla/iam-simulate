package condition

// StringNotEquals is the negation of StringEquals.
type StringNotEquals struct{}

func (StringNotEquals) Name() string { return "StringNotEquals" }

func (StringNotEquals) Eval(actual Value, expected Value) (bool, error) {
    eq := StringEquals{}
    m, err := eq.Eval(actual, expected)
    if err != nil { return false, err }
    return !m, nil
}
