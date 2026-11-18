# About qgate

## Project Summary

**qgate** (Post-Quantum Gate) is a crypto-agile software supply chain integrity framework for containerized workloads in Kubernetes.

As cryptographic standards transition toward post-quantum resilience (such as NIST FIPS 204 ML-DSA), organizations managing mission-critical Kubernetes infrastructure must adapt without causing operational disruption or vendor lock-in. **qgate** bridges this gap by enabling hybrid digital signatures (combining classical elliptic curve signatures like ECDSA with quantum-resistant lattice-based signatures like ML-DSA) and decoupling algorithm enforcement into declarative admission policies.

## Key Capabilities

1. **Hybrid Cryptographic Signing:** Signs image digests with multiple complementary algorithms in CI/CD pipelines.
2. **Crypto-Agile Admission Control:** Evaluates signatures in a Kubernetes validating webhook according to customizable policy modes:
   - `classical`: Standard elliptic-curve verification.
   - `hybrid`: Requires both classical and post-quantum signatures.
   - `pqc-only`: Fully quantum-resistant verification.
3. **Seamless Migration:** Modify cryptographic policy dynamically at runtime without recompiling images or changing workload manifests.
4. **Cloud-Native & OCI Compliant:** Built to store and verify signatures as standard OCI artifacts alongside container images.

## Architecture

```
CI / Build Phase                       Cluster Admission Phase
----------------                       -----------------------
[Image Digest]                         [Pod Admission Request]
      │                                           │
      ▼                                           ▼
[qgate sign]                           [qgate Validating Webhook]
├── ECDSA (NIST P-256)                            │
└── ML-DSA (FIPS 204)                             ▼
      │                                [Policy Evaluator]
      ▼                                Mode: hybrid / pqc-only
[OCI Registry] ───────────────────────────────────┘
```

## Maintainers

Maintained by [kavix](https://github.com/kavix).
Licensed under Apache 2.0.
