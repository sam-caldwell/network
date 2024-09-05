package core

// TcSfqRedStats - tc_sfqred_stats - Structure that holds statistics for the Stochastic Fair Queueing (SFQ) with Random Early Detection (RED).
//
// SFQ-RED is a combination of the Stochastic Fair Queueing (SFQ) and Random Early Detection (RED) algorithms.
// It is used in network traffic management to ensure fair bandwidth distribution while avoiding congestion.
//
// For more details on SFQ-RED, you can refer to the Linux Kernel source code:
// - https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_sched.h
type TcSfqRedStats struct {
	// ProbDrop - Count of packets dropped early, below the maximum threshold.
	ProbDrop uint32
	// ForcedDrop - Count of packets forcibly dropped, after the maximum threshold.
	ForcedDrop uint32
	// ProbMark - Count of packets marked early, below the maximum threshold.
	ProbMark uint32
	// ForcedMark - Count of packets forcibly marked, after the maximum threshold.
	ForcedMark uint32
	// ProbMarkHead - Count of packets marked early at the head of the queue, below the maximum threshold.
	ProbMarkHead uint32
	// ForcedMarkHead - Count of packets forcibly marked at the head of the queue, after the maximum threshold.
	ForcedMarkHead uint32
}
