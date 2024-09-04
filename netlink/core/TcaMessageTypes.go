package core

// TcaMessageTypes - Enumeration of Traffic Control Action (TCA) message types.
//
// These message types are used in the Linux kernel's Traffic Control system for various actions and statistics.
// They represent the types of attributes associated with traffic control (qdiscs, classes, filters, etc.).
//
// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_sched.h
type TcaMessageTypes uint8

const (
	// TCAUnspec - TCA_UNSPEC - Unspecified message type.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_sched.h
	TCAUnspec TcaMessageTypes = iota

	// TCAKind - TCA_KIND - Describes the type of traffic control action (e.g., qdisc, filter).
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_sched.h
	TCAKind

	// TCAOptions - TCA_OPTIONS - Options for the specified action.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_sched.h
	TCAOptions

	// TCAStats - TCA_STATS - Basic traffic control statistics.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_sched.h
	TCAStats

	// TCAxStats - TCA_XSTATS - Extended statistics for traffic control.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_sched.h
	TCAxStats

	// TCARate - TCA_RATE - Rate settings for traffic control.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_sched.h
	TCARate

	// TCAFcnt - TCA_FCNT - Filter count.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_sched.h
	TCAFcnt

	// TCAStats2 - TCA_STATS2 - Extended statistics for traffic control, including additional attributes.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_sched.h
	TCAStats2

	// TCAStab - TCA_STAB - Stability settings.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_sched.h
	TCAStab

	// TCAPad - TCA_PAD - Padding for alignment.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_sched.h
	TCAPad

	// TCADumpInvisible - TCA_DUMP_INVISIBLE - Dump invisible elements.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_sched.h
	TCADumpInvisible

	// TCAChain - TCA_CHAIN - Chain information for traffic control actions.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_sched.h
	TCAChain

	// TCAHwOffload - TCA_HW_OFFLOAD - Hardware offloading information.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_sched.h
	TCAHwOffload

	// TCAIngressBlock - TCA_INGRESS_BLOCK - Ingress block for actions.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_sched.h
	TCAIngressBlock

	// TCAEgressBlock - TCA_EGRESS_BLOCK - Egress block for actions.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_sched.h
	TCAEgressBlock

	// TCADumpFlags - TCA_DUMP_FLAGS - Flags for dump operations.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_sched.h
	TCADumpFlags

	// TCAMax - TCA_MAX - Maximum defined TCA message type.
	//
	// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_sched.h
	TCAMax = TCADumpFlags
)
