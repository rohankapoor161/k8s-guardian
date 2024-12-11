# Deployment Guide

## Prerequisites

- Kubernetes cluster 1.25+
- kubectl configured
- Helm 3.x

## Installation

### Option 1: Direct Manifests

```bash
kubectl apply -f deploy/manifests/
```

### Option 2: Helm Chart

```bash
helm install guardian ./deploy/helm
```

## Configuration

Create a ConfigMap with your policies:

```bash
kubectl create configmap guardian-config \
  --from-file=config.yaml
```
