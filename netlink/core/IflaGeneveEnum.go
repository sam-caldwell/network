package core

// IflaGeneveEnum - Enumeration for Geneve Tunnel Attributes
//
// GENEVE: Generic Network Virtualization Encapsulation
// https://github.com/torvalds/linux/blob/master/drivers/net/geneve.c
//
// This enumeration defines the various attributes used for configuring Geneve tunnels in Linux.
// Geneve is a flexible tunneling protocol used for network virtualization.

type IflaGeneveEnum uint8

const (
	// IflaGeneveUnspec - IFLA_GENEVE_UNSPEC - Unspecified attribute.
	// This is a placeholder for an unspecified or unknown attribute in the Geneve tunnel configuration.
	IflaGeneveUnspec IflaGeneveEnum = iota

	// IflaGeneveId - IFLA_GENEVE_ID - Geneve Tunnel ID.
	// This attribute specifies the Geneve tunnel ID, a 32-bit identifier for the tunnel.
	IflaGeneveId

	// IflaGeneveRemote - IFLA_GENEVE_REMOTE - Remote IPv4 address for the Geneve tunnel.
	// This attribute specifies the remote IPv4 address that the Geneve tunnel connects to.
	IflaGeneveRemote

	// IflaGeneveRemote6 - IFLA_GENEVE_REMOTE6 - Remote IPv6 address for the Geneve tunnel.
	// This attribute specifies the remote IPv6 address that the Geneve tunnel connects to.
	IflaGeneveRemote6

	// IflaGeneveTtl - IFLA_GENEVE_TTL - Time to Live (TTL) for Geneve tunnel packets.
	// This attribute specifies the TTL value for packets sent through the Geneve tunnel.
	IflaGeneveTtl

	// IflaGeneveTos - IFLA_GENEVE_TOS - Type of Service (TOS) for Geneve tunnel packets.
	// This attribute specifies the TOS field value for packets sent through the Geneve tunnel.
	IflaGeneveTos

	// IflaGeneveLabel - IFLA_GENEVE_LABEL - Traffic Class label for the Geneve tunnel.
	// This attribute specifies a 32-bit label that can be used for classifying traffic in the Geneve tunnel.
	IflaGeneveLabel

	// IflaGenevePort - IFLA_GENEVE_PORT - IpProtoUDP port number for the Geneve tunnel.
	// This attribute specifies the IpProtoUDP port number used by the Geneve tunnel.
	IflaGenevePort

	// IflaGeneveCollectMetadata - IFLA_GENEVE_COLLECT_METADATA - Collect metadata flag.
	// This attribute indicates whether the Geneve tunnel should collect metadata for use in routing or traffic
	// management.
	IflaGeneveCollectMetadata

	// IflaGeneveUdpCsum - IFLA_GENEVE_UDP_CSUM - IpProtoUDP checksum setting for the Geneve tunnel.
	// This attribute specifies whether IpProtoUDP checksums should be used for packets sent through the Geneve tunnel.
	IflaGeneveUdpCsum

	// IflaGeneveUdpZeroCsum6Tx - IFLA_GENEVE_UDP_ZERO_CSUM6_TX - Zero IpProtoUDP checksum for IPv6 transmit.
	// This attribute specifies whether to use a zero IpProtoUDP checksum for IPv6 packets transmitted by the Geneve tunnel.
	IflaGeneveUdpZeroCsum6Tx

	// IflaGeneveUdpZeroCsum6Rx - IFLA_GENEVE_UDP_ZERO_CSUM6_RX - Zero IpProtoUDP checksum for IPv6 receive.
	// This attribute specifies whether to accept packets with a zero IpProtoUDP checksum for IPv6 packets received by the
	// Geneve tunnel.
	IflaGeneveUdpZeroCsum6Rx

	// IflaGeneveTtlInherit - IFLA_GENEVE_TTL_INHERIT - TTL inheritance setting.
	// This attribute specifies whether the Geneve tunnel should inherit the TTL value from the inner packet.
	IflaGeneveTtlInherit

	// IflaGeneveDf - IFLA_GENEVE_DF - Don't Fragment flag.
	// This attribute specifies whether the Don't Fragment (DF) flag should be set on packets sent through the
	// Geneve tunnel.
	IflaGeneveDf

	// IflaGeneveInnerProtoInherit - IFLA_GENEVE_INNER_PROTO_INHERIT - Inner protocol inheritance flag.
	// This attribute indicates whether the Geneve tunnel should inherit the protocol information from the
	// encapsulated packet.
	IflaGeneveInnerProtoInherit

	// IflaGeneveMax - IFLA_GENEVE_MAX - Maximum attribute value.
	// This constant represents the maximum valid value for Geneve tunnel attributes, used for validation and boundary
	// checks.
	IflaGeneveMax = iota - 1
)
