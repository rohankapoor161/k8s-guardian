# Metrics Design

## Key Metrics

### Admission Metrics
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
- High error rate (>1%)
- High latency (P99 >100ms)

### Warning
- Cache approaching limit
- Elevated violation rate
