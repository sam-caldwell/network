package core

// TcaFqEnum - Enumeration of fq (Fair Queuing) qdisc attributes in Linux Traffic Control.
// These attributes allow you to configure the behavior of the fq qdisc, which is designed to minimize queue delays
// and distribute bandwidth fairly among multiple flows.
//
// For more details, refer to the Linux kernel source:
// - https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_sched.h
type TcaFqEnum uint8

const (
	// TcaFqUnspec - TCA_FQ_UNSPEC - Unspecified attribute.
	TcaFqUnspec TcaFqEnum = iota

	// TcaFqPlimit - TCA_FQ_PLIMIT - Total packet limit for the queue.
	// This defines the maximum number of packets allowed in the fq qdisc at any given time.
	TcaFqPlimit

	// TcaFqFlowPlimit - TCA_FQ_FLOW_PLIMIT - Packet limit per flow.
	// This sets the maximum number of packets allowed per flow (an individual traffic stream).
	TcaFqFlowPlimit

	// TcaFqQuantum - TCA_FQ_QUANTUM - Round-Robin (RR) quantum.
	// The number of bytes a flow can send in one round of the round-robin scheduling.
	TcaFqQuantum

	// TcaFqInitialQuantum - TCA_FQ_INITIAL_QUANTUM - Initial quantum for a new flow.
	// This value sets the initial quantum for new flows before they receive their regular quantum value.
	TcaFqInitialQuantum

	// TcaFqRateEnable - TCA_FQ_RATE_ENABLE - Enable or disable rate limiting.
	// Controls whether per-flow rate limiting is enabled or disabled.
	TcaFqRateEnable

	// TcaFqFlowDefaultRate - TCA_FQ_FLOW_DEFAULT_RATE - Obsolete, do not use.
	// This field was previously used but is now deprecated.
	TcaFqFlowDefaultRate

	// TcaFqFlowMaxRate - TCA_FQ_FLOW_MAX_RATE - Maximum rate per flow.
	// This sets the maximum bandwidth a flow is allowed to consume.
	TcaFqFlowMaxRate

	// TcaFqBucketsLog - TCA_FQ_BUCKETS_LOG - Log2 of the number of hash buckets.
	// This attribute sets the size of the hash table used for classifying flows in logarithmic scale.
	TcaFqBucketsLog

	// TcaFqFlowRefillDelay - TCA_FQ_FLOW_REFILL_DELAY - Delay for flow credit refill (in microseconds).
	// This controls how quickly a flow's credits are refilled.
	TcaFqFlowRefillDelay

	// TcaFqOrphanMask - TCA_FQ_ORPHAN_MASK - Mask applied to orphaned skb (socket buffer) hashes.
	// Orphaned packets are those with no associated user context, and this sets the mask for such packets.
	TcaFqOrphanMask

	// TcaFqLowRateThreshold - TCA_FQ_LOW_RATE_THRESHOLD - Delay threshold for low-rate flows.
	// Packets below this rate experience additional delay to prevent starvation.
	TcaFqLowRateThreshold

	// TcaFqCeThreshold - TCA_FQ_CE_THRESHOLD - Congestion Experienced (CE) marking threshold.
	// This sets the threshold for marking packets with CE (Congestion Experienced) to signal congestion.
	TcaFqCeThreshold

	// TcaFqTimerSlack - TCA_FQ_TIMER_SLACK - Slack time for scheduling timers.
	// This attribute defines how much slack is allowed in timer precision.
	TcaFqTimerSlack

	// TcaFqHorizon - TCA_FQ_HORIZON - Time horizon for scheduling (in microseconds).
	// This controls how far into the future packets can be scheduled.
	TcaFqHorizon

	// TcaFqHorizonDrop - TCA_FQ_HORIZON_DROP - Drop packets beyond the time horizon or cap their Earliest Departure Time (EDT).
	// This controls whether packets are dropped or capped when they exceed the time horizon.
	TcaFqHorizonDrop
)
