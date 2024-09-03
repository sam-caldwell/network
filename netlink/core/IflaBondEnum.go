package core

// IflaBondEnum - Bonding section
//
// This enumeration defines various attributes related to network bonding interfaces in the Linux kernel.
// Bonding allows multiple network interfaces to be aggregated together, providing redundancy and improved performance.
//
// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
type IflaBondEnum uint8

const (

	// IflaBondUnspec - IFLA_BOND_UNSPEC -
	// This is an unspecified placeholder value used in the enumeration.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBondUnspec IflaBondEnum = 0

	// IflaBondMode - IFLA_BOND_MODE -
	// This attribute specifies the bonding mode. Bonding modes determine how the interfaces in the bond
	// interact with each other and with the network. Examples include round-robin, active-backup, and 802.3ad (LACP).
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBondMode IflaBondEnum = 1

	// IflaBondActiveSlave - IFLA_BOND_ACTIVE_SLAVE -
	// This attribute identifies the currently active slave interface in the bond. In active-backup mode, this would
	// be the interface that is actively handling traffic.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBondActiveSlave IflaBondEnum = 2

	// IflaBondMiimon - IFLA_BOND_MIIMON -
	// This attribute specifies the MII (Media Independent Interface) monitoring interval in milliseconds. This
	// interval determines how often the bond checks the link status of its slave interfaces.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBondMiimon IflaBondEnum = 3

	// IflaBondUpdelay - IFLA_BOND_UPDELAY -
	// This attribute sets the delay in milliseconds before a newly active interface in the bond begins transmitting.
	// It ensures that the interface is fully operational before it is used for traffic.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBondUpdelay IflaBondEnum = 4

	// IflaBondDowndelay - IFLA_BOND_DOWNDELAY -
	// This attribute sets the delay in milliseconds before an interface in the bond is marked as inactive. It helps
	// to avoid brief outages or link flapping from triggering a failover.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBondDowndelay IflaBondEnum = 5

	// IflaBondUseCarrier - IFLA_BOND_USE_CARRIER -
	// This attribute controls whether the bond should use carrier state to determine the availability of a slave
	// interface. When enabled, the bond relies on the interface's carrier detection to decide if it is up or down.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBondUseCarrier IflaBondEnum = 6

	// IflaBondArpInterval - IFLA_BOND_ARP_INTERVAL -
	// This attribute specifies the interval in milliseconds between ARP requests sent to monitor the state of the
	// slave interfaces. It is used in ARP monitoring mode to detect failed interfaces.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBondArpInterval IflaBondEnum = 7

	// IflaBondArpIpTarget - IFLA_BOND_ARP_IP_TARGET -
	// This attribute lists the IP addresses targeted by the ARP requests in ARP monitoring mode.
	// These targets are used to verify the connectivity of the slave interfaces.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBondArpIpTarget IflaBondEnum = 8

	// IflaBondArpValidate - IFLA_BOND_ARP_VALIDATE -
	// This attribute controls ARP validation settings for the bond. ARP validation checks that ARP replies come
	// from the correct MAC address and are sent by the correct slave interface.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBondArpValidate IflaBondEnum = 9

	// IflaBondArpAllTargets - IFLA_BOND_ARP_ALL_TARGETS -
	// This attribute specifies how the bond handles ARP targets in ARP monitoring mode. It determines whether all
	// targets must be reachable for the bond to consider a slave interface as active.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBondArpAllTargets IflaBondEnum = 10

	// IflaBondPrimary - IFLA_BOND_PRIMARY -
	// This attribute specifies the primary slave interface in the bond. The primary slave is the preferred
	// interface for transmitting traffic when it is available.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBondPrimary IflaBondEnum = 11

	// IflaBondPrimaryReselect - IFLA_BOND_PRIMARY_RESELECT -
	// This attribute controls how and when the bond should reselect the primary slave interface. Reselection criteria
	// can include the availability of the primary interface or specific events.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBondPrimaryReselect IflaBondEnum = 12

	// IflaBondFailOverMac - IFLA_BOND_FAIL_OVER_MAC -
	// This attribute controls the failover behavior for the MAC address of the bond. It specifies whether the MAC
	// address should be taken from the active slave interface or remain constant.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBondFailOverMac IflaBondEnum = 13

	// IflaBondXmitHashPolicy - IFLA_BOND_XMIT_HASH_POLICY -
	// This attribute sets the transmit hash policy for the bond, which determines how traffic is distributed among
	// slave interfaces. Policies include layer 2, layer 3, and layer 4 hashing, influencing load balancing.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBondXmitHashPolicy IflaBondEnum = 14

	// IflaBondResendIgmp - IFLA_BOND_RESEND_IGMP -
	// This attribute controls the number of IGMP membership reports sent on failover. It helps in maintaining
	// multicast group membership when failover occurs in the bond.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBondResendIgmp IflaBondEnum = 15

	// IflaBondNumPeerNotif - IFLA_BOND_NUM_PEER_NOTIF -
	// This attribute specifies the number of peer notifications to be sent when a slave interface changes state. Peer
	// notifications inform neighboring devices about the change in the bond's configuration.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBondNumPeerNotif IflaBondEnum = 16

	// IflaBondAllSlavesActive - IFLA_BOND_ALL_SLAVES_ACTIVE -
	// This attribute indicates whether all slave interfaces in the bond should be active simultaneously. It is
	// relevant in active-active configurations where multiple slaves handle traffic concurrently.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBondAllSlavesActive IflaBondEnum = 17

	// IflaBondMinLinks - IFLA_BOND_MIN_LINKS -
	// This attribute sets the minimum number of active slave interfaces required for the bond to be considered
	// operational. It ensures that the bond has sufficient redundancy before it is used for traffic.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBondMinLinks IflaBondEnum = 18

	// IflaBondLpInterval - IFLA_BOND_LP_INTERVAL -
	// This attribute specifies the interval in milliseconds for sending Link Polling (LP) packets. LP is used to
	// monitor the status of the slave interfaces and ensure they are functioning correctly.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBondLpInterval IflaBondEnum = 19

	// IflaBondPacketsPerSlave - IFLA_BOND_PACKETS_PER_SLAVE -
	// This attribute sets the number of packets to be sent through each slave interface before switching to the next.
	// It controls load balancing by distributing traffic evenly across all active slaves.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBondPacketsPerSlave IflaBondEnum = 20

	// IflaBondAdLacpRate - IFLA_BOND_AD_LACP_RATE -
	// This attribute sets the rate at which Link Aggregation Control Protocol (LACP) packets are sent. LACP is used
	// in 802.3ad bonding mode to negotiate and manage link aggregation.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBondAdLacpRate IflaBondEnum = 21

	// IflaBondAdSelect - IFLA_BOND_AD_SELECT -
	// This attribute specifies the selection criteria for Active-Backup mode in 802.3ad bonding. It determines which
	// slave interface is selected as the active link in the aggregation.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBondAdSelect IflaBondEnum = 22

	// IflaBondAdInfo - IFLA_BOND_AD_INFO -
	// This attribute contains additional information related to 802.3ad bonding, including negotiated settings. It
	// provides details about the aggregated link and its operational state.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBondAdInfo IflaBondEnum = 23

	// IflaBondAdActorSysPrio - IFLA_BOND_AD_ACTOR_SYS_PRIO -
	// This attribute sets the system priority for the Actor in 802.3ad bonding. System priority is used in the
	// LACP negotiation process to select the active links.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBondAdActorSysPrio IflaBondEnum = 24

	// IflaBondAdUserPortKey - IFLA_BOND_AD_USER_PORT_KEY -
	// This attribute specifies the user-configured key for the ports in 802.3ad bonding. The port key is used to
	// group links that belong to the same aggregated link.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBondAdUserPortKey IflaBondEnum = 25

	// IflaBondAdActorSystem - IFLA_BOND_AD_ACTOR_SYSTEM -
	// This attribute sets the system MAC address used by the Actor in 802.3ad bonding. The Actor system identifier
	// is crucial in the LACP negotiation and aggregation process.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBondAdActorSystem IflaBondEnum = 26

	// IflaBondTlbDynamicLb - IFLA_BOND_TLB_DYNAMIC_LB -
	// This attribute enables dynamic load balancing in the TLB (Transmit Load Balancing) bonding mode. Dynamic load
	// balancing adjusts traffic distribution based on the current load on each slave.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBondTlbDynamicLb IflaBondEnum = 27

	// IflaBondPeerNotifDelay - IFLA_BOND_PEER_NOTIF_DELAY -
	// This attribute sets the delay before sending notifications to peers when a bond's configuration changes.
	// The delay helps to ensure that the bond's state is stable before peers are informed of the change.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBondPeerNotifDelay IflaBondEnum = 28

	// IflaBondAdLacpActive - IFLA_BOND_AD_LACP_ACTIVE -
	// This attribute controls whether LACP is in active or passive mode in 802.3ad bonding.
	// Active mode means that the bond will actively negotiate link aggregation with peers.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBondAdLacpActive IflaBondEnum = 29

	// IflaBondMissedMax - IFLA_BOND_MISSED_MAX -
	// This attribute sets the maximum number of missed LACPDU frames before the bond considers a link to be down.
	// It helps in determining the stability and reliability of the aggregated links.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBondMissedMax IflaBondEnum = 30

	// IflaBondNsIp6Target - IFLA_BOND_NS_IP6_TARGET -
	// This attribute specifies IPv6 targets for Neighbor Solicitation (NS) in monitoring mode.
	// NS targets are used to verify the reachability of IPv6 addresses via the bond's slave interfaces.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBondNsIp6Target IflaBondEnum = 31

	// IflaBondCoupledControl - IFLA_BOND_COUPLED_CONTROL -
	// This attribute controls whether the bond operates in coupled control mode.
	// Coupled control mode coordinates the operation of the bond with other networking features, such as STP or LACP.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBondCoupledControl IflaBondEnum = 32

	// IflaBondMax - IFLA_BOND_MAX -
	// This constant represents the maximum valid value for bonding attributes.
	// It is used as a boundary marker to validate or limit the range of bonding attributes.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaBondMax IflaBondEnum = iota - 1
)
