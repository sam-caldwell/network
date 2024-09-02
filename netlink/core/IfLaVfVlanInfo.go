package core

// IfLaVfVlanInfoStruct - ifla_vf_vlan_info - Vlan information.
//
// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
type IfLaVfVlanInfoStruct struct {
	VfVlan
	// VlanProto - VLAN protocol either 802.1Q or 802.1ad
	VlanProto uint16
}

const (
	// IfLaVfVlanInfoUnspec - IFLA_VF_VLAN_INFO_UNSPEC - constant in the Linux kernel's networking subsystem that
	// serves as a placeholder or default value, indicating that no specific VLAN information is specified for a
	// Virtual Function (VF). It is used in contexts where VLAN information for a VF is either not provided or is
	// irrelevant, typically within the SR-IOV (Single Root I/O Virtualization) framework.
	IfLaVfVlanInfoUnspec = iota

	// IfLaVfVlanInfo - IFLA_VF_VLAN_INFO - VLAN ID, QoS and VLAN protocol. Constant in the Linux kernel's networking
	// subsystem used to represent VLAN (Virtual Local Area Network) information for a Virtual Function (VF)
	// associated with a network interface. This constant is used in the context of configuring or querying VLAN
	// settings for VFs, which are part of SR-IOV (Single Root I/O Virtualization) enabled network interfaces.
	IfLaVfVlanInfo

	// IfLaVfVlanInfoMax - IFLA_VF_VLAN_INFO_MAX - constant in the Linux kernel's networking subsystem that represents
	// the maximum value or the upper limit of attributes related to VLAN information for a Virtual Function (VF) in
	// the context of SR-IOV (Single Root I/O Virtualization). It is used to validate or iterate over the set of
	// VLAN-related attributes associated with a VF, ensuring that the configuration or querying of these attributes
	// stays within the defined limits.
	IfLaVfVlanInfoMax
)
