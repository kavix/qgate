package policy

import "testing"

func TestHybridRequiresBoth(t *testing.T) {
	p := Policy{Mode: Hybrid, Require: []string{"ecdsa-p256", "ml-dsa-65"}}

	if p.Allowed(map[string]bool{"ecdsa-p256": true}) {
		t.Fatal("hybrid must reject when ML-DSA is missing")
	}
	if !p.Allowed(map[string]bool{"ecdsa-p256": true, "ml-dsa-65": true}) {
		t.Fatal("hybrid must accept when both verify")
	}
}
