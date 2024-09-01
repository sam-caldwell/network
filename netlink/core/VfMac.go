//go:build linux

package core

// VfMac - go port of c ifla_vf_mac struct
//
//	struct ifla_vf_mac {
//	  __u32 vf;
//	  __u8 mac[32]; /* MAX_ADDR_LEN */
//	};
type VfMac struct {
	Vf  uint32
	Mac [32]byte
}
