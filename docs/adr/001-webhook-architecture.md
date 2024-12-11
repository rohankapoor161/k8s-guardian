# ADR-001: Webhook Architecture

**Status:** Accepted  
**Date:** 2024-12-10  

## Context

Need to decide between validating and mutating webhooks for Guardian.

## Decision

Use validating webhooks for policy enforcement.
Mutating webhooks only for auto-labeling.

## Consequences

- Clear separation of concerns
- Easier to reason about policy violations
- Less magic, more explicit
