package core

// IPProto - Represents the `ip_proto` attribute in the Flower classifier.
// Flower is a flexible, flow-based packet classification mechanism used in traffic control.
//
// This enumeration maps commonly used Internet Protocol (IP) numbers to their constants, including TCP, UDP, SCTP, ICMP, and ICMPv6.
// These constants are from the `unix` package, which provides access to low-level system call interfaces.
//
// The `ip_proto` attribute in Flower is used to match the protocol number in the IP header of a packet.
//
// For more details, refer to:
// - Linux Kernel Source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_cls.h
type IPProto uint8

const (
	// IPPROTO_TCP - Transmission Control Protocol (TCP).
	// See: https://en.wikipedia.org/wiki/Transmission_Control_Protocol
	IPPROTO_TCP IPProto = unix.IPPROTO_TCP

	// IPPROTO_UDP - User Datagram Protocol (UDP).
	// See: https://en.wikipedia.org/wiki/User_Datagram_Protocol
	IPPROTO_UDP IPProto = unix.IPPROTO_UDP

	// IPPROTO_SCTP - Stream Control Transmission Protocol (SCTP).
	// See: https://en.wikipedia.org/wiki/Stream_Control_Transmission_Protocol
	IPPROTO_SCTP IPProto = unix.IPPROTO_SCTP

	// IPPROTO_ICMP - Internet Control Message Protocol (ICMP).
	// See: https://en.wikipedia.org/wiki/Internet_Control_Message_Protocol
	IPPROTO_ICMP IPProto = unix.IPPROTO_ICMP

	// IPPROTO_ICMPV6 - Internet Control Message Protocol for IPv6 (ICMPv6).
	// See: https://en.wikipedia.org/wiki/ICMPv6
	IPPROTO_ICMPV6 IPProto = unix.IPPROTO_ICMPV6
)
