package core

// String - Return string representation of the counter state
func (counter *CtAttrCountersEnum) String() string {
	switch *counter {
	case CtaCountersUnspec:
		return "CtaCountersUnspec"
	case CtaCountersPackets:
		return "CtaCountersPackets"
	case CtaCountersBytes:
		return "CtaCountersBytes"
	case CtaCounters32Packets:
		return "CtaCounters32Packets"
	case CtaCounters32Bytes:
		return "CtaCounters32Bytes"
	case CtaCountersPad:
		return "CtaCountersPad"
	default:
		panic("unhandled default case")
	}
}
