package core

// TcaFlowerEnum - Enumeration of Flower filter keys for Traffic Control (TC).
//
// Flower is a TC filter classifier that allows matching on a wide variety of
// fields in packet headers, including Ethernet, IP, TCP, and others.
//
// For more details on Flower classifier and its keys, see:
// - https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_cls.h
// - https://man7.org/linux/man-pages/man8/tc-flower.8.html
type TcaFlowerEnum uint8

const (
	// TcaFlowerUnspec - TCA_FLOWER_UNSPEC - Unspecified Flower attribute.
	// Placeholder for an undefined Flower key.
	TcaFlowerUnspec TcaFlowerEnum = iota

	// TcaFlowerClassID - TCA_FLOWER_CLASSID - Class ID for classification.
	// Used to assign packets to specific traffic classes.
	TcaFlowerClassID

	// TcaFlowerIndev - TCA_FLOWER_INDEV - Input device.
	// Filter packets based on the ingress network interface.
	TcaFlowerIndev

	// TcaFlowerAct - TCA_FLOWER_ACT - Associated actions.
	// List of actions to perform on matching packets.
	TcaFlowerAct

	// TcaFlowerKeyEthDst - TCA_FLOWER_KEY_ETH_DST - Ethernet destination MAC address.
	// Matches packets based on the destination MAC address.
	TcaFlowerKeyEthDst

	// TcaFlowerKeyEthDstMask - TCA_FLOWER_KEY_ETH_DST_MASK - Mask for destination MAC.
	// Apply a mask for matching on the destination MAC address.
	TcaFlowerKeyEthDstMask

	// TcaFlowerKeyEthSrc - TCA_FLOWER_KEY_ETH_SRC - Ethernet source MAC address.
	// Matches packets based on the source MAC address.
	TcaFlowerKeyEthSrc

	// TcaFlowerKeyEthSrcMask - TCA_FLOWER_KEY_ETH_SRC_MASK - Mask for source MAC.
	// Apply a mask for matching on the source MAC address.
	TcaFlowerKeyEthSrcMask

	// TcaFlowerKeyEthType - TCA_FLOWER_KEY_ETH_TYPE - Ethernet type.
	// Matches based on the EtherType field (e.g., IPv4, IPv6).
	TcaFlowerKeyEthType

	// TcaFlowerKeyIpProto - TCA_FLOWER_KEY_IP_PROTO - IP protocol.
	// Matches packets based on the IP protocol (e.g., TCP, IpProtoUDP).
	TcaFlowerKeyIpProto

	// TcaFlowerKeyIpv4Src - TCA_FLOWER_KEY_IPV4_SRC - IPv4 source address.
	// Matches packets based on the source IPv4 address.
	TcaFlowerKeyIpv4Src

	// TcaFlowerKeyIpv4SrcMask - TCA_FLOWER_KEY_IPV4_SRC_MASK - Mask for IPv4 source address.
	// Apply a mask for matching on the source IPv4 address.
	TcaFlowerKeyIpv4SrcMask

	// TcaFlowerKeyIpv4Dst - TCA_FLOWER_KEY_IPV4_DST - IPv4 destination address.
	// Matches packets based on the destination IPv4 address.
	TcaFlowerKeyIpv4Dst

	// TcaFlowerKeyIpv4DstMask - TCA_FLOWER_KEY_IPV4_DST_MASK - Mask for IPv4 destination address.
	// Apply a mask for matching on the destination IPv4 address.
	TcaFlowerKeyIpv4DstMask

	// TcaFlowerKeyIpv6Src - TCA_FLOWER_KEY_IPV6_SRC - IPv6 source address.
	// Matches packets based on the source IPv6 address.
	TcaFlowerKeyIpv6Src

	// TcaFlowerKeyIpv6SrcMask - TCA_FLOWER_KEY_IPV6_SRC_MASK - Mask for IPv6 source address.
	// Apply a mask for matching on the source IPv6 address.
	TcaFlowerKeyIpv6SrcMask

	// TcaFlowerKeyIpv6Dst - TCA_FLOWER_KEY_IPV6_DST - IPv6 destination address.
	// Matches packets based on the destination IPv6 address.
	TcaFlowerKeyIpv6Dst

	// TcaFlowerKeyIpv6DstMask - TCA_FLOWER_KEY_IPV6_DST_MASK - Mask for IPv6 destination address.
	// Apply a mask for matching on the destination IPv6 address.
	TcaFlowerKeyIpv6DstMask

	// TcaFlowerKeyTcpSrc - TCA_FLOWER_KEY_TCP_SRC - TCP source port.
	// Matches packets based on the TCP source port.
	TcaFlowerKeyTcpSrc

	// TcaFlowerKeyTcpDst - TCA_FLOWER_KEY_TCP_DST - TCP destination port.
	// Matches packets based on the TCP destination port.
	TcaFlowerKeyTcpDst

	// TcaFlowerKeyUdpSrc - TCA_FLOWER_KEY_UDP_SRC - IpProtoUDP source port.
	// Matches packets based on the IpProtoUDP source port.
	TcaFlowerKeyUdpSrc

	// TcaFlowerKeyUdpDst - TCA_FLOWER_KEY_UDP_DST - IpProtoUDP destination port.
	// Matches packets based on the IpProtoUDP destination port.
	TcaFlowerKeyUdpDst

	// More keys continue, following the same pattern for fields like VLAN, MPLS, SCTP, IpProtoIcmp, and encapsulation options.

	// TcaFlowerMax - Placeholder for the maximum value of the Flower key.
	__TcaFlowerMax = iota - 1
)
