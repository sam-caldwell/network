//go:build linux

package core

// VfStats - ifla_vf_stats - represents statistics for a virtual function (VF).
//
// See https://github.com/torvalds/linux/blob/master/include/linux/if_link.h
type VfStats struct {
	RxPackets uint64 // Number of received packets
	TxPackets uint64 // Number of transmitted packets
	RxBytes   uint64 // Number of received bytes
	TxBytes   uint64 // Number of transmitted bytes
	Multicast uint64 // Number of multicast packets
	Broadcast uint64 // Number of broadcast packets
	RxDropped uint64 // Number of dropped received packets
	TxDropped uint64 // Number of dropped transmitted packets
}
