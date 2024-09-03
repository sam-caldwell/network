package core

// IflaMacVlanEnum -  MACVLAN section -
//
// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
type IflaMacVlanEnum uint8

const (
	// IflaMacvlanUnspec - IFLA_MACVLAN_UNSPEC - This is used as a placeholder in the enumeration. It indicates an
	// unspecified or undefined state.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaMacvlanUnspec IflaMacVlanEnum = 0

	// IflaMacvlanMode - IFLA_MACVLAN_MODE - This represents the mode of the MACVLAN interface. MACVLAN has several
	// modes such as `private`, `vepa`, `bridge`, etc., which determine how the MACVLAN interface interacts with other
	// network interfaces.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaMacvlanMode IflaMacVlanEnum = 1

	// IflaMacvlanFlags - IFLA_MACVLAN_FLAGS - This represents various flags associated with the MACVLAN interface.
	// Flags can be used to control certain behaviors of the MACVLAN interface, such as whether it should forward
	// broadcast traffic or handle multicast in a specific way.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaMacvlanFlags IflaMacVlanEnum = 2

	// IflaMacvlanMacaddrMode - IFLA_MACVLAN_MACADDR_MODE - This identifier is used to specify the MAC address
	// assignment mode for the MACVLAN interface. It could determine whether the MAC addresses are generated randomly,
	// set manually, or assigned by the system.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaMacvlanMacaddrMode IflaMacVlanEnum = 3

	// IflaMacvlanMacaddr - IFLA_MACVLAN_MACADDR - This represents the actual MAC address of the MACVLAN interface.
	// It is used to set or retrieve the MAC address.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaMacvlanMacaddr IflaMacVlanEnum = 4

	// IflaMacvlanMacaddrData - IFLA_MACVLAN_MACADDR_DATA - This is used to provide additional data related to the
	// MAC address, such as a list of multiple MAC addresses if the interface supports multiple addresses.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaMacvlanMacaddrData IflaMacVlanEnum = 5

	// IflaMacvlanMacaddrCount - IFLA_MACVLAN_MACADDR_COUNT - This represents the count of MAC addresses associated
	// with the MACVLAN interface. It can be used to query or set the number of MAC addresses that the interface
	// can use.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaMacvlanMacaddrCount IflaMacVlanEnum = 6

	// IflaMacvlanBcQueueLen - IFLA_MACVLAN_BC_QUEUE_LEN - This represents the length of the broadcast queue for
	// the MACVLAN interface. It is used to set or get the number of broadcast packets that can be queued before
	// being dropped.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaMacvlanBcQueueLen IflaMacVlanEnum = 7

	// IflaMacvlanBcQueueLenUsed - IFLA_MACVLAN_BC_QUEUE_LEN_USED - This represents the number of entries currently
	// used in the broadcast queue of the MACVLAN interface. It can be queried to understand the current load on the
	// broadcast queue.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaMacvlanBcQueueLenUsed IflaMacVlanEnum = 8

	// IflaMacvlanBcCutoff - IFLA_MACVLAN_BC_CUTOFF - This defines a cutoff threshold for broadcast traffic on the
	// MACVLAN interface. It is used to limit the amount of broadcast traffic processed by the interface.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaMacvlanBcCutoff IflaMacVlanEnum = 9

	// IflaMacvlanMax - IFLA_MACVLAN_MAX - This defines the maximum valid value for MACVLAN attributes. It is often
	// used to validate the range of values in netlink communications.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaMacvlanMax IflaMacVlanEnum = 9
)
