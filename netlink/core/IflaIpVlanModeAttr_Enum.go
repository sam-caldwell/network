package core

// IflaIpVlanModeAttr - Enumeration for IP VLAN mode attributes
//
// This enumeration defines various attributes related to the configuration of IP VLANs in Linux.
// IP VLANs allow the creation of virtual network interfaces based on IP addressing.
//
// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
type IflaIpVlanModeAttr uint8

const (
	// IflaIpVlanUnspec - IFLA_IPVLAN_UNSPEC -
	// This is an unspecified placeholder value used in the enumeration.
	// It represents an undefined or unused attribute within the context of IP VLANs.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaIpVlanUnspec IflaIpVlanModeAttr = 0

	// IflaIpVlanMode - IFLA_IPVLAN_MODE -
	// This attribute specifies the mode of the IP VLAN. IP VLAN modes determine how packets are routed
	// between virtual network interfaces. Examples include L2 (Layer 2) and L3 (Layer 3) modes.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaIpVlanMode IflaIpVlanModeAttr = 1

	// IflaIpVlanFlags - IFLA_IPVLAN_FLAGS -
	// This attribute specifies additional flags for the IP VLAN configuration.
	// Flags may control various aspects of IP VLAN behavior, such as filtering or forwarding rules.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaIpVlanFlags IflaIpVlanModeAttr = 2

	// IflaIpVlanMax - IFLA_IPVLAN_MAX -
	// This constant represents the maximum valid value for IP VLAN mode attributes.
	// It is used as a boundary marker to validate or limit the range of IP VLAN attributes.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaIpVlanMax = iota - 1
)
