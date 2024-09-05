package core

// TcaTunnelKeyEnum - Represents various tunnel key actions in Linux traffic control.
//
// This enum mirrors the constants in the Linux kernel related to tunnel key encapsulation parameters.
// Tunnel keys are used to encapsulate traffic in various protocols such as IP-in-IP, GRE, or VXLAN.
// These parameters define different attributes for the tunnel key, such as source/destination addresses,
// encapsulation key, and time-to-live (TTL).
//
// For further details, refer to the Linux Kernel Documentation:
// - https://github.com/torvalds/linux/blob/master/include/uapi/linux/tc_act/tc_tunnel_key.h
// - https://man7.org/linux/man-pages/man8/tc-tunnel_key.8.html
type TcaTunnelKeyEnum uint8

const (
	// TcaTunnelKeyUnspec - TCA_TUNNEL_KEY_UNSPEC - Unspecified action for tunnel key.
	//
	// This is used as a placeholder or for actions that are not explicitly specified.
	TcaTunnelKeyUnspec TcaTunnelKeyEnum = iota

	// TcaTunnelKeyTm - TCA_TUNNEL_KEY_TM - Time-related metadata for tunnel key.
	//
	// This attribute stores time metadata such as timestamps.
	TcaTunnelKeyTm TcaTunnelKeyEnum = iota

	// TcaTunnelKeyParms - TCA_TUNNEL_KEY_PARMS - Parameters for the tunnel key encapsulation.
	//
	// This includes source and destination IPs, encapsulation options, and more.
	TcaTunnelKeyParms TcaTunnelKeyEnum = iota

	// TcaTunnelKeyEncIpv4Src - TCA_TUNNEL_KEY_ENC_IPV4_SRC - Source IPv4 address for encapsulation.
	//
	// This attribute holds the IPv4 source address for the tunnel key encapsulation.
	TcaTunnelKeyEncIpv4Src TcaTunnelKeyEnum = iota

	// TcaTunnelKeyEncIpv4Dst - TCA_TUNNEL_KEY_ENC_IPV4_DST - Destination IPv4 address for encapsulation.
	//
	// This attribute holds the IPv4 destination address for the tunnel key encapsulation.
	TcaTunnelKeyEncIpv4Dst TcaTunnelKeyEnum = iota

	// TcaTunnelKeyEncIpv6Src - TCA_TUNNEL_KEY_ENC_IPV6_SRC - Source IPv6 address for encapsulation.
	//
	// This attribute holds the IPv6 source address for the tunnel key encapsulation.
	TcaTunnelKeyEncIpv6Src TcaTunnelKeyEnum = iota

	// TcaTunnelKeyEncIpv6Dst - TCA_TUNNEL_KEY_ENC_IPV6_DST - Destination IPv6 address for encapsulation.
	//
	// This attribute holds the IPv6 destination address for the tunnel key encapsulation.
	TcaTunnelKeyEncIpv6Dst TcaTunnelKeyEnum = iota

	// TcaTunnelKeyEncKeyId - TCA_TUNNEL_KEY_ENC_KEY_ID - Encapsulation key identifier.
	//
	// This attribute is used to identify the encapsulation key for the tunnel.
	TcaTunnelKeyEncKeyId TcaTunnelKeyEnum = iota

	// TcaTunnelKeyPad - TCA_TUNNEL_KEY_PAD - Padding for alignment or future use.
	//
	// This attribute is used for padding, ensuring proper alignment in the structure.
	TcaTunnelKeyPad TcaTunnelKeyEnum = iota

	// TcaTunnelKeyEncDstPort - TCA_TUNNEL_KEY_ENC_DST_PORT - Encapsulation destination port.
	//
	// This attribute represents the destination port for UDP-based encapsulation, such as VXLAN.
	TcaTunnelKeyEncDstPort TcaTunnelKeyEnum = iota

	// TcaTunnelKeyNoCsum - TCA_TUNNEL_KEY_NO_CSUM - Disable checksum calculation.
	//
	// This flag indicates that no checksum is needed for the encapsulated traffic.
	TcaTunnelKeyNoCsum TcaTunnelKeyEnum = iota

	// TcaTunnelKeyEncOpts - TCA_TUNNEL_KEY_ENC_OPTS - Additional encapsulation options.
	//
	// This attribute stores additional options for the encapsulation process.
	TcaTunnelKeyEncOpts TcaTunnelKeyEnum = iota

	// TcaTunnelKeyEncTos - TCA_TUNNEL_KEY_ENC_TOS - Encapsulation Type of Service (TOS).
	//
	// This attribute defines the Type of Service (TOS) for the encapsulated packets.
	TcaTunnelKeyEncTos TcaTunnelKeyEnum = iota

	// TcaTunnelKeyEncTtl - TCA_TUNNEL_KEY_ENC_TTL - Encapsulation Time to Live (TTL).
	//
	// This attribute specifies the TTL value for the encapsulated traffic.
	TcaTunnelKeyEncTtl TcaTunnelKeyEnum = iota

	// TcaTunnelKeyMax - TCA_TUNNEL_KEY_MAX - Maximum value for tunnel key enum.
	//
	// This constant defines the maximum value for the tunnel key enum, often used for boundary checking.
	TcaTunnelKeyMax = iota - 1
)
