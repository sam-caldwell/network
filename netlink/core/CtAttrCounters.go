package core

// CtAttrCounters - https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/nfnetlink_conntrack.h
type CtAttrCounters int

const (
	CTA_COUNTERS_PACKETS   CtAttrCounters = 1 /* 64bit counters */
	CTA_COUNTERS_BYTES     CtAttrCounters = 2 /* 64bit counters */
	CTA_COUNTERS32_PACKETS CtAttrCounters = 3 /* old 32bit counters, unused */
	CTA_COUNTERS32_BYTES   CtAttrCounters = 4 /* old 32bit counters, unused */
	CTA_COUNTERS_PAD       CtAttrCounters = 5
	CTA_COUNTERS_MAX                      = CTA_COUNTERS_PAD
)
