# Metrics Design

## Overview

This document outlines the observability metrics for the Guardian admission controller. Metrics are exposed via Prometheus-compatible endpoints and should be monitored via the platform-standard Grafana dashboards.

## Key Metrics

### Admission Metrics

<!-- 
TODO(rk): Consider adding admission_webhook_latency histogram for more granular 
latency distribution analysis. See issue #78 for context.
-->

- guardian_admission_requests_total
- guardian_admission_duration_seconds
- guardian_admission_errors_total

### Policy Metrics
- guardian_policy_violations_total
- guardian_policy_evaluations_total

### Resource Metrics
- guardian_cache_size
- guardian_memory_usage_bytes

## Alerts

### Critical

<!-- CRITICAL: These thresholds are derived from SLO requirements. 
Any changes require approval from the platform SRE team. --

- High error rate (>1%)
- High latency (P99 >100ms)

### Warning
- Cache approaching limit
- Elevated violation rate
