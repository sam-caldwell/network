package core

import (
	"errors"
	"net"
)

// SrhSeg6Encode - Encodes an IPv6 Segment Routing Header (SRH) into a byte slice.
// This function constructs the SRH and appends the provided segment list (IPv6 addresses) to the header.
//
// Parameters:
//   - segments: A list of net.IP addresses representing the segment list in the SRH.
//
// Returns:
//   - []byte: The encoded SRH as a byte slice.
//   - error: An error if the segment list is empty.
func SrhSeg6Encode(segments []net.IP) ([]byte, error) {
	nsegs := len(segments) // number of segments
	if nsegs == 0 {
		return nil, errors.New("SrhSeg6Encode: No Segments")
	}

	// Create a byte slice with space for the SRH header (8 bytes) and segments (16 bytes each)
	b := make([]byte, 8, 8+len(segments)*16)

	// Populate the SRH fields
	b[0] = 0                         // srh.nextHdr (0 when calling netlink)
	b[1] = uint8(16 * nsegs >> 3)    // srh.hdrLen (in 8-octet units)
	b[2] = byte(Ipv6SrcrtType4)      // srh.routingType (assigned by IANA)
	b[3] = uint8(nsegs - 1)          // srh.segmentsLeft
	b[4] = uint8(nsegs - 1)          // srh.firstSegment
	b[5] = 0                         // srh.flags (set to 0)
	NativeEndian.PutUint16(b[6:], 0) // srh.reserved (set to 0)

	// Append each segment (16-byte IPv6 address) to the byte slice
	for _, netIP := range segments {
		b = append(b, netIP...)
	}

	return b, nil
}
