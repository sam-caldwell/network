package core

// VfVlan - Vlan information
//
//	struct ifla_vf_vlan {
//	  __u32 vf;
//	  __u32 vlan; /* 0 - 4095, 0 disables VLAN filter */
//	  __u32 qos;
//	};
type VfVlan struct {
	Vf   uint32
	Vlan uint32 // 0 - 4095, 0 disables VLAN filter
	Qos  uint32
}
