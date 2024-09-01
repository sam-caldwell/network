package core

// CtAttrCounters - https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/nfnetlink_conntrack.h
type CtAttrCounters int

const (
	// CtaCountersPackets - CTA_COUNTERS_PACKETS
	CtaCountersPackets CtAttrCounters = 1 /* 64bit counters */

	// CtaCountersBytes - CTA_COUNTERS_BYTES
	CtaCountersBytes CtAttrCounters = 2 /* 64bit counters */

	// CtaCounters32Packets - CTA_COUNTERS32_PACKETS
	CtaCounters32Packets CtAttrCounters = 3 /* old 32bit counters, unused */

	// CtaCounters32Bytes - CTA_COUNTERS32_BYTES
	CtaCounters32Bytes CtAttrCounters = 4 /* old 32bit counters, unused */

	// CtaCountersPad - CTA_COUNTERS_PAD
	CtaCountersPad CtAttrCounters = 5

	// CtaCountersMax - CTA_COUNTERS_MAX
	CtaCountersMax = CtaCountersPad
)
