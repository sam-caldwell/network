package core

// LwTunnelEncapsulation defines the encapsulation types for lightweight tunnels (LWTunnel).
// Lightweight tunnels are used to encapsulate packets and route them over different network protocols.
//
// These types provide support for various encapsulation methods, such as MPLS, IP, and IPv6.
//
// For more details, refer to the following links:
// https://github.com/torvalds/linux/blob/master/net/core/lwtunnel.c
// https://datatracker.ietf.org/doc/html/rfc4023
type LwTunnelEncapsulation uint8

const (
	// LwtunnelEncapNone - No encapsulation.
	// This type is used when no encapsulation is required for the tunnel.
	LwtunnelEncapNone LwTunnelEncapsulation = 0

	// LwtunnelEncapMpls - MPLS encapsulation.
	// Used for encapsulating packets in MPLS (Multi-Protocol Label Switching).
	// MPLS allows efficient data transport across networks.
	LwtunnelEncapMpls LwTunnelEncapsulation = 1

	// LwtunnelEncapIp - IP encapsulation.
	// Standard IP encapsulation for tunneling IPv4 packets.
	LwtunnelEncapIp LwTunnelEncapsulation = 2

	// LwtunnelEncapIla - Identifier Locator Addressing (ILA) encapsulation.
	// Used in ILA, which separates the identity and location of network elements.
	LwtunnelEncapIla LwTunnelEncapsulation = 3

	// LwtunnelEncapIp6 - IPv6 encapsulation.
	// Similar to LwtunnelEncapIp, but for tunneling IPv6 packets.
	LwtunnelEncapIp6 LwTunnelEncapsulation = 4

	// LwtunnelEncapSeg6 - Segment Routing with IPv6 (SRv6).
	// Used for segment routing in IPv6, where packets are routed based on predefined segments.
	LwtunnelEncapSeg6 LwTunnelEncapsulation = 5

	// LwtunnelEncapBpf - BPF-based encapsulation.
	// Allows the use of BPF (Berkeley Packet Filter) programs for packet encapsulation.
	LwtunnelEncapBpf LwTunnelEncapsulation = 6

	// LwtunnelEncapSeg6Local - Segment Routing with IPv6 (local actions).
	// Used for local actions in SRv6, providing more fine-grained control over packet handling.
	LwtunnelEncapSeg6Local LwTunnelEncapsulation = 7
)
