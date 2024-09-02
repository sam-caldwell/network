package core

// VfVlanInfo - ifla_vf_vlan_info - Vlan information.
//
// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
//
//	struct ifla_vf_vlan_info {
//	  __u32 vf;
//	  __u32 vlan; /* 0 - 4095, 0 disables VLAN filter */
//	  __u32 qos;
//	  __be16 vlan_proto; /* VLAN protocol either 802.1Q or 802.1ad */
//	};
type VfVlanInfo struct {
	VfVlan

	// VlanProto - VLAN protocol either 802.1Q or 802.1ad
	VlanProto uint16
}
