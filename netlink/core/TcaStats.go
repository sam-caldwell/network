package core

// TcaStats - Enumeration for Traffic Control Action (TCA) statistics.
//
// These constants represent various types of statistics used in the Linux kernel's Traffic Control system.
//
// For more information, see: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_sched.h
type TcaStats uint8

const (
	// TcaStatsUnspec - TCA_STATS_UNSPEC - Unspecified statistics type.
	//
	// For more information, see: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_sched.h
	TcaStatsUnspec TcaStats = iota

	// TcaStatsBasic - TCA_STATS_BASIC - Basic statistics.
	//
	// For more information, see: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_sched.h
	TcaStatsBasic TcaStats = iota

	// TcaStatsRateEst - TCA_STATS_RATE_EST - Rate estimation statistics.
	//
	// For more information, see: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_sched.h
	TcaStatsRateEst TcaStats = iota

	// TcaStatsQueue - TCA_STATS_QUEUE - Queue statistics.
	//
	// For more information, see: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_sched.h
	TcaStatsQueue TcaStats = iota

	// TcaStatsApp - TCA_STATS_APP - Application-specific statistics.
	//
	// For more information, see: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_sched.h
	TcaStatsApp TcaStats = iota

	// TcaStatsRateEst64 - TCA_STATS_RATE_EST64 - 64-bit rate estimation statistics.
	//
	// For more information, see: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_sched.h
	TcaStatsRateEst64 TcaStats = iota

	// TcaStatsPad - TCA_STATS_PAD - Padding for alignment.
	//
	// For more information, see: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_sched.h
	TcaStatsPad TcaStats = iota

	// TcaStatsBasicHw - TCA_STATS_BASIC_HW - Basic hardware statistics.
	//
	// For more information, see: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_sched.h
	TcaStatsBasicHw TcaStats = iota

	// TcaStatsPkt64 - TCA_STATS_PKT64 - 64-bit packet statistics.
	//
	// For more information, see: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_sched.h
	TcaStatsPkt64 TcaStats = iota

	// TcaStatsMax - TCA_STATS_MAX - Maximum valid statistics type.
	//
	// For more information, see: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_sched.h
	TcaStatsMax = iota - 1
)
