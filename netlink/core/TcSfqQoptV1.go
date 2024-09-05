package core

// TcSfqQoptV1 - Represents the extended options for the SFQ (Stochastic Fair Queueing) qdisc (version 1).
//
// This struct extends the basic tc_sfq_qopt with additional fields to manage
// SFQRED parameters and statistics. It includes attributes such as depth, head-drop, and RED (Random Early Detection) parameters.
// RED is a congestion avoidance algorithm, and SFQRED integrates RED with SFQ to provide better queue management.
//
// For more information, refer to:
// - Linux Kernel Source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_sched.h
// - SFQRED Documentation: https://man7.org/linux/man-pages/man8/tc-sfq.8.html
type TcSfqQoptV1 struct {
	// TcSfqQopt - Basic SFQ options inherited from TcSfqQopt
	TcSfqQopt
	// Depth - Maximum number of packets per flow
	Depth uint32
	// HeadDrop - Whether head drop is enabled
	HeadDrop uint32
	// Limit - Maximum flow queue length (in bytes)
	Limit uint32
	// QthMin - Minimum average length threshold (in bytes)
	QthMin uint32
	// QthMax - Maximum average length threshold (in bytes)
	QthMax uint32
	// Wlog - Logarithmic weight for RED
	Wlog byte
	// Plog - Logarithm of the maximum probability for RED
	Plog byte
	// ScellLog - Cell size for idle damping
	ScellLog byte
	// Flags - SFQRED configuration flags
	Flags byte
	// MaxP - High-resolution maximum probability
	MaxP uint32
	// TcSfqRedStats - SFQRED statistics
	TcSfqRedStats
}
