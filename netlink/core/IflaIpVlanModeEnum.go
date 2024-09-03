package core

// IflaIpVlanModeEnum - Enumeration for IP VLAN mode attributes
//
// This enumeration defines various attributes related to the configuration of IP VLANs in Linux.
// IP VLANs allow the creation of virtual network interfaces based on IP addressing.
type IflaIpVlanModeEnum uint8

const (
	// IflaIpVlanUnspec - IFLA_IPVLAN_UNSPEC -
	// This is an unspecified placeholder value used in the enumeration.
	// It represents an undefined or unused attribute within the context of IP VLANs.
	IflaIpVlanUnspec IflaIpVlanModeEnum = 0

	// IflaIpVlanMode - IFLA_IPVLAN_MODE -
	// This attribute specifies the mode of the IP VLAN. IP VLAN modes determine how packets are routed
	// between virtual network interfaces. Examples include L2 (Layer 2) and L3 (Layer 3) modes.
	IflaIpVlanMode IflaIpVlanModeEnum = 1

	// IflaIpVlanFlags - IFLA_IPVLAN_FLAGS -
	// This attribute specifies additional flags for the IP VLAN configuration.
	// Flags may control various aspects of IP VLAN behavior, such as filtering or forwarding rules.
	IflaIpVlanFlags IflaIpVlanModeEnum = 2

	// IflaIpVlanMax - IFLA_IPVLAN_MAX -
	// This constant represents the maximum valid value for IP VLAN mode attributes.
	// It is used as a boundary marker to validate or limit the range of IP VLAN attributes.
	IflaIpVlanMax = iota - 1
)
