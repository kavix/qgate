package signer

import (
	"crypto/sha256"
	"testing"
)

func TestECDSASignAndVerify(t *testing.T) {
	priv, err := GenerateECDSAKey()
	if err != nil {
		t.Fatalf("failed to generate key: %v", err)
	}

	signer, err := NewECDSASigner(priv)
	if err != nil {
		t.Fatalf("failed to create signer: %v", err)
	}

	verifier, err := NewECDSAVerifier(&priv.PublicKey)
	if err != nil {
		t.Fatalf("failed to create verifier: %v", err)
	}

	if signer.Algorithm() != AlgorithmECDSAP256 {
		t.Errorf("expected algorithm %s, got %s", AlgorithmECDSAP256, signer.Algorithm())
	}
	if verifier.Algorithm() != AlgorithmECDSAP256 {
		t.Errorf("expected algorithm %s, got %s", AlgorithmECDSAP256, verifier.Algorithm())
	}

	payload := []byte("sha256:d87508499264c146dccb67bf91d3fd65825de756318357a731ef78f65eb5cf24")
	digest := sha256.Sum256(payload)

	sig, err := signer.Sign(digest[:])
	if err != nil {
		t.Fatalf("failed to sign digest: %v", err)
	}
	if len(sig) == 0 {
		t.Fatal("expected non-empty signature")
	}

	valid, err := verifier.Verify(digest[:], sig)
	if err != nil {
		t.Fatalf("unexpected verification error: %v", err)
	}
	if !valid {
		t.Fatal("expected signature verification to succeed")
	}

	// Tampered digest check
	tamperedDigest := sha256.Sum256([]byte("tampered"))
	validTampered, err := verifier.Verify(tamperedDigest[:], sig)
	if err != nil {
		t.Fatalf("unexpected error on tampered check: %v", err)
	}
	if validTampered {
		t.Fatal("expected signature verification to fail on tampered digest")
	}
}

func TestECDSANilHandling(t *testing.T) {
	_, err := NewECDSASigner(nil)
	if err == nil {
		t.Fatal("expected error for nil private key")
	}

	_, err = NewECDSAVerifier(nil)
	if err == nil {
		t.Fatal("expected error for nil public key")
	}
}
