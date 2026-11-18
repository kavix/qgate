// Package signer defines the algorithm-agnostic signing interface.
// Each algorithm (ECDSA, ML-DSA, ...) is a backend behind it, which is
// what makes the system crypto-agile.
package signer

// Signer signs an image digest such as "sha256:abc123...".
type Signer interface {
	Algorithm() string
	Sign(digest []byte) ([]byte, error)
}

// Verifier checks a signature over an image digest.
type Verifier interface {
	Algorithm() string
	Verify(digest, signature []byte) (bool, error)
}

// TODO: ecdsa.go (stdlib crypto/ecdsa), mldsa.go (cloudflare/circl).
