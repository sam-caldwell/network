package core

// number of nested RTATTR
// from include/uapi/linux/seg6_iptunnel.h
const (
	SEG6_IPTUNNEL_UNSPEC = iota
	SEG6_IPTUNNEL_SRH
	__SEG6_IPTUNNEL_MAX
)
const (
	SEG6_IPTUNNEL_MAX = __SEG6_IPTUNNEL_MAX - 1
)
