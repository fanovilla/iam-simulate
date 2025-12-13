iam-simulate-go (work in progress)

This directory contains the Go port of `@cloud-copilot/iam-simulate`.

- Module: `github.com/cloud-copilot/iam-simulate-go`
- Minimum Go: 1.22+

Packages:
- pkg/simulate: Public API (`RunSimulation`, `RunUnsafeSimulation`) and core types.
- pkg/policy: AWS policy document types.
- pkg/arn: ARN parsing and matching.
- pkg/action: Action parsing and matching.
- pkg/condition: Condition operators and registry.
- pkg/contextkeys: Context key types and discovery.
- pkg/explain: Explain data structures.
- internal/engine: Core evaluation engine.
- internal/validate: Input validation and normalization.
- internal/util: Helpers.

Status: scaffolding started — API and core types are being defined.
