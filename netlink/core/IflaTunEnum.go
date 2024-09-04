package core

// IflaTunEnum - Enumeration for TUN/TAP interface attributes.
//
// These attributes are used to configure and manage TUN/TAP interfaces in the Linux kernel.
// TUN/TAP interfaces are virtual network kernel devices used for packet tunneling.
//
// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
type IflaTunEnum uint8

const (
	// IflaTunUnspec - IFLA_TUN_UNSPEC -
	// Unspecified attribute, used as a placeholder for unknown or default values.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaTunUnspec IflaTunEnum = 0

	// IflaTunOwner - IFLA_TUN_OWNER -
	// Specifies the owner UID of the TUN/TAP interface.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaTunOwner IflaTunEnum = 1

	// IflaTunGroup - IFLA_TUN_GROUP -
	// Specifies the group ID associated with the TUN/TAP interface.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaTunGroup IflaTunEnum = 2

	// IflaTunType - IFLA_TUN_TYPE -
	// Specifies the type of the TUN/TAP interface (TUN or TAP).
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaTunType IflaTunEnum = 3

	// IflaTunPi - IFLA_TUN_PI -
	// Indicates whether packet information (PI) is included in the TUN/TAP interface.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaTunPi IflaTunEnum = 4

	// IflaTunVnetHdr - IFLA_TUN_VNET_HDR -
	// Indicates whether the TUN/TAP interface uses virtual network header support.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaTunVnetHdr IflaTunEnum = 5

	// IflaTunPersist - IFLA_TUN_PERSIST -
	// Specifies whether the TUN/TAP interface is persistent.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaTunPersist IflaTunEnum = 6

	// IflaTunMultiQueue - IFLA_TUN_MULTI_QUEUE -
	// Indicates whether the TUN/TAP interface supports multiple queues.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaTunMultiQueue IflaTunEnum = 7

	// IflaTunNumQueues - IFLA_TUN_NUM_QUEUES -
	// Specifies the number of queues for the TUN/TAP interface.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaTunNumQueues IflaTunEnum = 8

	// IflaTunNumDisabledQueues - IFLA_TUN_NUM_DISABLED_QUEUES -
	// Specifies the number of disabled queues for the TUN/TAP interface.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaTunNumDisabledQueues IflaTunEnum = 9

	// IflaTunMax - IFLA_TUN_MAX -
	// The maximum value for TUN/TAP attributes, used for validation purposes.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaTunMax = iota - 1
)
