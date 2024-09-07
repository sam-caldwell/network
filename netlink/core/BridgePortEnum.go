package core

// BridgePortEnum - Bridge mode enum definition
//
// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
type BridgePortEnum uint8

const (

	// IflaBrPortUnspec - IFLA_BRPORT_UNSPEC - unspecified value used as a placeholder
	IflaBrPortUnspec BridgePortEnum = 0

	// IflaBrPortState - IFLA_BRPORT_STATE - Spanning tree state
	//
	//  The operation state of the port. Here are the valid values.
	//
	//     0 - port is in STP *DISABLED* state. Make this port completely inactive for STP. This is also called BPDU
	//         filter and could be used to disable STP on an untrusted port, like a leaf virtual device. The traffic
	//         forwarding is also stopped on this port.
	//     1 - port is in STP *LISTENING* state. Only valid if STP is enabled on the bridge. In this state the port
	//         listens for STP BPDUs and drops all other traffic frames.
	//     2 - port is in STP *LEARNING* state. Only valid if STP is enabled on the bridge. In this state the port
	//         will accept traffic only for the purpose of updating MAC address tables.
	//     3 - port is in STP *FORWARDING* state. Port is fully active.
	//     4 - port is in STP *BLOCKING* state. Only valid if STP is enabled on the bridge. This state is used during
	//         the STP election process. In this state, port will only process STP BPDUs.
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBrPortState BridgePortEnum = 1

	// IflaBrPortPriority - IFLA_BRPORT_PRIORITY - The STP port priority. The valid values are between 0 and 255.
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBrPortPriority BridgePortEnum = 2

	// IflaBrPortCost -IFLA_BRPORT_COST - The STP path cost of the port. The valid values are between 1 and 65535.
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBrPortCost BridgePortEnum = 3

	// IflaBrPortMode - IFLA_BRPORT_MODE - Set the bridge port (hairpin) mode.
	//
	// See *BRIDGE_MODE_HAIRPIN* for more details.
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBrPortMode BridgePortEnum = 4

	// IflaBrPortGuard - IFLA_BRPORT_GUARD - Controls whether STP BPDUs will be processed by the bridge port. By
	// default, the flag is turned off to allow BPDU processing. Turning this flag on will disable the bridge port
	// if an STP BPDU packet is received.
	//
	// If the bridge has Spanning Tree enabled, hostile devices on the network may send BPDU on a port and cause
	// network failure. Setting *guard on* will detect and stop this by disabling the port. The port will be restarted
	// if the link is brought down, or removed and reattached.
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBrPortGuard BridgePortEnum = 5

	// IflaBrPortProtect - IFLA_BRPORT_PROTECT - Controls whether a given port is allowed to become a root port or
	// not. Only used when STP is enabled on the bridge. By default, the flag is off. This feature is also called root
	// port guard. If BPDU is received from a leaf (edge) port, it should not be elected as root port. This could be
	// used if using STP on a bridge and the downstream bridges are not fully trusted; this prevents a hostile guest
	// from rerouting traffic.
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBrPortProtect BridgePortEnum = 6 /* root port protection    */

	// IflaBrPortFastLeave - IFLA_BRPORT_FAST_LEAVE - This flag allows the bridge to immediately stop multicast traffic
	// forwarding on a port that receives an IGMP Leave message. It is only used when IGMP snooping is enabled on the
	// bridge. By default, the flag is off.
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBrPortFastLeave BridgePortEnum = 7 /* multicast fast leave    */

	// IflaBrPortLearning - IFLA_BRPORT_LEARNING - Controls whether a given port will learn *source* MAC addresses
	// from received traffic or not. Also controls whether dynamic FDB entries (which can also be added by software)
	// will be refreshed by incoming traffic. By default, this flag is on.
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBrPortLearning BridgePortEnum = 8 /* mac learning */

	// IflaBrPortUnicastFlood - IFLA_BRPORT_UNICAST_FLOOD -  Controls whether unicast traffic for which there is no
	// FDB entry will be flooded towards this port. By default, this flag is on.
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBrPortUnicastFlood BridgePortEnum = 9 /* flood unicast traffic */

	// IflaBrPortProxyArp - IFLA_BRPORT_PROXYARP - Enable proxy ARP on this port.
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBrPortProxyArp BridgePortEnum = 10 /* proxy ARP */

	// IflaBrPortLearningSync - IFLA_BRPORT_LEARNING_SYNC -Controls whether a given port will sync MAC addresses
	// learned on device port to bridge FDB.
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBrPortLearningSync BridgePortEnum = 11 /* mac learning sync from device */

	// IflaBrPortProxyArpWifi -IFLA_BRPORT_PROXYARP_WIFI - Enable proxy ARP on this port which meets extended
	// requirements by IEEE 802.11 and Hotspot 2.0 specifications.
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBrPortProxyArpWifi BridgePortEnum = 12 /* proxy ARP for Wi-Fi */

	// IflaBrPortRootId - IFLA_BRPORT_ROOT_ID - designated root
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBrPortRootId BridgePortEnum = 13 /* designated root */

	// IflaBrPortBridgeId - IFLA_BRPORT_BRIDGE_ID - designated bridge
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBrPortBridgeId BridgePortEnum = 14 /* designated bridge */

	// IflaBrPortDesignatedPort - IFLA_BRPORT_DESIGNATED_PORT -
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBrPortDesignatedPort BridgePortEnum = 15

	// IflaBrPortDesignatedCost - IFLA_BRPORT_DESIGNATED_COST -
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBrPortDesignatedCost BridgePortEnum = 16

	// IflaBrPortId - IFLA_BRPORT_ID -
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBrPortId BridgePortEnum = 17

	// IflaBrPortNo - IFLA_BRPORT_NO -
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBrPortNo BridgePortEnum = 18

	// IflaBrPortTopologyChangeAck - IFLA_BRPORT_TOPOLOGY_CHANGE_ACK -
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBrPortTopologyChangeAck BridgePortEnum = 19

	// IflaBrPortConfigPending - IFLA_BRPORT_CONFIG_PENDING -
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBrPortConfigPending BridgePortEnum = 20

	// IflaBrPortMessageAgeTimer - IFLA_BRPORT_MESSAGE_AGE_TIMER -
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBrPortMessageAgeTimer BridgePortEnum = 21

	// IflaBrPortForwardDelayTimer - IFLA_BRPORT_FORWARD_DELAY_TIMER -
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBrPortForwardDelayTimer BridgePortEnum = 22

	// IflaBrPortHoldTimer - IFLA_BRPORT_HOLD_TIMER -
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBrPortHoldTimer BridgePortEnum = 23

	// IflaBrPortFlush - IFLA_BRPORT_FLUSH - Flush bridge ports' fdb dynamic entries.
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBrPortFlush BridgePortEnum = 24

	// IflaBrPortMulticastRouter - IFLA_BRPORT_MULTICAST_ROUTER - Configure the port's multicast router presence.
	// A port with a multicast router will receive all multicast traffic.
	//
	// The valid values are:
	//   0 disable multicast routers on this port
	//   1 let the system detect the presence of routers (default)
	//   2 permanently enable multicast traffic forwarding on this port
	//   3 enable multicast routers temporarily on this port, not depending
	//     on incoming queries.
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBrPortMulticastRouter BridgePortEnum = 25

	// IflaBrPortPad - IFLA_BRPORT_PAD -
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBrPortPad BridgePortEnum = 26

	// IflaBrPortMcastFlood - IFLA_BRPORT_MCAST_FLOOD - Controls whether a given port will flood multicast traffic
	// for which there is no MDB entry. By default, this flag is on.
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBrPortMcastFlood BridgePortEnum = 27

	// IflaBrPortMcastToUcast - IFLA_BRPORT_MCAST_TO_UCAST - Controls whether a given port will replicate packets using
	// unicast instead of multicast. By default, this flag is off.
	//
	// This is done by copying the packet per host and changing the multicast destination MAC to an unicast one
	// accordingly.
	//
	// *mcast_to_unicast* works on top of the multicast snooping feature of the bridge. Which means unicast copies
	// are only delivered to hosts which are interested in unicast and signaled this via IGMP/MLD reports previously.
	//
	// This feature is intended for interface types which have a more reliable and/or efficient way to deliver unicast
	// packets than broadcast ones (e.g. Wi-Fi).
	//
	// However, it should only be enabled on interfaces where no IGMPv2/MLDv1 report suppression takes place.
	// IGMP/MLD report suppression issue is usually overcome by the network daemon (supplicant) enabling AP
	// isolation and by that separating all STAs.
	//
	// Delivery of STA-to-STA IP multicast is made possible again by enabling and utilizing the bridge hairpin mode,
	// which considers the incoming port as a potential outgoing port, too (see *BRIDGE_MODE_HAIRPIN* option). Hairpin
	// mode is performed after multicast snooping, therefore leading to only deliver reports to STAs running a
	// multicast router.
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBrPortMcastToUcast BridgePortEnum = 28

	// IflaBrPortVlanTunnel - IFLA_BRPORT_VLAN_TUNNEL - Controls whether vlan to tunnel mapping is enabled on the port.
	// By default, this flag is off.
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBrPortVlanTunnel BridgePortEnum = 29

	// IflaBrPortBcastFlood - IFLA_BRPORT_BCAST_FLOOD - Controls flooding of broadcast traffic on the given port.
	// By default, this flag is on.
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBrPortBcastFlood BridgePortEnum = 30

	// IflaBrPortGroupFwdMask - IFLA_BRPORT_GROUP_FWD_MASK - Set the group forward mask. This is a bitmask that is
	// applied to decide whether to forward incoming frames destined to link-local addresses. The addresses of the
	// form are 01:80:C2:00:00:0X (defaults to 0, which means the bridge does not forward any link-local frames
	// coming on this port).
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBrPortGroupFwdMask BridgePortEnum = 31

	// IflaBrPortNeighSuppress - IFLA_BRPORT_NEIGH_SUPPRESS - Controls whether neighbor discovery (arp and nd) proxy
	// and suppression is enabled on the port. By default, this flag is off.
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBrPortNeighSuppress BridgePortEnum = 32

	// IflaBrPortIsolated - IFLA_BRPORT_ISOLATED - Controls whether a given port will be isolated, which means it will
	// be able to communicate with non-isolated ports only. By default, this flag is off.
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBrPortIsolated BridgePortEnum = 33

	// IflaBrPortBackupPort - IFLA_BRPORT_BACKUP_PORT - Set a backup port. If the port loses carrier all traffic will
	// be redirected to the configured backup port. Set the value to 0 to disable it.
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBrPortBackupPort BridgePortEnum = 34

	// IflaBrPortMrpRingOpen - IFLA_BRPORT_MRP_RING_OPEN -
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBrPortMrpRingOpen BridgePortEnum = 35

	// IflaBrPortMrpInOpen - IFLA_BRPORT_MRP_IN_OPEN -
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBrPortMrpInOpen BridgePortEnum = 36

	// IflaBrPortMcastEhtHostsLimit - IFLA_BRPORT_MCAST_EHT_HOSTS_LIMIT - The number of per-port EHT hosts limit.
	// The default value is 512. Setting to 0 is not allowed.
	//
	// "Efficient Host Tracking" (EHT). Efficient Host Tracking is essential for optimizing the management of
	// multicast traffic. By limiting the number of multicast group members tracked on each port, the bridge can
	// efficiently manage and forward multicast traffic only to the relevant ports, reducing unnecessary network load
	// and improving performance.
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBrPortMcastEhtHostsLimit BridgePortEnum = 37

	// IflaBrPortMcastEhtHostsCnt - IFLA_BRPORT_MCAST_EHT_HOSTS_CNT - The current number of tracked hosts, read only.
	//
	// See IflaBrPortMcastEhtHostsLimit comment for definition of "Efficient Host Tracking" (EHT).
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBrPortMcastEhtHostsCnt BridgePortEnum = 38

	// IflaBrPortLocked - IFLA_BRPORT_LOCKED - Controls whether a port will be locked, meaning that hosts behind the
	// port will not be able to communicate through the port unless an FDB entry with the unit's MAC address is in
	// the FDB. The common use case is that hosts are allowed access through authentication with the IEEE 802.1X
	// protocol or based on whitelists. By default, this flag is off.
	//
	// Please note that secure 802.1X deployments should always use the *BR_BOOLOPT_NO_LL_LEARN* flag, to not permit
	// the bridge to populate its FDB based on link-local (EAPOL) traffic received on the port.
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBrPortLocked BridgePortEnum = 39

	// IflaBrPortMab - IFLA_BRPORT_MAB - Controls whether a port will use MAC Authentication Bypass (MAB), a technique
	// through which select MAC addresses may be allowed on a locked port, without using 802.1X authentication. Packets
	// with an unknown source MAC address generates a "locked" FDB entry on the incoming bridge port. The common use
	// case is for user space to react to these bridge FDB notifications and optionally replace the locked FDB entry
	// with a normal one, allowing traffic to pass for whitelisted MAC addresses.
	//
	// Setting this flag also requires *IFLA_BRPORT_LOCKED* and *IFLA_BRPORT_LEARNING*. *IFLA_BRPORT_LOCKED* ensures
	// that unauthorized data packets are dropped, and *IFLA_BRPORT_LEARNING* allows the dynamic FDB entries installed
	// by user space (as replacements for the locked FDB entries) to be refreshed and/or aged out.
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBrPortMab BridgePortEnum = 40

	// IflaBrPortMcastNGroups - IFLA_BRPORT_MCAST_N_GROUPS -
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBrPortMcastNGroups BridgePortEnum = 41

	// IflaBrPortMcastMaxGroups - IFLA_BRPORT_MCAST_MAX_GROUPS -Sets the maximum number of MDB entries that can be
	// registered for a given port. Attempts to register more MDB entries at the port than this limit allows will be
	// rejected, whether they are done through netlink (e.g. the bridge tool), or IGMP or MLD membership reports.
	// Setting a limit of 0 disables the limit. The default value is 0.
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBrPortMcastMaxGroups BridgePortEnum = 42

	// IflaBrPortNeighVlanSuppress - IFLA_BRPORT_NEIGH_VLAN_SUPPRESS - Controls whether neighbor discovery (arp and nd)
	// proxy and suppression is enabled for a given port. By default, this flag is off.
	//
	// Note that this option only takes effect when *IFLA_BRPORT_NEIGH_SUPPRESS* is enabled for a given port.
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBrPortNeighVlanSuppress BridgePortEnum = 43

	// IflaBrPortBackupNhid - IFLA_BRPORT_BACKUP_NHID - The FDB nexthop object ID to attach to packets being redirected
	// to a backup port that has VLAN tunnel mapping enabled (via the *IFLA_BRPORT_VLAN_TUNNEL* option). Setting a
	// value of 0 (default) has the effect of not attaching any ID.v
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBrPortBackupNhid BridgePortEnum = 44

	// IflaBrPortMax - IFLA_BRPORT_NEIGH_VLAN_SUPPRESS -
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBrPortMax = 45 - 1
)
