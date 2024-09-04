package core

// TcNetemRate represents the rate settings for a netem qdisc in the Linux Traffic Control system.
//
// This struct defines the rate parameters such as the actual rate in bytes per second and any overhead
// related to the packet or cell size.
//
// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_sched.h
type TcNetemRate struct {
	// Rate - Defines the bytes per second (bps) rate for the netem qdisc.
	Rate uint32

	// PacketOverhead - The additional overhead (in bytes) added to each packet.
	PacketOverhead int32

	// CellSize - The size of each cell in bytes for rate-limited cells (used for ATM or other technologies).
	CellSize uint32

	// CellOverhead - The overhead in bytes added to each cell.
	CellOverhead int32
}
