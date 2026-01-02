package condition

import "strings"

// Value is a generic condition value (string, number, bool, list, etc.).
type Value = any

// Context provides access to request context keys.
type Context map[string]Value

// Operator evaluates a condition key against an expected value using a specific
// operator semantics (e.g., StringEquals, Bool, NumericLessThan, etc.).
type Operator interface {
	Name() string
	Eval(actual Value, expected Value) (bool, error)
}

// Registry holds known operators by canonical name.
type Registry struct {
	ops map[string]Operator
}

// NewRegistry creates an empty registry.
func NewRegistry() *Registry { return &Registry{ops: map[string]Operator{}} }

// Register adds or replaces an operator by name.
func (r *Registry) Register(op Operator) {
	// Register both canonical and lower-case names for convenience
	r.ops[op.Name()] = op
	r.ops[strings.ToLower(op.Name())] = op
}

// Get returns an operator by name, or nil if not found.
func (r *Registry) Get(name string) Operator {
	if op, ok := r.ops[name]; ok {
		return op
	}
	if op, ok := r.ops[strings.ToLower(name)]; ok {
		return op
	}
	return nil
}

// Default registry with a minimal set of operators.
var Default = func() *Registry {
	reg := NewRegistry()
	reg.Register(StringEquals{})
	reg.Register(StringNotEquals{})
	reg.Register(StringEqualsIgnoreCase{})
	reg.Register(StringNotEqualsIgnoreCase{})
	reg.Register(Bool{})
	reg.Register(NumericEquals{})
	reg.Register(NumericNotEquals{})
	reg.Register(NumericLessThan{})
	reg.Register(NumericLessThanEquals{})
	reg.Register(NumericGreaterThan{})
	reg.Register(NumericGreaterThanEquals{})
	reg.Register(ArnLike{})
	reg.Register(ArnNotLike{})
	return reg
}()
