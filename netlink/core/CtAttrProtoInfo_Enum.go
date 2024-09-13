package core

// CtAttrProtoInfo - ctattr_protoinfo is a structure in the Linux Netfilter framework that holds protocol-specific
// information for a tracked connection within the connection tracking system.
//
// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/nfnetlink_conntrack.h
type CtAttrProtoInfo uint8

const (
	// CtaProtoInfoUnspec - CTA_PROTOINFO_UNSPEC - indicate that no specific protocol information is provided or
	// applicable in the connection tracking system.
	CtaProtoInfoUnspec CtAttrProtoInfo = 0

	// CtaProtoInfoTcp - CTA_PROTOINFO_TCP - specify that the protocol-specific information pertains to a TCP
	// connection within the connection tracking system.
	CtaProtoInfoTcp CtAttrProtoInfo = 1

	// CtaProtoInfoDccp - CTA_PROTOINFO_DCCP - indicate that the protocol-specific information pertains to a
	// DCCP (Datagram Congestion Control Protocol)
	CtaProtoInfoDccp CtAttrProtoInfo = 2

	// CtaProtoInfoSctp - CTA_PROTOINFO_SCTP -  indicate that the protocol-specific information pertains to an
	// SCTP (Stream Control Transmission Protocol) connection within the connection tracking system.
	CtaProtoInfoSctp CtAttrProtoInfo = 3
)
