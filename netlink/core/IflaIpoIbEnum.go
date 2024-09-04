package core

// IflaIpoIbEnum - Enumeration for IP-over-InfiniBand (IPoIB) interface attributes.
//
// These attributes are used to configure and manage IP-over-InfiniBand (IPoIB) interfaces in the Linux kernel.
// IPoIB is a network protocol that allows IP traffic to be sent over InfiniBand networks.
//
// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
type IflaIpoIbEnum uint8

const (
	// IflaIpoibUnspec - IFLA_IPOIB_UNSPEC - Unspecified attribute, used as a placeholder for unknown or default values.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaIpoibUnspec IflaIpoIbEnum = iota

	// IflaIpoibPkey - IFLA_IPOIB_PKEY - Specifies the Partition Key (PKEY) for the IPoIB interface.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaIpoibPkey

	// IflaIpoibMode - IFLA_IPOIB_MODE - Specifies the mode of the IPoIB interface (e.g., connected or datagram mode).
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaIpoibMode

	// IflaIpoibUmcast - IFLA_IPOIB_UMCAST - Indicates whether multicast support is enabled for the IPoIB interface.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaIpoibUmcast

	// IflaIpoibMax - IFLA_IPOIB_MAX - The maximum value for IPoIB attributes, used for validation purposes.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaIpoibMax = iota - 1
)
