package signer

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"errors"
	"fmt"
)

const AlgorithmECDSAP256 = "ecdsa-p256"

// ECDSASigner signs digests using an ECDSA P-256 private key.
type ECDSASigner struct {
	priv *ecdsa.PrivateKey
}

// NewECDSASigner creates an ECDSASigner from an existing private key.
func NewECDSASigner(priv *ecdsa.PrivateKey) (*ECDSASigner, error) {
	if priv == nil {
		return nil, errors.New("private key cannot be nil")
	}
	return &ECDSASigner{priv: priv}, nil
}

// GenerateECDSAKey generates a fresh ECDSA P-256 private key.
func GenerateECDSAKey() (*ecdsa.PrivateKey, error) {
	return ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
}

// Algorithm returns the cryptographic algorithm identifier.
func (s *ECDSASigner) Algorithm() string {
	return AlgorithmECDSAP256
}

// Sign signs the provided digest using ECDSA ASN.1 encoding.
func (s *ECDSASigner) Sign(digest []byte) ([]byte, error) {
	if len(digest) == 0 {
		return nil, errors.New("digest cannot be empty")
	}
	sig, err := ecdsa.SignASN1(rand.Reader, s.priv, digest)
	if err != nil {
		return nil, fmt.Errorf("failed to sign digest: %w", err)
	}
	return sig, nil
}

// ECDSAVerifier verifies ECDSA P-256 signatures over digests.
type ECDSAVerifier struct {
	pub *ecdsa.PublicKey
}

// NewECDSAVerifier creates an ECDSAVerifier from a public key.
func NewECDSAVerifier(pub *ecdsa.PublicKey) (*ECDSAVerifier, error) {
	if pub == nil {
		return nil, errors.New("public key cannot be nil")
	}
	return &ECDSAVerifier{pub: pub}, nil
}

// Algorithm returns the cryptographic algorithm identifier.
func (v *ECDSAVerifier) Algorithm() string {
	return AlgorithmECDSAP256
}

// Verify verifies that signature is a valid ECDSA signature over digest.
func (v *ECDSAVerifier) Verify(digest, signature []byte) (bool, error) {
	if len(digest) == 0 {
		return false, errors.New("digest cannot be empty")
	}
	if len(signature) == 0 {
		return false, errors.New("signature cannot be empty")
	}
	valid := ecdsa.VerifyASN1(v.pub, digest, signature)
	return valid, nil
}
