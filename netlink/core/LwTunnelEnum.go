package core

// LwTunnelEnum is an enumeration of route attribute IDs for lightweight tunnels (LWTunnel) in IPv6.
//
// These attribute IDs are derived from the Linux kernel's definitions for LWTunnel, particularly for IPv6.
// The values are specified in the Linux kernel's IP tunnel core implementation.
//
// References:
// - https://github.com/jow-/ucode/blob/master/include/linux/lwtunnel.h
// - https://github.com/torvalds/linux/blob/master/net/ipv4/ip_tunnel_core.c
type LwTunnelEnum uint8

const (
	// LwtunnelIp6Unspec - LWTUNNEL_IP6_UNSPEC - Unspecified attribute (0)
	// This is a placeholder for an unspecified or unknown attribute.
	LwtunnelIp6Unspec LwTunnelEnum = 0

	// LwtunnelIp6Id - LWTUNNEL_IP6_ID - ID attribute (1)
	// This attribute specifies the ID for the lightweight tunnel.
	LwtunnelIp6Id LwTunnelEnum = 1

	// LwtunnelIp6Dst - LWTUNNEL_IP6_DST - Destination attribute (2)
	// This attribute specifies the destination address for the tunnel.
	LwtunnelIp6Dst LwTunnelEnum = 2

	// LwtunnelIp6Src - LWTUNNEL_IP6_SRC - Source attribute (3)
	// This attribute specifies the source address for the tunnel.
	LwtunnelIp6Src LwTunnelEnum = 3

	// LwtunnelIp6Hoplimit - LWTUNNEL_IP6_HOPLIMIT - Hop limit attribute (4)
	// This attribute specifies the hop limit (TTL) for the tunnel.
	LwtunnelIp6Hoplimit LwTunnelEnum = 4

	// LwtunnelIp6Tc - LWTUNNEL_IP6_TC - Traffic class attribute (5)
	// This attribute specifies the traffic class for the tunnel.
	LwtunnelIp6Tc LwTunnelEnum = 5

	// LwtunnelIp6Flags - LWTUNNEL_IP6_FLAGS - Flags attribute (6)
	// This attribute specifies various flags for the tunnel.
	LwtunnelIp6Flags LwTunnelEnum = 6

	// LwtunnelIp6Pad - LWTUNNEL_IP6_PAD - Padding attribute (7)
	// This attribute is reserved for padding and is not implemented.
	LwtunnelIp6Pad LwTunnelEnum = 7

	// LwtunnelIp6Opts - LWTUNNEL_IP6_OPTS - Options attribute (8)
	// This attribute is reserved for additional options and is not implemented.
	LwtunnelIp6Opts LwTunnelEnum = 8

	// LwtunnelIp6Max - LWTUNNEL_IP6_MAX - Maximum attribute value (7)
	// This constant represents the maximum valid value for the LWTunnel attributes.
	LwtunnelIp6Max = iota - 1
)
