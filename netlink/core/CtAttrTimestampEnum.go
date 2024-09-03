package core

// CtAttrTimestampEnum - Enumeration for Connection Tracking (Conntrack) Timestamp Attributes
//
// This enumeration defines the attributes related to timestamps in the Connection Tracking (Conntrack) subsystem.
// These attributes represent specific timestamp values used in tracking the start and stop times of network connections.
//
// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/nfnetlink_conntrack.h
type CtAttrTimestampEnum uint64

const (
	// CtaTimestampStart - CTA_TIMESTAMP_START -
	// This attribute represents the timestamp indicating when a network connection started.
	// It is used in connection tracking to record the initiation time of a tracked connection.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/nfnetlink_conntrack.h
	CtaTimestampStart CtAttrTimestampEnum = 1

	// CtaTimestampStop - CTA_TIMESTAMP_STOP -
	// This attribute represents the timestamp indicating when a network connection ended.
	// It is used in connection tracking to record the termination time of a tracked connection.
	//
	// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/nfnetlink_conntrack.h
	CtaTimestampStop CtAttrTimestampEnum = 2
)
