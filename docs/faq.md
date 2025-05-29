# FAQ

## Why are my deployments failing validation?

Check that all containers have:
- Resource limits defined
- Health probes configured
- Required labels present

## Can I skip validation for specific resources?

Use the `guardian.io/skip-validation: "true"` annotation.

Performance issues?

See [Performance](performance.md) for benchmarks.
