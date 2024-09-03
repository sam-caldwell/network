//go:build geneve

package core

import "unsafe"

// GenevePolicy represents the Geneve tunnel attribute policy in Go.
var GenevePolicy = [IflaGeneveMax + 1]NlaPolicy{
	// IFLA_GENEVE_UNSPEC - Unspecified attribute.
	// This is a placeholder for an unspecified or unknown attribute.
	IflaGeneveUnspec: {
		StrictStartType: IflaGeneveInnerProtoInherit,
	},

	// IFLA_GENEVE_ID - Geneve Tunnel ID.
	// 32-bit identifier for the tunnel.
	IflaGeneveId: {
		Type: NlaU32,
	},

	// IFLA_GENEVE_REMOTE - Remote IPv4 address for the Geneve tunnel.
	// Size of the daddr field in the iphdr struct (IPv4 address).
	IflaGeneveRemote: {
		Len: int(unsafe.Sizeof(uint32(0))),
	},

	// IFLA_GENEVE_REMOTE6 - Remote IPv6 address for the Geneve tunnel.
	// Size of the in6_addr struct (IPv6 address).
	IflaGeneveRemote6: {
		Len: int(unsafe.Sizeof([16]byte{})),
	},

	// IFLA_GENEVE_TTL - Time to Live (TTL) for Geneve tunnel packets.
	// 8-bit TTL value.
	IflaGeneveTtl: {
		Type: NlaU8,
	},

	// IFLA_GENEVE_TOS - Type of Service (TOS) for Geneve tunnel packets.
	// 8-bit TOS value.
	IflaGeneveTos: {
		Type: NlaU8,
	},

	// IFLA_GENEVE_LABEL - Traffic Class label for the Geneve tunnel.
	// 32-bit label for classifying traffic.
	IflaGeneveLabel: {
		Type: NlaU32,
	},

	// IFLA_GENEVE_PORT - UDP port number for the Geneve tunnel.
	// 16-bit UDP port number.
	IflaGenevePort: {
		Type: NlaU16,
	},

	// IFLA_GENEVE_COLLECT_METADATA - Collect metadata flag.
	// Boolean flag to indicate metadata collection.
	IflaGeneveCollectMetadata: {
		Type: NlaFlag,
	},

	// IFLA_GENEVE_UDP_CSUM - UDP checksum setting for the Geneve tunnel.
	// 8-bit value indicating UDP checksum behavior.
	IflaGeneveUdpCsum: {
		Type: NlaU8,
	},

	// IFLA_GENEVE_UDP_ZERO_CSUM6_TX - Zero UDP checksum for IPv6 transmit.
	// 8-bit value indicating zero checksum for IPv6 transmit.
	IflaGeneveUdpZeroCsum6Tx: {
		Type: NlaU8,
	},

	// IFLA_GENEVE_UDP_ZERO_CSUM6_RX - Zero UDP checksum for IPv6 receive.
	// 8-bit value indicating zero checksum for IPv6 receive.
	IflaGeneveUdpZeroCsum6Rx: {
		Type: NlaU8,
	},

	// IFLA_GENEVE_TTL_INHERIT - TTL inheritance setting.
	// 8-bit value indicating TTL inheritance from inner packet.
	IflaGeneveTtlInherit: {
		Type: NlaU8,
	},

	// IFLA_GENEVE_DF - Don't Fragment flag.
	// 8-bit value indicating DF flag behavior.
	IflaGeneveDf: {
		Type: NlaU8,
	},

	// IFLA_GENEVE_INNER_PROTO_INHERIT - Inner protocol inheritance flag.
	// Boolean flag to indicate protocol inheritance from the inner packet.
	IflaGeneveInnerProtoInherit: {
		Type: NlaFlag,
	},
}
