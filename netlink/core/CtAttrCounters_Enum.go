package core

// CtAttrCounters - https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/nfnetlink_conntrack.h
type CtAttrCounters uint8

const (
	// CtaCountersUnspec - unspecified value placeholder
	CtaCountersUnspec CtAttrCounters = 0
	// CtaCountersPackets - CTA_COUNTERS_PACKETS - 64bit counters
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/nfnetlink_conntrack.h
	CtaCountersPackets CtAttrCounters = 1

	// CtaCountersBytes - CTA_COUNTERS_BYTES - 64bit counters
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/nfnetlink_conntrack.h
	CtaCountersBytes CtAttrCounters = 2

	// CtaCounters32Packets - CTA_COUNTERS32_PACKETS - old 32bit counters, unused
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/nfnetlink_conntrack.h
	CtaCounters32Packets CtAttrCounters = 3

	// CtaCounters32Bytes - CTA_COUNTERS32_BYTES - old 32bit counters, unused
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/nfnetlink_conntrack.h
	CtaCounters32Bytes CtAttrCounters = 4

	// CtaCountersPad - CTA_COUNTERS_PAD
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/nfnetlink_conntrack.h
	CtaCountersPad CtAttrCounters = 5

	// CtaCountersMax - CTA_COUNTERS_MAX
	//
	// https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/nfnetlink_conntrack.h
	CtaCountersMax = iota - 1
)
