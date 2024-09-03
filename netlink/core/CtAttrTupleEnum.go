package core

// CtAttrTupleEnum - ctattr_tuple - enumerated type in the Linux Netfilter framework that holds the tuple information
// for a tracked connection, including attributes like source and destination IP addresses, ports, and protocol
// numbers, which uniquely identify the connection within the connection tracking system.
//
// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/netfilter/nfnetlink_conntrack.h
type CtAttrTupleEnum uint8

const (

	// CtaTupleUnspec - CTA_TUPLE_UNSPEC - constant in the Linux Netfilter framework that represents an unspecified
	// or uninitialized tuple attribute within the connection tracking system, often used as a placeholder or
	// default value.
	CtaTupleUnspec CtAttrTupleEnum = 0

	// CtaTupleIp - CTA_TUPLE_IP - constant in the Linux Netfilter framework that represents the IP layer tuple,
	// including source and destination IP addresses, for a tracked connection within the connection tracking system.
	CtaTupleIp CtAttrTupleEnum = 1

	// CtaTupleProto - CTA_TUPLE_PROTO - constant in the Linux Netfilter framework that represents the
	// protocol-specific part of a tuple, including attributes like source and destination ports, and protocol number,
	// for a tracked connection within the connection tracking system.
	CtaTupleProto CtAttrTupleEnum = 2

	// CtaTupleZone - CTA_TUPLE_ZONE - constant in the Linux Netfilter framework that represents the zone information
	// associated with a tuple, used to segregate connections into different security or administrative zones within
	// the connection tracking system.
	CtaTupleZone CtAttrTupleEnum = 3

	// CtaTupleMax - CTA_TUPLE_MAX - constant in the Linux Netfilter framework that represents the maximum value for
	// tuple-related attributes, indicating the upper limit of tuple types defined within the connection tracking
	// system.
	CtaTupleMax CtAttrTupleEnum = CtaTupleZone
)
