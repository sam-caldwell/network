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
	nextHdr      uint8  // Next header type (e.g., TCP, IpProtoUDP, etc.)
	hdrLen       uint8  // Length of the SRH
	routingType  uint8  // Routing type for Segment Routing
	segmentsLeft uint8  // Number of segments left to visit
	firstSegment uint8  // Index of the first segment
	flags        uint8  // Flags (e.g., O-bit)
	reserved     uint16 // Reserved field for future use

	Segments []net.IP // List of segments (intermediate node IP addresses)
}

// Equal - Return whether two IPv6SrHdr structs are equal
func (s1 *IPv6SrHdr) Equal(s2 IPv6SrHdr) bool {

	if len(s1.Segments) != len(s2.Segments) {
		return false
	}

	for i := range s1.Segments {
		if !s1.Segments[i].Equal(s2.Segments[i]) {
			return false
		}
	}

	return s1.nextHdr == s2.nextHdr &&
		s1.hdrLen == s2.hdrLen &&
		s1.routingType == s2.routingType &&
		s1.segmentsLeft == s2.segmentsLeft &&
		s1.firstSegment == s2.firstSegment &&
		s1.flags == s2.flags
	// reserved doesn't need to be identical.
}
