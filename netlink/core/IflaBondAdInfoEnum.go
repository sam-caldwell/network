package core

// IflaBondAdInfoEnum - Enumeration for 802.3ad bonding information attributes
//
// This enumeration defines attributes related to the 802.3ad (LACP) bonding information,
// used to query or set parameters for link aggregation in Linux.
//
// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
type IflaBondAdInfoEnum uint8

const (

	// IflaBondAdInfoUnspec - IFLA_BOND_AD_INFO_UNSPEC -
	// This is an unspecified placeholder value used in the enumeration.
	// It represents an undefined or unused attribute within the context of 802.3ad bonding information.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBondAdInfoUnspec IflaBondAdInfoEnum = 0

	// IflaBondAdInfoAggregator - IFLA_BOND_AD_INFO_AGGREGATOR -
	// This attribute identifies the aggregator that a particular bonding interface is using.
	// In the context of LACP, an aggregator groups multiple physical links into a single logical link.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBondAdInfoAggregator IflaBondAdInfoEnum = 1

	// IflaBondAdInfoNumPorts - IFLA_BOND_AD_INFO_NUM_PORTS -
	// This attribute specifies the number of ports that are currently active in the LACP bond.
	// It reflects the number of interfaces that are part of the aggregated link.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBondAdInfoNumPorts IflaBondAdInfoEnum = 2

	// IflaBondAdInfoActorKey - IFLA_BOND_AD_INFO_ACTOR_KEY -
	// This attribute represents the key assigned to the bond's actor (the local system) in LACP.
	// The actor key is used to identify the aggregated link from the perspective of the local system.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBondAdInfoActorKey IflaBondAdInfoEnum = 3

	// IflaBondAdInfoPartnerKey - IFLA_BOND_AD_INFO_PARTNER_KEY -
	// This attribute represents the key assigned to the bond's partner (the remote system) in LACP.
	// The partner key is used to identify the aggregated link from the perspective of the remote system.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBondAdInfoPartnerKey IflaBondAdInfoEnum = 4

	// IflaBondAdInfoPartnerMac - IFLA_BOND_AD_INFO_PARTNER_MAC -
	// This attribute specifies the MAC address of the partner (the remote system) in the LACP bond.
	// It is used to identify the remote system that is participating in the link aggregation.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBondAdInfoPartnerMac IflaBondAdInfoEnum = 5

	// IflaBondAdInfoMax - IFLA_BOND_AD_INFO_MAX -
	// This constant represents the maximum valid value for 802.3ad bonding information attributes.
	// It is used as a boundary marker to validate or limit the range of these attributes.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBondAdInfoMax IflaBondAdInfoEnum = iota - 1
)
