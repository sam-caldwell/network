package core

import (
	"net"
)

// IPv6SrHdr - Represents an IPv6 Segment Routing Header (SRH)
//
// The SRH is an IPv6 extension header used for source routing. It allows the sender to specify a list of intermediate
// nodes (segments) that the packet must traverse before reaching its final destination.
//
// This structure includes fields such as the next header, routing type, flags, and an array of segments representing
// the intermediate nodes.
//
// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/seg6.h
type IPv6SrHdr struct {
	nextHdr      uint8  // Next header type (e.g., TCP, UDP, etc.)
	hdrLen       uint8  // Length of the SRH
	routingType  uint8  // Routing type for Segment Routing
	segmentsLeft uint8  // Number of segments left to visit
	firstSegment uint8  // Index of the first segment
	flags        uint8  // Flags (e.g., O-bit)
	reserved     uint16 // Reserved field for future use

	Segments []net.IP // List of segments (intermediate node IP addresses)
}
