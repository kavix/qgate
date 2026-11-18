# Contributing to qgate

Thank you for your interest in contributing to **qgate**! We welcome community contributions, bug reports, documentation improvements, and architectural suggestions.

## Code of Conduct

All contributors and maintainers are expected to adhere to our [Code of Conduct](CODE_OF_CONDUCT.md).

## Getting Started

1. **Fork and Clone**
   Fork the repository on GitHub, then clone your fork locally:
   ```bash
   git clone https://github.com/<your-username>/qgate.git
   cd qgate
   ```

2. **Prerequisites**
   - Go 1.23+ installed.
   - `golangci-lint` (recommended for local lint checks).

3. **Verify Local Environment**
   Run the test suite and build targets:
   ```bash
   make test
   make build
   ```

## Development Workflow

1. Create a feature branch with a descriptive name:
   ```bash
   git checkout -b feature/my-new-feature
   ```

2. Make your modifications. Keep changes focused and atomic.

3. Format and lint your code:
   ```bash
   go fmt ./...
   go vet ./...
   ```

4. Ensure all unit and integration tests pass:
   ```bash
   make test
   ```

5. Write tests for any new functionality or bug fixes.

## Pull Request Guidelines

- Open a Pull Request against the `main` branch.
- Fill out the PR template describing the intent of the change and referencing any related issues.
- Ensure all automated GitHub Actions checks pass.
- Maintainers will review the PR, request changes if necessary, and merge once approved.

## Reporting Issues

- **Bugs & Feature Requests:** Please use the issue templates provided in the repository.
- **Security Vulnerabilities:** Do **not** open a public issue. Follow the instructions in [SECURITY.md](SECURITY.md).
