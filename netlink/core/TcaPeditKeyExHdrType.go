package core

// TcaPeditKeyExHdrType represents the header types used in the pedit action within the Linux kernel's
// traffic control (TC). These values define the header types that can be manipulated by the extended pedit key
// feature.
//
// # See PeditHeaderType
//
// For more details, refer to the Linux Kernel Source:
// https://github.com/torvalds/linux/blob/master/include/uapi/linux/tc_act/tc_pedit.h
type TcaPeditKeyExHdrType uint16

const (
	// TcaPeditKeyExHdrTypeNetwork - TcaPeditKeyEx_HDR_TYPE_NETWORK - represents a network layer header.
	TcaPeditKeyExHdrTypeNetwork PeditHeaderType = iota

	// TcaPeditKeyExHdrTypeEth - TcaPeditKeyExHdrTypeEth - represents an Ethernet header.
	TcaPeditKeyExHdrTypeEth PeditHeaderType = iota

	// TcaPeditKeyExHdrTypeIP4 - TcaPeditKeyExHdrTypeIP4 - represents an IPv4 header.
	TcaPeditKeyExHdrTypeIP4 PeditHeaderType = iota

	// TcaPeditKeyExHdrTypeIP6 - TcaPeditKeyExHdrTypeIP6 - represents an IPv6 header.
	TcaPeditKeyExHdrTypeIP6 PeditHeaderType = iota

	// TcaPeditKeyExHdrTypeTCP - TcaPeditKeyExHdrTypeTCP - represents a TCP header.
	TcaPeditKeyExHdrTypeTCP PeditHeaderType = iota

	// TcaPeditKeyExHdrTypeUDP - TcaPeditKeyExHdrTypeUDP - represents a UDP header.
	TcaPeditKeyExHdrTypeUDP PeditHeaderType = iota

	// PeditHdrTypeMax - __PEDIT_HDR_TYPE_MAX - represents the maximum number of header types.
	PeditHdrTypeMax = iota - 1
)
