package core

// TcaHfscEnum - Enumeration of HFSC (Hierarchical Fair Service Curve) qdisc parameters.
// HFSC is a Linux qdisc used for advanced traffic shaping, allowing hierarchical bandwidth sharing.
//
// For more details, refer to:
// - Linux Kernel Source: https://github.com/torvalds/linux/blob/master/include/uapi/linux/pkt_sched.h
type TcaHfscEnum uint8

const (
	// TcaHfscUnspec - TCA_HFSC_UNSPEC - Unspecified parameter.
	// This is a placeholder for undefined or unspecified values in HFSC qdisc configuration.
	TcaHfscUnspec TcaHfscEnum = iota

	// TcaHfscRsc - TCA_HFSC_RSC - Realtime Service Curve (RSC) configuration.
	// Defines the minimum bandwidth guaranteed for real-time traffic in the HFSC hierarchy.
	TcaHfscRsc

	// TcaHfscFsc - TCA_HFSC_FSC - Fair Service Curve (FSC) configuration.
	// Provides a fair bandwidth allocation for non-realtime traffic, ensuring a balanced distribution.
	TcaHfscFsc

	// TcaHfscUsc - TCA_HFSC_USC - Upper Service Curve (USC) configuration.
	// Sets the upper bound on bandwidth usage for specific classes within the HFSC hierarchy.
	TcaHfscUsc
)
