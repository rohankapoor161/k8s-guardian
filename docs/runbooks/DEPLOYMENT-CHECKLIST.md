# Deployment Checklist

## Pre-Deployment Validation

- [ ] All unit tests pass (`make test`)
- [ ] Integration tests complete without errors
- [ ] Docker image builds successfully
- [ ] Security scan shows no critical vulnerabilities
- [ ] CHANGELOG.md updated with version notes

## Staging Deployment

- [ ] Deploy to staging cluster (`make deploy-staging`)
- [ ] Verify health endpoints respond correctly
- [ ] Run smoke tests against staging environment
- [ ] Check logs for unexpected errors
- [ ] Confirm metrics are reporting correctly

## Production Deployment

- [ ] Announce deployment in #deployments channel
- [ ] Verify canary deployment metrics (5-minute window)
- [ ] Monitor error rates for 15 minutes post-deploy
- [ ] Update deployment runbook with any issues encountered

## Post-Deployment

- [ ] Tag release in GitHub
- [ ] Update documentation if API changes introduced
- [ ] Schedule 24-hour follow-up review

---
*Last updated: 2026-03-05*
*Owner: Platform Engineering*
