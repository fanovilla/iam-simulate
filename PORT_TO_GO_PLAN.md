# Go Port Plan: iam-simulate TypeScript to Go

Progress markers: * = in progress, ✓ = completed in this session, ! = failed

1. Discovery and specification alignment ✓
   - Extract behavior from ARCHITECTURE.md and TypeScript sources (evaluate.ts, simulation.ts, condition/, context_keys/). ✓
   - Draft Behavioral Spec covering precedence rules, account scoping, validation rules, and implicit RCPFullAWSAccess. ✓
   - Identify external datasets needed to replace iam-data for action/resource metadata. *
2. Go package/layout design and API shaping ✓
   - Choose module name and packages (pkg/simulate, pkg/policy, pkg/arn, pkg/action, pkg/condition, pkg/contextkeys, pkg/explain, internal/engine, internal/validate, internal/util).
   - Define public API (RunSimulation, RunUnsafeSimulation), core types (Simulation, SimulationResult, RequestAnalysis), and option pattern.
3. Core data model ✓
   - Implement structs mirroring Simulation and evaluation types from TS.
   - Implement AWS policy document model with JSON tags and custom unmarshalling for string-or-array fields.
4. Condition operators and helpers
   - Implement StringEquals/Like, Arn*, Numeric*, Bool*, Date*, IfExists, ForAnyValue/ForAllValues.
   - Build operator registry and value typing/coercions.
5. Matching primitives *
   - ARN parsing/matching (wildcards, service-specific nuances such as S3 forms).
   - Action parsing/matching (service:Operation and wildcards).
   - Principal matching (account root, user/role ARNs, service principals, wildcard).
6. Core evaluation engine by layers
   - Identity, Resource, SCP, RCP (with implicit RCPFullAWSAccess), Permission Boundary, VPC Endpoint.
   - Apply precedence: explicit deny overrides; allow requires allows across layers with no denies.
7. Simulation engine (validation, normalization, orchestration)
   - Validate actions/resources/context keys (initially mirror TS lax global keys).
   - Discovery/strict modes via options.
   - Merge per-layer analyses into overall EvaluationResult.
8. Explain model and formatting
   - Mirror TS explain structures; optional CLI formatter.
9. Tests and fixtures
   - Unit tests for operators/matching and layer evaluation.
   - Golden fixtures: TS Simulation JSON -> expected RequestAnalysis JSON; assert Go matches.
   - Fuzz/property tests for ARN/conditions where feasible.
10. Performance and benchmarking
    - Micro-benchmarks and end-to-end scenarios; optimize hot paths.
11. Documentation and examples
    - README with Go usage; ported architecture guide; Godoc comments.
12. CI, versioning, release
    - Go 1.22+, vet/staticcheck/golangci-lint/tests/race; coverage targets; GitHub Actions; semantic versioning starting v0.1.0.

## Progress Log

- 2025-12-13: Created initial plan and began discovery and spec alignment. ✓
- 2025-12-13: Drafted Behavioral Spec and recorded open decisions. ✓
 - 2025-12-13: Initialized Go module and public API skeleton (RunSimulation, types). ✓
 - 2025-12-13: Implemented core policy document types and JSON helpers. ✓
 - 2025-12-13: Began scaffolding matching primitives (ARN and action). *

## Artifacts

- Behavioral Spec (draft): included inline below

---

## Behavioral Spec (Draft)

Core Concepts
- EvaluationResult values: Allowed, ExplicitlyDenied, ImplicitlyDenied
- ResourceEvaluationResult values: NotApplicable, Allowed, ExplicitlyDenied, AllowedForAccount, DeniedForAccount, ImplicitlyDenied
- RequestAnalysis aggregates per-layer analyses and overall result; includes sameAccount, ignoredConditions, ignoredRoleSessionName

Layers and Precedence
- Layers: Identity, Resource, SCPs, RCPs, Permission Boundaries, VPC Endpoint
- Precedence:
  - Any explicit Deny anywhere -> ExplicitlyDenied
  - Allow requires identity allow AND resource allow (if applicable) AND no blocking SCP/RCP/Boundary/Endpoint
  - Otherwise -> ImplicitlyDenied

Same vs Cross Account
- sameAccount is true when principal.accountId equals resource.accountId
- Some resource policy principal matching differs by same vs cross account

Validation and Normalization
- Validate action name exists (via action metadata provider)
- Validate resource ARN shape is compatible with the action
- Validate context keys (initially allow all global keys to mirror current TS behavior)
- Normalize inputs: arrays vs singletons, canonicalize ARNs/action casing where applicable

Condition Operators
- Support: String*, Arn*, Numeric*, Bool*, Date*, IfExists, ForAnyValue, ForAllValues
- Respect wildcard semantics and case sensitivity rules per family

Matching Primitives
- Action matching: exact, service:*, *
- ARN matching: partition/region/account/service/resource; service specific shapes (e.g., S3)
- Principal matching: wildcard, AWS account root, IAM users/roles, service principals

Explain Model
- Per statement: fields for effect, identifier, matches, actionMatch, principalMatch, resourceMatch, conditionMatch; lists for matched resources/actions
- Per layer: allow/deny/unmatched statements

Discovery Mode
- Record ignored conditions and ignoredRoleSessionName similarly to TS

Implicit Policies
- RCPFullAWSAccess is implicitly present and always applied

Errors
- SimulationResult.Errors holds structured validation errors; analysis may be absent when errors exist

Options (Go)
- Discovery mode on/off; context key strictness; trust policy customization; action metadata provider override

Data Sources
- Default action/resource metadata: vendored JSON snapshot compatible with iam-data concepts; pluggable provider interface
