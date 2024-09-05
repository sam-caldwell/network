package core

// TcaFqCodelEnum - Enumeration of fq_codel qdisc (queue discipline) attributes in Linux traffic control.
//
// fq_codel (Fair Queuing Controlled Delay) is a queue management algorithm designed to minimize latency in packet queues,
// while maintaining fairness in the distribution of bandwidth. These attributes are used to control its behavior.
//
// For more details, refer to the Linux kernel source:
// - https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_sched.h
type TcaFqCodelEnum uint8

const (
	// TcaFqCodelUnspec - TCA_FQ_CODEL_UNSPEC - Unspecified fq_codel attribute.
	TcaFqCodelUnspec TcaFqCodelEnum = iota

	// TcaFqCodelTarget - TCA_FQ_CODEL_TARGET - The target delay for packets in the queue.
	TcaFqCodelTarget

	// TcaFqCodelLimit - TCA_FQ_CODEL_LIMIT - Maximum number of packets allowed in the queue.
	TcaFqCodelLimit

	// TcaFqCodelInterval - TCA_FQ_CODEL_INTERVAL - Time interval used to judge delay.
	TcaFqCodelInterval

	// TcaFqCodelEcn - TCA_FQ_CODEL_ECN - ECN (Explicit Congestion Notification) marking configuration.
	TcaFqCodelEcn

	// TcaFqCodelFlows - TCA_FQ_CODEL_FLOWS - Number of queues (flows) in fq_codel.
	TcaFqCodelFlows

	// TcaFqCodelQuantum - TCA_FQ_CODEL_QUANTUM - The number of bytes sent in one round per queue.
	TcaFqCodelQuantum

	// TcaFqCodelCeThreshold - TCA_FQ_CODEL_CE_THRESHOLD - Threshold for setting ECN CE (Congestion Experienced) mark.
	TcaFqCodelCeThreshold

	// TcaFqCodelDropBatchSize - TCA_FQ_CODEL_DROP_BATCH_SIZE - Number of packets dropped in one batch.
	TcaFqCodelDropBatchSize

	// TcaFqCodelMemoryLimit - TCA_FQ_CODEL_MEMORY_LIMIT - Memory limit for the fq_codel qdisc.
	TcaFqCodelMemoryLimit
)
