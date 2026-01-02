package simulate

import (
	"context"

	eg "github.com/fanovilla/iam-simulate-go/internal/engine"
)

// RunSimulation validates inputs and runs the full evaluation engine.
// It returns a SimulationResult or an error when validation fails.
func RunSimulation(ctx context.Context, sim Simulation, opts ...Option) (SimulationResult, error) {
	_ = ctx
	_ = opts
	// TODO: add validation/normalization in later steps
	idp, err := eg.DecodePolicies(sim.IdentityPolicies)
	if err != nil {
		return SimulationResult{}, err
	}
	rsp, err := eg.DecodePolicies(sim.ResourcePolicies)
	if err != nil {
		return SimulationResult{}, err
	}
	scp, err := eg.DecodePolicies(sim.SCPs)
	if err != nil {
		return SimulationResult{}, err
	}
	r, a, err := eg.Evaluate(sim.Action, sim.Resource, sim.Principal, sim.Context, idp, rsp, scp)
	if err != nil {
		return SimulationResult{}, err
	}
	res := SimulationResult{Result: mapEngineResult(r), Analysis: &RequestAnalysis{}}
	_ = a // placeholder until RequestAnalysis is fully wired
	return res, nil
}

// RunUnsafeSimulation skips certain validations and runs the core evaluation.
// Intended for advanced scenarios where the caller guarantees inputs.
func RunUnsafeSimulation(ctx context.Context, sim Simulation, opts ...Option) (SimulationResult, error) {
	_ = ctx
	_ = opts
	idp, err := eg.DecodePolicies(sim.IdentityPolicies)
	if err != nil {
		return SimulationResult{}, err
	}
	rsp, err := eg.DecodePolicies(sim.ResourcePolicies)
	if err != nil {
		return SimulationResult{}, err
	}
	scp, err := eg.DecodePolicies(sim.SCPs)
	if err != nil {
		return SimulationResult{}, err
	}
	r, a, err := eg.Evaluate(sim.Action, sim.Resource, sim.Principal, sim.Context, idp, rsp, scp)
	if err != nil {
		return SimulationResult{}, err
	}
	res := SimulationResult{Result: mapEngineResult(r), Analysis: &RequestAnalysis{}}
	_ = a
	return res, nil
}

func mapEngineResult(r eg.Result) EvaluationResult {
	switch r {
	case eg.ResultAllowed:
		return EvaluationResultAllowed
	case eg.ResultExplicitlyDenied:
		return EvaluationResultExplicitlyDenied
	default:
		return EvaluationResultImplicitlyDenied
	}
}
