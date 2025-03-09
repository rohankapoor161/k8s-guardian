# Release Process

## Versioning

We follow semantic versioning:
- MAJOR: Breaking changes
- MINOR: New features
- PATCH: Bug fixes

## Creating a Release

1. Update CHANGELOG.md
2. Create tag: `git tag vX.Y.Z`
3. Push: `git push origin vX.Y.Z`
4. CI will build and publish release
