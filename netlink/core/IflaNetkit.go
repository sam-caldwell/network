package core

// IlfaNetkit - Enumeration for NETKIT attributes in Linux kernel -
//
// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
type IlfaNetkit uint8

const (

	// IflaNetkitUnspec - IFLA_NETKIT_UNSPEC -
	// This is an unspecified placeholder value used in the enumeration.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaNetkitUnspec IlfaNetkit = iota

	// IflaNetkitPeerInfo - IFLA_NETKIT_PEER_INFO -
	// This attribute represents information about the peer in a NETKIT context.
	// It is typically used to get or set details about the peer connection or association.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaNetkitPeerInfo IlfaNetkit = iota

	// IflaNetkitPrimary - IFLA_NETKIT_PRIMARY -
	// This attribute indicates the primary setting or role within the NETKIT setup.
	// It is used to specify or query which entity or connection is considered primary.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaNetkitPrimary IlfaNetkit = iota

	// IflaNetkitPolicy - IFLA_NETKIT_POLICY -
	// This attribute represents the policy associated with NETKIT.
	// It is used to define or retrieve the policy governing NETKIT's behavior or configuration.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaNetkitPolicy IlfaNetkit = iota

	// IflaNetkitPeerPolicy - IFLA_NETKIT_PEER_POLICY -
	// This attribute represents the policy specific to the peer in a NETKIT setup.
	// It allows configuration or querying of rules and behaviors that apply to the peer entity.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaNetkitPeerPolicy IlfaNetkit = iota

	// IflaNetkitMode - IFLA_NETKIT_MODE -
	// This attribute represents the mode of operation for NETKIT.
	// The mode could define how NETKIT interfaces behave or interact under various conditions.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaNetkitMode IlfaNetkit = iota

	// IflaNetkitMax - IFLA_NETKIT_MAX -
	// This constant represents the maximum valid value for the NETKIT attributes.
	// It is often used as a boundary marker to validate or limit the range of attributes in NETKIT.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaNetkitMax = IflaNetkitMode
)
