# pqc-gate

Hybrid (ECDSA + ML-DSA) container image signing and Kubernetes admission
enforcement, with a policy file that switches algorithms without redeploying apps.

> Status: early scaffold. Research focus: post-quantum cryptographic agility for
> software supply chain integrity in GitOps-managed Kubernetes.

## Idea

1. **Sign** an image digest in CI with classical + post-quantum keys.
2. **Verify** at admission time with a Kubernetes validating webhook.
3. **Policy** decides what is required: `classical` -> `hybrid` -> `pqc-only`.

## Layout

| Path | Purpose |
|------|---------|
| `cmd/pqcgate` | CLI: `keygen`, `sign`, `verify` |
| `cmd/webhook` | Kubernetes validating admission webhook |
| `internal/signer` | Signer/Verifier interface, ECDSA and ML-DSA backends |
| `internal/policy` | Policy model and evaluation |
| `deploy/` | Webhook Kubernetes manifests |
| `docs/` | Literature review and design notes |

## Roadmap

- [ ] ML-DSA sign/verify (cloudflare/circl)
- [ ] Store signatures in the registry as OCI artifacts
- [ ] Webhook verifies signatures per policy
- [ ] ArgoCD demo: signed image deploys, tampered image rejected
- [ ] Benchmarks: signature size, sign/verify time, admission latency
