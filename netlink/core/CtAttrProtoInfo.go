package core

// CtAttrProtoInfo - ctattr_protoinfo is a structure in the Linux Netfilter framework that holds protocol-specific
// information for a tracked connection within the connection tracking system.
//
// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/nfnetlink_conntrack.h
type CtAttrProtoInfo uint8

const (
	// CtaProtoinfoUnspec - CTA_PROTOINFO_UNSPEC - indicate that no specific protocol information is provided or
	// applicable in the connection tracking system.
	CtaProtoinfoUnspec CtAttrProtoInfo = 0

	// CtaProtoinfoTcp - CTA_PROTOINFO_TCP - specify that the protocol-specific information pertains to a TCP
	// connection within the connection tracking system.
	CtaProtoinfoTcp CtAttrProtoInfo = 1

	// CtaProtoinfoDccp - CTA_PROTOINFO_DCCP - indicate that the protocol-specific information pertains to a
	// DCCP (Datagram Congestion Control Protocol)
	CtaProtoinfoDccp CtAttrProtoInfo = 2

	// CtaProtoinfoSctp - CTA_PROTOINFO_SCTP -  indicate that the protocol-specific information pertains to an
	// SCTP (Stream Control Transmission Protocol) connection within the connection tracking system.
	CtaProtoinfoSctp CtAttrProtoInfo = 3
)
