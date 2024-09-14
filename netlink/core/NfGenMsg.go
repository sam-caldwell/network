//go:build linux

package core

import "unsafe"

// NfGenMsgSize - Track the message sizes for the correct serialization/deserialization
const NfGenMsgSize = int(unsafe.Sizeof(Nfgenmsg{}))

// Nfgenmsg - nfgenmsg - General form of address family dependent message.
//
// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/nfnetlink.h
type Nfgenmsg struct {
	/* Address Family (e.g., AF_INET for IPv4, AF_INET6 for IPv6) */
	NfgenFamily uint8

	/* Netfilter netlink message version */
	Version uint8

	/* Resource identifier (stored in big endian byte order) */
	ResId uint16
}
