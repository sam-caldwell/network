package core

// IflaIpOverInfiniband - Enumeration for IP-over-InfiniBand (IpOIb) interface attributes.
//
// These attributes are used to configure and manage IP-over-InfiniBand (IpOIb) interfaces in the Linux kernel.
// IpOverInfiniband is a network protocol that allows IP traffic to be sent over InfiniBand networks.
//
// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
type IflaIpOverInfiniband uint8

const (
	// IflaIpOverInfinibandUnspec - IFLA_IpOverInfiniband_UNSPEC - Unspecified attribute, used as a placeholder
	// for unknown or default values.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaIpOverInfinibandUnspec IflaIpOverInfiniband = iota

	// IflaIpOverInfinibandPkey - IFLA_IpOverInfiniband_PKEY - Specifies the Partition Key (PKEY) for the
	// IpOverInfiniband interface.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaIpOverInfinibandPkey

	// IflaIpOverInfinibandMode - IFLA_IpOverInfiniband_MODE - Specifies the mode of the IpOverInfiniband
	// interface (e.g., connected or datagram mode).
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaIpOverInfinibandMode

	// IflaIpOverInfinibandUmcast - IFLA_IpOverInfiniband_UMCAST - Indicates whether multicast support is enabled
	// for the IpOverInfiniband interface.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaIpOverInfinibandUmcast

	// IflaIpOverInfinibandMax - IFLA_IpOverInfiniband_MAX - The maximum value for IpOverInfiniband attributes,
	// used for validation purposes.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaIpOverInfinibandMax = iota - 1
)
