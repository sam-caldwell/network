package core

// IflaVxLanEnum - Enumeration for VXLAN attributes in Linux kernel -
//
// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
type IflaVxLanEnum uint8

const (
	// IflaVxlanUnspec - IFLA_VXLAN_UNSPEC -
	// This is an unspecified placeholder value used in the enumeration.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaVxlanUnspec IflaVxLanEnum = 0

	// IflaVxlanId - IFLA_VXLAN_ID -
	// This attribute represents the VXLAN Network Identifier (VNI), a 24-bit identifier
	// that distinguishes the VXLAN overlay network.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaVxlanId IflaVxLanEnum = 1

	// IflaVxlanGroup - IFLA_VXLAN_GROUP -
	// This attribute represents the multicast group or remote address used for VXLAN.
	// It specifies the group to which the VXLAN packets are sent for encapsulation.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaVxlanGroup IflaVxLanEnum = 2

	// IflaVxlanLink - IFLA_VXLAN_LINK -
	// This attribute specifies the underlying physical network interface that the VXLAN
	// is bound to. It links the VXLAN interface with a physical device.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaVxlanLink IflaVxLanEnum = 3

	// IflaVxlanLocal - IFLA_VXLAN_LOCAL -
	// This attribute represents the local IP address used for VXLAN encapsulation.
	// It is the source IP address from which VXLAN packets are sent.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaVxlanLocal IflaVxLanEnum = 4

	// IflaVxlanTtl - IFLA_VXLAN_TTL -
	// This attribute specifies the Time to Live (TTL) value for VXLAN packets.
	// TTL controls the lifetime of a packet in the network, preventing it from looping indefinitely.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaVxlanTtl IflaVxLanEnum = 5

	// IflaVxlanTos - IFLA_VXLAN_TOS -
	// This attribute specifies the Type of Service (TOS) field in the VXLAN header.
	// TOS is used for QoS (Quality of Service) to prioritize traffic.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaVxlanTos IflaVxLanEnum = 6

	// IflaVxlanLearning - IFLA_VXLAN_LEARNING -
	// This attribute controls whether the VXLAN interface should perform MAC address learning.
	// When enabled, the VXLAN interface learns the MAC addresses of remote peers.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaVxlanLearning IflaVxLanEnum = 7

	// IflaVxlanAgeing - IFLA_VXLAN_AGEING -
	// This attribute specifies the ageing time for learned MAC addresses on the VXLAN interface.
	// It determines how long a learned MAC address remains in the forwarding table.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaVxlanAgeing IflaVxLanEnum = 8

	// IflaVxlanLimit - IFLA_VXLAN_LIMIT -
	// This attribute sets a limit on the number of MAC addresses that can be learned by the VXLAN interface.
	// It helps in controlling the size of the forwarding table.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaVxlanLimit IflaVxLanEnum = 9

	// IflaVxlanPortRange - IFLA_VXLAN_PORT_RANGE -
	// This attribute specifies the range of source ports that can be used for VXLAN packets.
	// Source port selection within this range can be used for load balancing.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaVxlanPortRange IflaVxLanEnum = 10

	// IflaVxlanProxy - IFLA_VXLAN_PROXY -
	// This attribute enables or disables proxy ARP/NDP for the VXLAN interface.
	// When enabled, the VXLAN interface can respond to ARP/NDP requests on behalf of other hosts.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaVxlanProxy IflaVxLanEnum = 11

	// IflaVxlanRsc - IFLA_VXLAN_RSC -
	// This attribute stands for "Route Short Circuit," and it enables optimizations that
	// bypass the routing table lookup for VXLAN packets destined for known hosts.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaVxlanRsc IflaVxLanEnum = 12

	// IflaVxlanL2miss - IFLA_VXLAN_L2MISS -
	// This attribute enables L2 miss notifications for the VXLAN interface.
	// It triggers notifications when packets are received for unknown Layer 2 destinations.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaVxlanL2miss IflaVxLanEnum = 13

	// IflaVxlanL3miss - IFLA_VXLAN_L3MISS -
	// This attribute enables L3 miss notifications for the VXLAN interface.
	// It triggers notifications when packets are received for unknown Layer 3 destinations.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaVxlanL3miss IflaVxLanEnum = 14

	// IflaVxlanPort - IFLA_VXLAN_PORT -
	// This attribute specifies the destination port used for VXLAN traffic.
	// The default port for VXLAN is 4789, but it can be changed using this attribute.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaVxlanPort IflaVxLanEnum = 15

	// IflaVxlanGroup6 - IFLA_VXLAN_GROUP6 -
	// This attribute specifies the IPv6 multicast group address used for VXLAN.
	// It allows the use of IPv6 addresses in the VXLAN encapsulation process.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaVxlanGroup6 IflaVxLanEnum = 16

	// IflaVxlanLocal6 - IFLA_VXLAN_LOCAL6 -
	// This attribute specifies the local IPv6 address used for VXLAN encapsulation.
	// It allows the VXLAN interface to operate over IPv6 networks.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaVxlanLocal6 IflaVxLanEnum = 17

	// IflaVxlanUdpCsum - IFLA_VXLAN_UDP_CSUM -
	// This attribute controls whether IpProtoUDP checksum is enabled for VXLAN packets.
	// Enabling IpProtoUDP checksum can help in verifying packet integrity.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaVxlanUdpCsum IflaVxLanEnum = 18

	// IflaVxlanUdpZeroCsum6Tx - IFLA_VXLAN_UDP_ZERO_CSUM6_TX -
	// This attribute controls whether to disable IpProtoUDP checksum for transmitted IPv6 VXLAN packets.
	// It sets the checksum to zero for outgoing IPv6 VXLAN traffic.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaVxlanUdpZeroCsum6Tx IflaVxLanEnum = 19

	// IflaVxlanUdpZeroCsum6Rx - IFLA_VXLAN_UDP_ZERO_CSUM6_RX -
	// This attribute controls whether to allow received IPv6 VXLAN packets with zero IpProtoUDP checksum.
	// It allows for the acceptance of incoming IPv6 VXLAN packets even if the checksum is zero.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaVxlanUdpZeroCsum6Rx IflaVxLanEnum = 20

	// IflaVxlanRemcsumTx - IFLA_VXLAN_REMCSUM_TX -
	// This attribute controls whether to transmit remote checksum offload (RCO) for VXLAN.
	// RCO allows the offloading of checksum calculation to remote peers.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaVxlanRemcsumTx IflaVxLanEnum = 21

	// IflaVxlanRemcsumRx - IFLA_VXLAN_REMCSUM_RX -
	// This attribute controls whether to receive remote checksum offload (RCO) for VXLAN.
	// It enables the VXLAN interface to handle packets with checksums offloaded by remote peers.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaVxlanRemcsumRx IflaVxLanEnum = 22

	// IflaVxlanGbp - IFLA_VXLAN_GBP -
	// This attribute enables Group Policy extension for VXLAN.
	// Group Policy provides additional security and policy control in VXLAN environments.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaVxlanGbp IflaVxLanEnum = 23

	// IflaVxlanRemcsumNopartial - IFLA_VXLAN_REMCSUM_NOPARTIAL -
	// This attribute disables partial checksum offload for VXLAN.
	// It ensures that the full checksum is calculated and verified, rather than relying on partial offloading.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaVxlanRemcsumNopartial IflaVxLanEnum = 24

	// IflaVxlanCollectMetadata - IFLA_VXLAN_COLLECT_METADATA -
	// This attribute enables the collection of metadata for VXLAN packets.
	// Metadata can include various additional information used for advanced networking features.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaVxlanCollectMetadata IflaVxLanEnum = 25

	// IflaVxlanLabel - IFLA_VXLAN_LABEL -
	// This attribute specifies a label for VXLAN traffic, which can be used for traffic classification or policy enforcement.
	// The label helps in distinguishing and managing different types of VXLAN traffic.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaVxlanLabel IflaVxLanEnum = 26

	// IflaVxlanGpe - IFLA_VXLAN_GPE -
	// This attribute enables VXLAN-GPE (Generic Protocol Extension).
	// VXLAN-GPE allows for the encapsulation of various protocols beyond just Ethernet, enabling broader use cases.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaVxlanGpe IflaVxLanEnum = 27

	// IflaVxlanTtlInherit - IFLA_VXLAN_TTL_INHERIT -
	// This attribute controls whether the TTL value of the inner packet should be inherited by the outer VXLAN packet.
	// TTL inheritance helps in maintaining the expected packet lifespan across the VXLAN tunnel.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaVxlanTtlInherit IflaVxLanEnum = 28

	// IflaVxlanDf - IFLA_VXLAN_DF -
	// This attribute controls the "Don't Fragment" (DF) flag for VXLAN packets.
	// Setting this flag prevents the VXLAN packets from being fragmented during transit.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaVxlanDf IflaVxLanEnum = 29

	// IflaVxlanVnifilter - IFLA_VXLAN_VNIFILTER -
	// This attribute enables VNI filtering, which restricts VXLAN traffic based on the VNI.
	// It is particularly applicable when metadata collection is enabled.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaVxlanVnifilter IflaVxLanEnum = 30

	// IflaVxlanLocalbypass - IFLA_VXLAN_LOCALBYPASS -
	// This attribute controls whether locally generated VXLAN traffic should bypass the VXLAN overlay.
	// It can be used to optimize traffic paths in certain network setups.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaVxlanLocalbypass IflaVxLanEnum = 31

	// IflaVxlanLabelPolicy - IFLA_VXLAN_LABEL_POLICY -
	// This attribute specifies the policy for handling IPv6 flow labels in VXLAN.
	// It determines how flow labels are applied to or interpreted by VXLAN interfaces.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaVxlanLabelPolicy IflaVxLanEnum = 32

	// IflaVxlanMax - IFLA_VXLAN_MAX -
	// This constant represents the maximum valid value for VXLAN attributes.
	// It is used as a boundary marker to validate or limit the range of VXLAN attributes.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
	IflaVxlanMax = IflaVxlanLabelPolicy
)
