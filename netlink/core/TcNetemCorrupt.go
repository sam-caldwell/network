package core

// TcNetemCorrupt - tc_netem_corrupt - Represents the tc_netem_corrupt structure, which specifies packet corruption
// settings in the Linux Traffic Control (netem) system.
//
// The structure defines the probability and correlation of packet corruption. These settings are used to simulate
// packet corruption in a network environment.
//
// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_sched.h
type TcNetemCorrupt struct {

	// Probability - The likelihood of packet corruption (0=none ~0=100%)
	Probability uint32

	// Correlation - The correlation factor for packet corruption
	Correlation uint32
}
