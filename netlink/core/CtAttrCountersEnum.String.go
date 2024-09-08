package core

// String - Return string representation of the counter state
func (counter *CtAttrCountersEnum) String() string {
	switch *counter {
	case CtaCountersPackets:
		return "CTA_COUNTERS_PACKETS"
	case CtaCountersBytes:
		return "CTA_COUNTERS_BYTES"
	case CtaCounters32Packets:
		return "CTA_COUNTERS32_PACKETS"
	case CtaCounters32Bytes:
		return "CTA_COUNTERS32_BYTES"
	case CtaCountersPad:
		return "CTA_COUNTERS_PAD"
	default:
		panic("unhandled default case")
	}
}
