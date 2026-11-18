# qgate

[![CI](https://github.com/kavix/qgate/actions/workflows/ci.yml/badge.svg)](https://github.com/kavix/qgate/actions/workflows/ci.yml)
[![Go Version](https://img.shields.io/github/go-mod/go-version/kavix/qgate)](https://go.dev/)
[![License: Apache 2.0](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](LICENSE)
[![OpenSSF Best Practices](https://bestpractices.coreinfrastructure.org/projects/default/badge)](https://bestpractices.coreinfrastructure.org/)

**qgate** is a crypto-agile supply chain security framework providing hybrid (ECDSA + ML-DSA) container image signing and Kubernetes admission enforcement. It allows platform engineers and security teams to smoothly transition from classical to post-quantum cryptography (PQC) without redeploying applications.

> **Status:** Early scaffold / active research. Research focus: post-quantum cryptographic agility for software supply chain integrity in GitOps-managed Kubernetes.

---

## Overview

As quantum computers approach the capability to break elliptic-curve cryptography, software supply chain signatures (e.g., Cosign, Notation) must evolve. **qgate** implements a hybrid cryptographic strategy combining NIST-standardized Post-Quantum Cryptography (FIPS 204 ML-DSA) with classical signatures (ECDSA P-256), evaluated against a flexible runtime admission policy.

```
+-----------------------------------------------------------------------------------+
|                                  CI / Pipeline                                    |
|                                                                                   |
|  Container Image  ---> [ qgate sign ] ---> Signs with ECDSA + ML-DSA keys        |
|                             |                                                     |
|                             v                                                     |
|               OCI Registry (Image + Signatures)                                   |
+-----------------------------+-----------------------------------------------------+
                              |
                              v
+-----------------------------------------------------------------------------------+
|                            Kubernetes Cluster                                     |
|                                                                                   |
|  Pod Deployment  ---> [ Validating Webhook: qgate-webhook ]                       |
|                                   |                                               |
|                                   v                                               |
|                      [ Policy Engine: policy.yaml ]                               |
|                       mode: classical | hybrid | pqc-only                         |
|                                   |                                               |
|                   +---------------+---------------+                               |
|                   |                               |                               |
|             [ Allowed ]                     [ Rejected ]                          |
+-----------------------------------------------------------------------------------+
```

---

## Features

- **Dual/Hybrid Signatures:** Produce and verify both classical (ECDSA) and post-quantum (ML-DSA) signatures for container image digests.
- **Cryptographic Agility:** Transition between migration phases (`classical` &rarr; `hybrid` &rarr; `pqc-only`) dynamically through policy without rewriting application deployments.
- **Kubernetes Native:** Admission webhook intercepts pod creation requests and enforces cryptographic policy before pods are scheduled.
- **OCI Standard Artifacts:** Signature storage adhering to OCI image and artifact specifications.

---

## Repository Layout

```
.
├── cmd/
│   ├── qgate/              # CLI for key generation, signing, and verification
│   └── webhook/            # Kubernetes validating admission webhook service
├── internal/
│   ├── signer/             # Algorithm-agnostic Signer/Verifier interfaces (ECDSA, ML-DSA)
│   └── policy/             # Policy parser and evaluation engine
├── deploy/                 # Kubernetes deployment manifests & cert-manager config
├── docs/                   # Architecture, literature review, and design specifications
├── Makefile                # Build, test, and lint automations
└── policy.example.yaml     # Sample admission policy configuration
```

---

## Getting Started

### Prerequisites

- [Go](https://go.dev/dl/) 1.23 or later
- [Docker](https://www.docker.com/) or [Podman](https://podman.io/) (optional, for container builds)
- A Kubernetes cluster (v1.28+) or [kind](https://kind.sigs.k8s.io/) for local webhook testing

### Building

Clone the repository and build the binaries using `make`:

```bash
git clone https://github.com/kavix/qgate.git
cd qgate
make build
```

This will produce:
- `bin/qgate`: The CLI tool
- `bin/webhook`: The admission webhook server

### Running Tests

```bash
make test
```

### CLI Usage

```bash
# Generate key pairs (ECDSA & ML-DSA)
./bin/qgate keygen

# Sign an image digest with hybrid keys
./bin/qgate sign --digest sha256:...

# Verify signatures against a policy
./bin/qgate verify --digest sha256:... --policy policy.example.yaml
```

---

## Policy Configuration

Admission enforcement is defined in a declarative policy file:

```yaml
mode: hybrid          # Options: classical | hybrid | pqc-only
require:
  - ecdsa-p256
  - ml-dsa-65
```

- **`classical`**: Requires classical signatures only (e.g. `ecdsa-p256`).
- **`hybrid`**: Requires both classical and post-quantum signatures to be valid.
- **`pqc-only`**: Requires post-quantum signatures (e.g. `ml-dsa-65`).

---

## Roadmap

- [ ] ML-DSA sign/verify implementation ([cloudflare/circl](https://github.com/cloudflare/circl))
- [ ] Store signatures in container registries as OCI artifacts
- [ ] Webhook signature verification engine
- [ ] GitOps / ArgoCD integration demo: signed image deploys, unverified image rejected
- [ ] Benchmarks: signature size overhead, sign/verify timings, admission latency impact

---

## Contributing

We welcome contributions! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct, development workflow, and pull request process.

---

## Security

For security vulnerability reports, please refer to our [Security Policy](SECURITY.md).

---

## License

This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details.
