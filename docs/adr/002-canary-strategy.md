# ADR 002: Canary Deployment Strategy

## Status

Accepted

## Context

Need to safely roll out changes with automatic rollback.

## Decision

Use automated canary analysis with Prometheus metrics as gate.

## Consequences

- Requires Prometheus in cluster
- Automatic rollback on threshold breach
- Slower but safer deployments
