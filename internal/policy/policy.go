// Package policy decides whether an image's signatures satisfy the
// configured mode.
package policy

type Mode string

const (
	Classical Mode = "classical"
	Hybrid    Mode = "hybrid"
	PQCOnly   Mode = "pqc-only"
)

type Policy struct {
	Mode    Mode
	Require []string // e.g. "ecdsa-p256", "ml-dsa-65"
}

// Allowed reports whether the set of algorithms that verified
// successfully satisfies the policy: every required one must be present.
func (p Policy) Allowed(verified map[string]bool) bool {
	for _, alg := range p.Require {
		if !verified[alg] {
			return false
		}
	}
	return len(p.Require) > 0
}

// TODO: load policy from YAML.
