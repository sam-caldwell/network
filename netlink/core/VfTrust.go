//go:build linux

package core

// VfTrust - ifla_vf_trust
type VfTrust struct {
	Vf      uint32
	Setting uint32
}
