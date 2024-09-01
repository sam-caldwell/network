//go:build linux

package core

// VfTrust - ifla_vf_trust
//
//	struct ifla_vf_trust {
//	  __u32 vf;
//	  __u32 setting;
//	};
type VfTrust struct {
	Vf      uint32
	Setting uint32
}
