package core

// IfLaVfVlanInfoStruct - ifla_vf_vlan_info - Vlan information.
//
// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
type IfLaVfVlanInfoStruct struct {
	VfVlan
	// VlanProto - VLAN protocol either 802.1Q or 802.1ad
	VlanProto IfLaVfVlanInfoEnum
}
