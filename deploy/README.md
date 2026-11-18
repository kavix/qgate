# Kubernetes Deployment Manifests

This directory contains the production-ready manifests to deploy the **qgate** validating admission webhook into a Kubernetes cluster.

## Manifest Inventory

| File | Resource | Purpose |
| ---- | -------- | ------- |
| `namespace.yaml` | Namespace | Creates the dedicated `qgate-system` namespace |
| `rbac.yaml` | ServiceAccount, ClusterRole, ClusterRoleBinding | Grants necessary permissions to the webhook service |
| `deployment.yaml` | Deployment | Runs the webhook pod replicas with non-root security context |
| `service.yaml` | Service | Exposes the HTTPS port (443 -> 8443) internally |
| `webhook-configuration.yaml` | ValidatingWebhookConfiguration | Intercepts pod admission requests for cryptographic verification |

## Quickstart

### 1. Create TLS Certificates

The webhook requires TLS. We recommend using [cert-manager](https://cert-manager.io/):

```bash
kubectl apply -f deploy/namespace.yaml
```

Generate the secret `qgate-webhook-tls` in `qgate-system` containing `tls.crt` and `tls.key`.

### 2. Create the Policy ConfigMap

```bash
kubectl create configmap qgate-policy \
  --namespace qgate-system \
  --from-file=policy.yaml=policy.example.yaml
```

### 3. Deploy the Webhook

```bash
kubectl apply -f deploy/rbac.yaml
kubectl apply -f deploy/deployment.yaml
kubectl apply -f deploy/service.yaml
kubectl apply -f deploy/webhook-configuration.yaml
```

### 4. Verify Status

```bash
kubectl rollout status deployment/qgate-webhook -n qgate-system
```
