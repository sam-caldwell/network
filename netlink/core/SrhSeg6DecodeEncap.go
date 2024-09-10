package core

import (
	"encoding/binary"
	"fmt"
	"net"
)

// SrhSeg6DecodeEncap - Decodes a Segment Routing Header (SRH) from a byte buffer.
//
// This function parses the SRH for IPv6 encapsulation, extracting the mode and segments.
// The segments represent a list of intermediate nodes for the routing path.
//
// Parameters:
//   - buf: A byte slice containing the SRH to decode.
//
// Returns:
//   - mode: The mode of encapsulation.
//   - segments: A list of net.IP segments extracted from the SRH.
//   - error: Returns an error if the segment list is malformed or invalid.
func SrhSeg6DecodeEncap(buf []byte) (int, []net.IP, error) {

	// Extract the mode from the first 4 bytes using BigEndian (network byte order)
	mode := int(binary.BigEndian.Uint32(buf))

	// Parse the IPv6 Segment Routing Header (SRH) from the buffer
	srh := IPv6SrHdr{
		nextHdr:      buf[4],
		hdrLen:       buf[5],
		routingType:  buf[6],
		segmentsLeft: buf[7],
		firstSegment: buf[8],
		flags:        buf[9],
		reserved:     binary.BigEndian.Uint16(buf[10:12]),
	}

	// Move the buffer forward, leaving the 12-byte SRH header
	buf = buf[12:]

	// Ensure the remaining buffer length is a multiple of 16 (for IPv6 segment addresses)
	if len(buf)%16 != 0 {
		err := fmt.Errorf("SrhSeg6DecodeEncap: error parsing Segment List (buf len: %d)", len(buf))
		return mode, nil, err
	}

	// Extract each 16-byte IPv6 address and add it to the segment list
	for len(buf) > 0 {
		srh.Segments = append(srh.Segments, net.IP(buf[:16]))
		buf = buf[16:]
	}

	// Return the mode and segment list
	return mode, srh.Segments, nil
}
