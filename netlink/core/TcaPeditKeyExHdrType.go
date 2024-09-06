package core

// TcaPeditKeyExHdrType represents the header types used in the pedit action within the Linux kernel's traffic control (TC).
// These values define the header types that can be manipulated by the extended pedit key feature.
//
// For more details, refer to the Linux Kernel Source:
// https://github.com/torvalds/linux/blob/master/include/uapi/linux/tc_act/tc_pedit.h
type TcaPeditKeyExHdrType uint8

const (
	// TcaPeditKeyExHdrTypeNetwork - TCA_PEDIT_KEY_EX_HDR_TYPE_NETWORK - represents a network layer header.
	TcaPeditKeyExHdrTypeNetwork TcaPeditKeyExHdrType = iota

	// TcaPeditKeyExHdrTypeEth - TCA_PEDIT_KEY_EX_HDR_TYPE_ETH - represents an Ethernet header.
	TcaPeditKeyExHdrTypeEth TcaPeditKeyExHdrType = iota

	// TcaPeditKeyExHdrTypeIP4 - TCA_PEDIT_KEY_EX_HDR_TYPE_IP4 - represents an IPv4 header.
	TcaPeditKeyExHdrTypeIP4 TcaPeditKeyExHdrType = iota

	// TcaPeditKeyExHdrTypeIP6 - TCA_PEDIT_KEY_EX_HDR_TYPE_IP6 - represents an IPv6 header.
	TcaPeditKeyExHdrTypeIP6 TcaPeditKeyExHdrType = iota

	// TcaPeditKeyExHdrTypeTCP - TCA_PEDIT_KEY_EX_HDR_TYPE_TCP - represents a TCP header.
	TcaPeditKeyExHdrTypeTCP TcaPeditKeyExHdrType = iota

	// TcaPeditKeyExHdrTypeUDP - TCA_PEDIT_KEY_EX_HDR_TYPE_UDP - represents a UDP header.
	TcaPeditKeyExHdrTypeUDP TcaPeditKeyExHdrType = iota

	// PeditHdrTypeMax - __PEDIT_HDR_TYPE_MAX - represents the maximum number of header types.
	PeditHdrTypeMax = iota - 1
)
