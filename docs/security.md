# Security

## Principle of Least Privilege

The Guardian webhook runs with minimal required permissions:
- Read-only access to validating webhook configurations
- No cluster admin privileges
- TLS certificates for secure communication

## TLS Configuration

Webhooks require valid TLS certificates.

## RBAC

See `deploy/manifests/rbac.yaml` for the minimal RBAC configuration.
