package core

// IflaBareUdp - Enumeration for Bare IpProtoUDP interface attributes.
//
// These attributes are used to configure and manage Bare IpProtoUDP interfaces in the Linux kernel.
// Bare IpProtoUDP is a lightweight tunneling protocol that encapsulates Layer 3 protocols in IpProtoUDP.
//
// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
type IflaBareUdp uint8

const (
	// IflaBareUdpUnspec - IFLA_BAREUDP_UNSPEC -
	// Unspecified attribute, used as a placeholder for unknown or default values.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBareUdpUnspec IflaBareUdp = 0

	// IflaBareUdpPort - IFLA_BAREUDP_PORT -
	// Specifies the IpProtoUDP destination port for the Bare IpProtoUDP interface.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBareUdpPort IflaBareUdp = 1

	// IflaBareUdpEthertype - IFLA_BAREUDP_ETHERTYPE -
	// Specifies the EtherType for encapsulated packets on the Bare IpProtoUDP interface.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBareUdpEthertype IflaBareUdp = 2

	// IflaBareUdpSrcportMin - IFLA_BAREUDP_SRCPORT_MIN -
	// Specifies the minimum IpProtoUDP source port for the Bare IpProtoUDP interface.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBareUdpSrcportMin IflaBareUdp = 3

	// IflaBareUdpMultiprotoMode - IFLA_BAREUDP_MULTIPROTO_MODE -
	// Indicates whether multiple protocols are supported on the Bare IpProtoUDP interface.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBareUdpMultiprotoMode IflaBareUdp = 4

	// IflaBareUdpMax - IFLA_BAREUDP_MAX -
	// The maximum value for Bare IpProtoUDP attributes, used for validation purposes.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBareUdpMax = iota - 1
)
