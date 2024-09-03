package core

// IflaBondSlaveEnum - Enumeration for Bonding Slave attributes
//
// This enumeration defines various attributes related to slave interfaces in a network bond.
// These attributes are used to query or set parameters for each slave interface within the bond.
//
// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
type IflaBondSlaveEnum uint8

const (
	// IflaBondSlaveUnspec - IFLA_BOND_SLAVE_UNSPEC -
	// This is an unspecified placeholder value used in the enumeration.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBondSlaveUnspec IflaBondSlaveEnum = 0

	// IflaBondSlaveState - IFLA_BOND_SLAVE_STATE -
	// This attribute represents the state of the slave interface in the bond. The state can indicate whether the
	// slave is active, inactive, or in another state within the bonding group.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBondSlaveState IflaBondSlaveEnum = 1

	// IflaBondSlaveMiiStatus - IFLA_BOND_SLAVE_MII_STATUS -
	// This attribute reports the MII (Media Independent Interface) status of the slave interface. The MII status
	// provides information about the link status and operational state of the slave.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBondSlaveMiiStatus IflaBondSlaveEnum = 2

	// IflaBondSlaveLinkFailureCount - IFLA_BOND_SLAVE_LINK_FAILURE_COUNT -
	// This attribute counts the number of link failures detected on the slave interface. It is used to track the
	// reliability and stability of the slave within the bond.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBondSlaveLinkFailureCount IflaBondSlaveEnum = 3

	// IflaBondSlavePermHwAddr - IFLA_BOND_SLAVE_PERM_HWADDR -
	// This attribute specifies the permanent hardware address (MAC address) of the slave interface. The permanent
	// MAC address is typically burned into the network interface card (NIC) and does not change.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBondSlavePermHwAddr IflaBondSlaveEnum = 4

	// IflaBondSlaveQueueId - IFLA_BOND_SLAVE_QUEUE_ID -
	// This attribute identifies the queue ID associated with the slave interface.
	// Queue IDs are used for traffic management and load balancing within the bond.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBondSlaveQueueId IflaBondSlaveEnum = 5

	// IflaBondSlaveAdAggregatorId - IFLA_BOND_SLAVE_AD_AGGREGATOR_ID -
	// This attribute represents the Aggregator ID in 802.3ad (LACP) mode for the slave interface.
	// The Aggregator ID is used to group multiple slave interfaces into a single logical link for traffic distribution.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBondSlaveAdAggregatorId IflaBondSlaveEnum = 6

	// IflaBondSlaveAdActorOperPortState - IFLA_BOND_SLAVE_AD_ACTOR_OPER_PORT_STATE -
	// This attribute reports the operational port state of the Actor in 802.3ad (LACP) mode for the slave interface.
	// The operational state provides information about how the Actor perceives the current state of the port.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBondSlaveAdActorOperPortState IflaBondSlaveEnum = 7

	// IflaBondSlaveAdPartnerOperPortState - IFLA_BOND_SLAVE_AD_PARTNER_OPER_PORT_STATE -
	// This attribute reports the operational port state of the Partner in 802.3ad (LACP) mode for the slave interface.
	// The operational state provides information about how the Partner perceives the current state of the port.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBondSlaveAdPartnerOperPortState IflaBondSlaveEnum = 8

	// IflaBondSlavePrio - IFLA_BOND_SLAVE_PRIO -
	// This attribute specifies the priority of the slave interface within the bond.
	// The priority can influence the selection of active slaves, especially in failover scenarios.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBondSlavePrio IflaBondSlaveEnum = 9

	// IflaBondSlaveMax - IFLA_BOND_SLAVE_MAX -
	// This constant represents the maximum valid value for bonding slave attributes.
	// It is used as a boundary marker to validate or limit the range of slave attributes within the bond.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBondSlaveMax = iota - 1
)
