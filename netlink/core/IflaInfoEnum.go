package core

// IflaInfoEnum - Enumeration for interface info attributes in Linux kernel
//
// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
type IflaInfoEnum uint8

const (

	// IflaInfoUnspec - IFLA_INFO_UNSPEC - This is an unspecified placeholder value used in the enumeration.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaInfoUnspec IflaInfoEnum = 0

	// IflaInfoKind - IFLA_INFO_KIND - This attribute represents the kind of the network device.
	// For example, it can indicate if the device is a VLAN, bridge, or another type of network interface.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaInfoKind IflaInfoEnum = 1

	// IflaInfoData - IFLA_INFO_DATA - This attribute contains the data specific to the type of network interface
	// specified by IFLA_INFO_KIND. It could include various parameters that are relevant to the network device.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaInfoData IflaInfoEnum = 2

	// IflaInfoXstats - IFLA_INFO_XSTATS - This attribute is used to represent extended statistics for the network
	// interface. It provides additional statistical data beyond the basic interface statistics.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaInfoXstats IflaInfoEnum = 3

	// IflaInfoSlaveKind - IFLA_INFO_SLAVE_KIND - This attribute represents the kind of slave interface.
	// Slave interfaces are typically part of a bonding or bridging setup.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaInfoSlaveKind IflaInfoEnum = 4

	// IflaInfoSlaveData - IFLA_INFO_SLAVE_DATA - This attribute contains the data specific to the slave interface,
	// similar to how IFLA_INFO_DATA contains data for the primary interface.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaInfoSlaveData IflaInfoEnum = 5

	// IflaInfoMax - IFLA_INFO_MAX - This constant represents the maximum valid value for the IflaInfoEnum enumeration.
	// It is often used as a boundary marker for validation purposes.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaInfoMax = iota - 1
)
