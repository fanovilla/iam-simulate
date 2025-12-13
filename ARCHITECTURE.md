# Architecture Overview

This document explains how `@cloud-copilot/iam-simulate` evaluates simulations and how the codebase is organized.

## High-level Design

The simulator determines if a principal can perform an action on a resource, considering:

- Identity policies
- Resource policies
- Service Control Policies (SCPs)
- Resource Control Policies (RCPs)
- Permission boundaries
- VPC endpoint policies

It validates inputs (actions, resources, context keys), runs the evaluation, and produces a detailed explain model.

## Public API

- `runSimulation(simulation: Simulation, options?: SimulationOptions): Promise<SimulationResult>`
- `runUnsafeSimulation(...)`
- Types and utilities exported from `src/index.ts`

## Key Modules

- `src/simulation_engine/simulationEngine.ts`: Orchestrates simulations, validates inputs, normalizes parameters, emits `SimulationResult`.
- `src/core_engine/CoreSimulatorEngine.ts`: Core IAM evaluation across layers; returns `RequestAnalysis` with per-layer analyses.
- `src/evaluate.ts`: Analysis and result types (`EvaluationResult`, `RequestAnalysis`, `IgnoredConditions`, etc.).
- `src/condition/**`: Implementations for AWS condition operators and helpers.
- `src/context_keys/**`: Context key types, discovery, and strictness rules.
- `src/explain/**`: Explain data model and optional CLI formatting.

## Data Flow

1. Build a `Simulation` (see `src/simulation_engine/simulation.ts`).
2. `runSimulation` validates actions, resources, and context keys; normalizes input.
3. `CoreSimulatorEngine` evaluates identity/resource/SCP/RCP/permission boundary/endpoint layers.
4. Each layer records matched/unmatched statements with explains (`StatementExplain`).
5. Results are merged to overall `EvaluationResult` (`Allowed`, `ExplicitlyDenied`, `ImplicitlyDenied`).

## Policy Semantics (summary)

- Explicit deny overrides allows.
- Allows must pass identity/resource checks and be within SCP/RCP/permission boundary constraints.
- RCPFullAWSAccess is implicitly present; callers need not add it.

## Extension Points

- Add condition operators under `src/condition/<family>` and register in `condition.ts`.
- Extend context key types and discovery in `src/context_keys/**`.
- Customize explain formatting via `src/explain/displayExplainCli.ts`.

## Tests

Unit tests exist for conditions, context keys, and engine behaviors (`**/*.test.ts`). Core engine fixtures live under `src/core_engine/coreEngineTests/**`.
