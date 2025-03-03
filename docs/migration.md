# Migration Guide

## From v0.0.x to v0.1.0

Breaking changes in configuration:
- `gates.resources.max_memory` is now `gates.resources.maxMemoryLimit`
- Add `tier` label to skip PDB requirements

## Upgrading

1. Update config file
2. Update CRDs
3. Validate existing resources
