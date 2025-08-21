# Integrations

## Prometheus

Configure metrics endpoint:
```yaml
prometheus:
  url: http://prometheus:9090
```

## Slack

Add webhook URL for notifications:
```yaml
slack:
  webhook: https://hooks.slack.com/...
```

## PagerDuty

```yaml
pagerduty:
    service_key: ${PD_KEY}
```
