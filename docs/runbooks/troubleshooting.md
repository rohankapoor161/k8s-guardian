# Guardian Troubleshooting

## Common Issues

### Webhook not responding
1. Check pod status: kubectl get pods -n guardian
2. Check logs: kubectl logs -n guardian deployment/guardian
3. Verify certificate: kubectl get secret guardian-tls

### Policy not enforced
1. Check ValidatingWebhookConfiguration
2. Verify namespace labels
3. Check policy scope

### Performance issues
1. Check resource limits
2. Review cache size
3. Enable pprof profiling
