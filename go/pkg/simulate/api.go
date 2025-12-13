package simulate

import (
    "context"
    "errors"
)

// RunSimulation validates inputs and runs the full evaluation engine.
// It returns a SimulationResult or an error when validation fails.
func RunSimulation(ctx context.Context, sim Simulation, opts ...Option) (SimulationResult, error) {
    _ = ctx
    _ = opts
    // Placeholder implementation; wired in later steps.
    return SimulationResult{Result: EvaluationResultImplicitlyDenied}, errors.New("not implemented: RunSimulation")
}

// RunUnsafeSimulation skips certain validations and runs the core evaluation.
// Intended for advanced scenarios where the caller guarantees inputs.
func RunUnsafeSimulation(ctx context.Context, sim Simulation, opts ...Option) (SimulationResult, error) {
    _ = ctx
    _ = opts
    // Placeholder implementation; wired in later steps.
    return SimulationResult{Result: EvaluationResultImplicitlyDenied}, errors.New("not implemented: RunUnsafeSimulation")
}
