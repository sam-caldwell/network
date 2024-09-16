package core

// IflaBrEnum - Enumeration for bridge interface attributes in Linux.
//
// These attributes correspond to the settings and state information for Linux bridge interfaces.
// They are used to configure and retrieve bridge parameters via Netlink.
//
// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
type IflaBrEnum uint8

const (
	// IflaBrUnspec - IFLA_BR_UNSPEC - Unspecified attribute, used as a placeholder.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBrUnspec IflaBrEnum = iota

	// IflaBrForwardDelay - IFLA_BR_FORWARD_DELAY - The time the bridge will stay in the listening and learning
	// states before forwarding traffic.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBrForwardDelay

	// IflaBrHelloTime - IFLA_BR_HELLO_TIME - The time between sending bridge protocol data units (BPDU) on a
	// bridge interface.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBrHelloTime

	// IflaBrMaxAge - IFLA_BR_MAX_AGE - The maximum time a BPDU is kept in the bridge before it is discarded.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBrMaxAge

	// IflaBrAgeingTime - IFLA_BR_AGEING_TIME - The time after which a MAC address will be removed from the
	// bridge's forwarding table.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBrAgeingTime

	// IflaBrStpState - IFLA_BR_STP_STATE - The state of the Spanning Tree Protocol (STP) on the bridge.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBrStpState

	// IflaBrPriority - IFLA_BR_PRIORITY - The priority of the bridge in the Spanning Tree Protocol.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBrPriority

	// IflaBrVlanFiltering - IFLA_BR_VLAN_FILTERING - Indicates whether VLAN filtering is enabled on the bridge.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBrVlanFiltering

	// IflaBrVlanProtocol - IFLA_BR_VLAN_PROTOCOL - The protocol used for VLAN tagging on the bridge.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBrVlanProtocol

	// IflaBrGroupFwdMask - IFLA_BR_GROUP_FWD_MASK - A mask specifying which multicast groups are forwarded by the
	// bridge.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBrGroupFwdMask

	// IflaBrRootId - IFLA_BR_ROOT_ID - The ID of the root bridge in the Spanning Tree Protocol.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBrRootId

	// IflaBrBridgeId - IFLA_BR_BRIDGE_ID - The ID of the bridge itself in the Spanning Tree Protocol.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBrBridgeId

	// IflaBrRootPort - IFLA_BR_ROOT_PORT - The port on the bridge that is the root port in the Spanning Tree Protocol.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBrRootPort

	// IflaBrRootPathCost - IFLA_BR_ROOT_PATH_COST - The cost of the path to the root bridge in the Spanning
	// Tree Protocol.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBrRootPathCost

	// IflaBrTopologyChange - IFLA_BR_TOPOLOGY_CHANGE - Indicates if a topology change has occurred in the
	// Spanning Tree Protocol.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBrTopologyChange

	// IflaBrTopologyChangeDetected - IFLA_BR_TOPOLOGY_CHANGE_DETECTED - Indicates if a topology change has been
	// detected on the bridge.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBrTopologyChangeDetected

	// IflaBrHelloTimer - IFLA_BR_HELLO_TIMER - The timer for sending hello packets in the Spanning Tree Protocol.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBrHelloTimer

	// IflaBrTcnTimer - IFLA_BR_TCN_TIMER - The timer for the topology change notification (TCN) in the Spanning
	// Tree Protocol.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBrTcnTimer

	// IflaBrTopologyChangeTimer - IFLA_BR_TOPOLOGY_CHANGE_TIMER - The timer that runs when a topology change has
	// occurred.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBrTopologyChangeTimer

	// IflaBrGcTimer - IFLA_BR_GC_TIMER - The garbage collection timer for the bridge's forwarding table.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBrGcTimer

	// IflaBrGroupAddr - IFLA_BR_GROUP_ADDR - The multicast group MAC address used by the bridge.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBrGroupAddr

	// IflaBrFdbFlush - IFLA_BR_FDB_FLUSH - Flushes the forwarding database (FDB) of the bridge.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBrFdbFlush

	// IflaBrMcastRouter - IFLA_BR_MCAST_ROUTER - The multicast router setting for the bridge.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBrMcastRouter

	// IflaBrMcastSnooping - IFLA_BR_MCAST_SNOOPING - Indicates if multicast snooping is enabled on the bridge.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBrMcastSnooping

	// IflaBrMcastQueryUseIfaddr - IFLA_BR_MCAST_QUERY_USE_IFADDR - Indicates if the bridge should use its own IP
	// address when sending IGMP/MLD queries.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBrMcastQueryUseIfaddr

	// IflaBrMcastQuerier - IFLA_BR_MCAST_QUERIER - Indicates if the bridge should act as a multicast querier.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBrMcastQuerier

	// IflaBrMcastHashElasticity - IFLA_BR_MCAST_HASH_ELASTICITY - The multicast hash elasticity setting for the
	// bridge.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBrMcastHashElasticity

	// IflaBrMcastHashMax - IFLA_BR_MCAST_HASH_MAX - The maximum number of hash buckets for multicast entries.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBrMcastHashMax

	// IflaBrMcastLastMemberCnt - IFLA_BR_MCAST_LAST_MEMBER_CNT - The number of last-member queries sent before the
	// bridge removes a multicast group.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBrMcastLastMemberCnt

	// IflaBrMcastStartupQueryCnt - IFLA_BR_MCAST_STARTUP_QUERY_CNT - The number of queries sent during the startup
	// phase of IGMP/MLD querying.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBrMcastStartupQueryCnt

	// IflaBrMcastLastMemberIntvl - IFLA_BR_MCAST_LAST_MEMBER_INTVL - The interval between last-member queries.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBrMcastLastMemberIntvl

	// IflaBrMcastMembershipIntvl - IFLA_BR_MCAST_MEMBERSHIP_INTVL - The interval during which a multicast
	// membership remains valid.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBrMcastMembershipIntvl

	// IflaBrMcastQuerierIntvl - IFLA_BR_MCAST_QUERIER_INTVL - The interval between general queries sent by the
	// bridge acting as a multicast querier.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBrMcastQuerierIntvl

	// IflaBrMcastQueryIntvl - IFLA_BR_MCAST_QUERY_INTVL - The interval between multicast queries sent by the bridge.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBrMcastQueryIntvl

	// IflaBrMcastQueryResponseIntvl - IFLA_BR_MCAST_QUERY_RESPONSE_INTVL - The maximum response time advertised
	// in IGMP/MLD queries.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBrMcastQueryResponseIntvl

	// IflaBrMcastStartupQueryIntvl - IFLA_BR_MCAST_STARTUP_QUERY_INTVL - The interval between startup queries
	// sent by the bridge when it first becomes a multicast querier.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBrMcastStartupQueryIntvl

	// IflaBrNfCallIptables - IFLA_BR_NF_CALL_IPTABLES - Indicates if iptables should be called for bridged IPv4
	// packets.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBrNfCallIptables

	// IflaBrNfCallIp6tables - IFLA_BR_NF_CALL_IP6TABLES - Indicates if ip6tables should be called for bridged
	// IPv6 packets.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBrNfCallIp6tables

	// IflaBrNfCallArptables - IFLA_BR_NF_CALL_ARPTABLES - Indicates if arptables should be called for bridged
	// ARP packets.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBrNfCallArptables

	// IflaBrVlanDefaultPvid - IFLA_BR_VLAN_DEFAULT_PVID - The default VLAN ID for untagged packets on the bridge.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBrVlanDefaultPvid

	// IflaBrPad - IFLA_BR_PAD - Padding attribute for alignment purposes, not used directly.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBrPad

	// IflaBrVlanStatsEnabled - IFLA_BR_VLAN_STATS_ENABLED - Indicates if VLAN statistics are enabled on the bridge.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBrVlanStatsEnabled

	// IflaBrMcastStatsEnabled - IFLA_BR_MCAST_STATS_ENABLED - Indicates if multicast statistics are enabled on
	// the bridge.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBrMcastStatsEnabled

	// IflaBrMcastIgmpVersion - IFLA_BR_MCAST_IGMP_VERSION - The IGMP version used by the bridge for multicast
	// communication.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBrMcastIgmpVersion

	// IflaBrMcastMldVersion - IFLA_BR_MCAST_MLD_VERSION - The MLD version used by the bridge for multicast
	// communication.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBrMcastMldVersion

	// IflaBrMax - IFLA_BR_MAX - The maximum value for bridge attributes, used for validation purposes.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBrMax = iota - 1
)
