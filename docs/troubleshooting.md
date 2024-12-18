# Troubleshooting

## Common Issues

### Webhook certificate errors

If you see certificate errors when the webhook starts:

```bash
# Generate self-signed certs
make certs
```

### Pod validation failing

Check that your config is loaded:

```bash
kubectl logs -n guardian-system deployment/guardian-webhook
```

### Resource limits not enforced

Ensure the webhook is configured with failurePolicy: Fail
