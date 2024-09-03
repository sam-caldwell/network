package core

// CtAttrProtoInfoTcpEnum (ctattr_protoinfo_tcp)
//
// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/nfnetlink_conntrack.h
type CtAttrProtoInfoTcpEnum uint8

const (
	// CtaProtoinfoTcpUnspec - CTA_PROTOINFO_TCP_UNSPEC - indicates that specific TCP protocol information is
	// unspecified or not applicable in connection tracking.
	CtaProtoinfoTcpUnspec CtAttrProtoInfoTcpEnum = 0

	// CtaProtoinfoTcpState - CTA_PROTOINFO_TCP_STATE - represent the state of a TCP connection within the
	// connection tracking system.
	CtaProtoinfoTcpState CtAttrProtoInfoTcpEnum = 1

	// CtaProtoinfoTcpWscaleOriginal - CTA_PROTOINFO_TCP_WSCALE_ORIGINAL - represent the original window scale factor
	// of a TCP connection as it was negotiated during the connection establishment phase, within the connection tracking system.
	CtaProtoinfoTcpWscaleOriginal CtAttrProtoInfoTcpEnum = 2

	// CtaProtoinfoTcpWscaleReply - CTA_PROTOINFO_TCP_WSCALE_REPLY - represent the window scale factor in the reply
	// direction of a TCP connection, as it was negotiated during the connection establishment phase, within the
	// connection tracking system.
	CtaProtoinfoTcpWscaleReply CtAttrProtoInfoTcpEnum = 3

	// CtaProtoinfoTcpFlagsOriginal - CTA_PROTOINFO_TCP_FLAGS_ORIGINAL - represent the original TCP flags (such as
	// SYN, ACK, etc.) that were set in the initial packet of a TCP connection, within the connection tracking system.
	CtaProtoinfoTcpFlagsOriginal CtAttrProtoInfoTcpEnum = 4

	// CtaProtoinfoTcpFlagsReply - CTA_PROTOINFO_TCP_FLAGS_REPLY - represent the TCP flags in the reply direction of
	// a connection, as observed in the connection tracking system.
	CtaProtoinfoTcpFlagsReply CtAttrProtoInfoTcpEnum = 5
)
