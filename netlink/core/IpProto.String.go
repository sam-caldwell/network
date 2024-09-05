package core

import "fmt"

// String - Returns the string representation of the IP protocol.
// This function maps specific protocol constants (such as TCP, UDP, SCTP, etc.)
// to their corresponding string names. If the protocol doesn't match any predefined
// case, the function returns the numeric value of the protocol.
//
// Supported protocols:
// - IPPROTO_TCP: "tcp"
// - IPPROTO_UDP: "udp"
// - IPPROTO_SCTP: "sctp"
// - IPPROTO_ICMP: "icmp"
// - IPPROTO_ICMPV6: "icmpv6"
//
// Any other protocol values will be returned as the numeric protocol identifier.
//
// Returns a string representing the IP protocol.
func (i *IPProto) String() string {
	switch *i {
	case IPPROTO_TCP:
		return "tcp"
	case IPPROTO_UDP:
		return "udp"
	case IPPROTO_SCTP:
		return "sctp"
	case IPPROTO_ICMP:
		return "icmp"
	case IPPROTO_ICMPV6:
		return "icmpv6"
	default:
		return fmt.Sprintf("%d", *i) // Return numeric value if protocol not matched
	}
}
