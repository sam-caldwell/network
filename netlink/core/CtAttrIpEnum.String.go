package core

// String - Return string representation of the counter state
func (counter *CtAttrIpEnum) String() string {
	switch *counter {
	case CtaIpUnspec:
		return "CTA_IP_UNSPEC"
	case CtaIpV4Src:
		return "CTA_IP_V4_SRC"
	case CtaIpV4Dst:
		return "CTA_IP_V4_DST"
	case CtaIpV6Src:
		return "CTA_IP_V6_SRC"
	case CtaIpV6Dst:
		return "CTA_IP_V6_DST"
	default:
		panic("unhandled default case")
	}
}
