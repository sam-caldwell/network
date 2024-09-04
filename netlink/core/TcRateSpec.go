package core

// TcRateSpec - Represents the tc_ratespec structure used in Traffic Control (TC) in Linux.
//
// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_sched.h
type TcRateSpec struct {
	// CellLog - Specifies the logarithm of the cell size used in rate calculation.
	// It helps define the granularity of the rate control.
	CellLog uint8

	// Linklayer - Represents the link layer type (e.g., Ethernet).
	// This field determines the type of network link layer in use.
	Linklayer uint8

	// Overhead - Defines any additional overhead for the data link (e.g., Ethernet, ATM).
	// It is used to calculate the effective rate after accounting for overhead.
	Overhead uint16

	// CellAlign - Describes alignment requirements for the cell size in rate calculation.
	// This helps in ensuring proper alignment of the data in the traffic queue.
	CellAlign int16

	// Mpu - Defines the Minimum Packet Unit (MPU).
	// This parameter is the minimum size that will be used for traffic shaping calculations.
	Mpu uint16

	// Rate - The rate at which traffic is shaped, typically in bits per second (bps).
	// It defines the maximum transmission rate for the traffic.
	Rate uint32
}
