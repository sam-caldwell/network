package core

// IflaIpVlan - Enumeration for IP VLAN attributes
//
// This enumeration defines various attributes related to IP VLANs in Linux.
// IP VLANs allow the configuration of virtual LANs based on IP addressing.
//
// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h

type IflaIpVlanEnum uint8

const (

	// IflaVfInfoUnspec - IFLA_VF_INFO_UNSPEC -
	// This is an unspecified placeholder value used in the enumeration.
	// It represents an undefined or unused attribute within the context of IP VLANs.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaVfInfoUnspec IflaIpVlanEnum = 0

	// IflaVfInfo - IFLA_VF_INFO -
	// This attribute represents information related to a Virtual Function (VF) in the context of IP VLANs.
	// It is used to query or configure specific VF-related settings for IP VLANs.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaVfInfo IflaIpVlanEnum = 1

	// IflaVfInfoMax - IFLA_VF_INFO_MAX -
	// This constant represents the maximum valid value for IP VLAN attributes.
	// It is used as a boundary marker to validate or limit the range of IP VLAN attributes.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaVfInfoMax = iota - 1
)
