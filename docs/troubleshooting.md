# Troubleshooting

## Common Issues

### Webhook certificate errors

If you see certificate errors when the webhook starts:

```bash
# Generate self-signed certs
make certs
```

**Note**: In production, use certificates from your PKI or cert-manager.

### Pod validation failing

Check that your config is loaded:

```bash
kubectl logs -n guardian-system deployment/guardian-webhook
```

Common root causes:
- ConfigMap not mounted correctly
- RBAC permissions missing for config read
- Invalid YAML syntax in config

### Resource limits not enforced

Ensure the webhook is configured with failurePolicy: Fail

## Performance Issues

### High memory usage

If the guardian webhook is consuming excessive memory:

1. Check for memory leaks in validation cache
2. Review Prometheus metrics for validation volume
3. Consider increasing resource limits temporarily

### Slow validation responses

- Verify etcd cluster health
- Check webhook timeout settings (default: 10s)
- Review validation rule complexity

---
*Last updated: 2026-03-05*
