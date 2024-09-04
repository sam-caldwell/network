package core

// TcaPoliceEnum - Enum for policing attributes in traffic control actions.
//
// These constants represent various traffic policing attributes in Linux, such as token bucket filter (TBF),
// rate limits, and burst size for traffic shaping. These enums are used in the Linux kernel for setting
// up traffic policing rules via netlink.
//
// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_cls.h
type TcaPoliceEnum uint8

const (
	// TcaPoliceUnspec - TCA_POLICE_UNSPEC - unspecified police attribute
	TcaPoliceUnspec TcaPoliceEnum = iota

	// TcaPoliceTbf - TCA_POLICE_TBF - token bucket filter for rate-limiting
	TcaPoliceTbf

	// TcaPoliceRate - TCA_POLICE_RATE - rate limit for the filter
	TcaPoliceRate

	// TcaPolicePeakRate - TCA_POLICE_PEAKRATE - peak rate for traffic
	TcaPolicePeakRate

	// TcaPoliceAvrate - TCA_POLICE_AVRATE - average rate for traffic
	TcaPoliceAvrate

	// TcaPoliceResult - TCA_POLICE_RESULT - result for the filter action
	TcaPoliceResult

	// TcaPoliceTm - TCA_POLICE_TM - timing for the police filter
	TcaPoliceTm

	// TcaPolicePad - TCA_POLICE_PAD - padding for alignment
	TcaPolicePad

	// TcaPoliceRate64 - TCA_POLICE_RATE64 - 64-bit rate for large-scale limits
	TcaPoliceRate64

	// TcaPolicePeakRate64 - TCA_POLICE_PEAKRATE64 - 64-bit peak rate
	TcaPolicePeakRate64

	// TcaPolicePktRate64 - TCA_POLICE_PKTRATE64 - 64-bit packet rate
	TcaPolicePktRate64

	// TcaPolicePktBurst64 - TCA_POLICE_PKTBURST64 - 64-bit packet burst rate
	TcaPolicePktBurst64

	// TcaPoliceMax - TCA_POLICE_MAX - maximum police attribute value
	TcaPoliceMax = iota - 1
)
