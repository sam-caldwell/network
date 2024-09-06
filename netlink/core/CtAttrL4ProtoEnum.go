package core

// CtAttrL4ProtoEnum - ctattr_l4proto - structure in the Linux Netfilter framework that stores layer 4 protocol-specific
// attributes for a tracked connection, such as protocol-specific information related to TCP, IpProtoUDP, SCTP, or DCCP,
// within the connection tracking system.
//
// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/nfnetlink_conntrack.h
type CtAttrL4ProtoEnum uint8

const (

	// CtaProtoNum - CTA_PROTO_NUM - specify the layer 4 protocol number (e.g., TCP, IpProtoUDP, IpProtoIcmp) for a tracked
	// connection within the connection tracking system.
	CtaProtoNum CtAttrL4ProtoEnum = 1

	// CtaProtoSrcPort - CTA_PROTO_SRC_PORT - represents the source port number of a layer 4 protocol (such as TCP
	// or IpProtoUDP) for a tracked connection within the connection tracking system.
	CtaProtoSrcPort CtAttrL4ProtoEnum = 2

	// CtaProtoDstPort - CTA_PROTO_DST_PORT - represents the destination port number of a layer 4 protocol (such as
	// TCP or IpProtoUDP) for a tracked connection within the connection tracking system.
	CtaProtoDstPort CtAttrL4ProtoEnum = 3
)
