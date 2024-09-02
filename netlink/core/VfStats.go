//go:build linux

package core

// VfStats - ifla_vf_stats - represents statistics for a virtual function (VF).
//
// See https://github.com/torvalds/linux/blob/master/include/linux/if_link.h
type VfStats struct {

	// RxPackets - Number of received packets
	RxPackets uint64

	// TxPackets - Number of transmitted packets
	TxPackets uint64

	// RxBytes - Number of received bytes
	RxBytes uint64

	// TxBytes - Number of transmitted bytes
	TxBytes uint64

	// Multicast - Number of multicast packets
	Multicast uint64

	// Broadcast - Number of broadcast packets
	Broadcast uint64

	// RxDropped - Number of dropped received packets
	RxDropped uint64

	// TxDropped - Number of dropped transmitted packets
	TxDropped uint64
}
