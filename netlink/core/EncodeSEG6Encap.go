package core

import (
	"errors"
	"net"
)

// EncodeSEG6Encap - Encodes a Segment Routing Header (SRH) for IPv6 encapsulation.
//
// This function encodes the SRH with a given mode and a list of IPv6 segments.
// The segments represent a set of intermediate nodes the packet should traverse.
// The function returns the byte-encoded SRH or an error if no segments are provided.
//
// Parameters:
//   - mode: The mode used for encoding (e.g., routing mode).
//   - segments: A list of net.IP representing the routing segments.
//
// See: https://github.com/torvalds/linux/blob/master/include/uapi/linux/seg6.h
func EncodeSEG6Encap(mode int, segments []net.IP) ([]byte, error) {
	nsegs := len(segments) // Number of segments

	if nsegs == 0 {
		return nil, errors.New("EncodeSEG6Encap: No Segment in srh")
	}

	// Create a byte slice with initial length 12 (fixed header) and capacity for the segments
	b := make([]byte, 12, 12+len(segments)*16)

	// Encode mode in the first 4 bytes
	NativeEndian.PutUint32(b, uint32(mode))

	// Populate the SRH fields
	b[4] = 0                          // srh.nextHdr: Placeholder
	b[5] = uint8(16 * nsegs >> 3)     // srh.hdrLen: Length of SRH in 8-octet units
	b[6] = Ipv6SrcrtType4             // srh.routingType: Segment Routing Type (IANA-assigned)
	b[7] = uint8(nsegs - 1)           // srh.segmentsLeft: Segments left to process
	b[8] = uint8(nsegs - 1)           // srh.firstSegment: Index of the first segment
	b[9] = 0                          // srh.flags: Flags (e.g., HMAC)
	NativeEndian.PutUint16(b[10:], 0) // srh.reserved: Reserved field (Tag in some drafts)

	// Append each segment (16 bytes per IPv6 address) to the byte slice
	for _, netIP := range segments {
		b = append(b, netIP...)
	}

	return b, nil
}
