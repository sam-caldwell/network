package core

// TcSfqQopt - This structure represents the options for the SFQ (Stochastic Fairness Queueing) scheduler in Linux Traffic Control.
// SFQ is a simple stochastic queueing algorithm that allocates bandwidth equally across flows without maintaining per-flow state.
// For more details on SFQ, refer to:
// - https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_sched.h
// - https://man7.org/linux/man-pages/man8/tc-sfq.8.html
type TcSfqQopt struct {
	// Quantum - Bytes per round allocated to flow. Determines how much data each flow gets to send per round.
	// The higher the quantum, the more data per round a flow can transmit.
	Quantum uint8

	// Perturb - Period of hash perturbation. The hash table used to assign packets to flows is perturbed periodically
	// to avoid any long-term bias.
	Perturb int32

	// Limit - Maximum packets allowed in the queue. This field controls the maximum number of packets that can be queued
	// before packets start being dropped.
	Limit uint32

	// Divisor - Hash divisor used for determining the number of hash buckets. A smaller divisor means fewer buckets,
	// which leads to fewer flows being tracked, but can increase hash collisions.
	Divisor uint8

	// Flows - Maximum number of flows. Controls the total number of simultaneous flows that can be managed by the SFQ algorithm.
	Flows uint8
}
