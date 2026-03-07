# Log Analysis Runbook

## Overview

This runbook provides step-by-step procedures for analyzing Guardian webhook logs during incident response or routine troubleshooting.

## Prerequisites

- kubectl access to the guardian-system namespace
- Log aggregation system access (Splunk/Datadog/ELK)
- Basic understanding of Kubernetes admission webhooks

## Common Log Patterns

### High Error Rate Investigation

1. Check webhook pod health:
   ```bash
   kubectl get pods -n guardian-system -l app=guardian
   ```

2. Fetch recent error logs:
   ```bash
   kubectl logs -n guardian-system -l app=guardian --tail=500 | grep -i error
   ```

3. Look for specific error patterns:
   - `tls: bad certificate` - Certificate rotation needed
   - `context deadline exceeded` - Review timeout settings
   - `admission denied` - Check policy configuration

### Latency Investigation

<!-- 
FIXME(rk): Add percentile-based queries once we've standardized 
on the log format. Currently blocked on platform observability 
team's log schema review.
-->

1. Check recent request latency distribution
2. Compare against baseline metrics
3. Identify potential resource constraints

## Follow-up Actions

After completing log analysis:
- Document findings in incident tracker
- Update relevant dashboards if needed
- Escalate to SRE team if persistent patterns identified

---
*Last updated: March 2026*
