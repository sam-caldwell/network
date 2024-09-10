package core

import (
	"fmt"
	"net"
)

// SrhSeg6Decode - Decodes an IPv6 Segment Routing Header (SRH) from a byte buffer.
//
// This function extracts an SRH from a byte slice and parses the segments to return a list of net.IP addresses.
// SRH is used in segment routing for IPv6 to define a list of waypoints for a packet to travel through.
//
// Parameters:
//   - buf: A byte slice containing the SRH.
//
// Returns:
//   - segments: A slice of net.IP representing the segment list in the SRH.
//   - error: An error if the buffer length is invalid or malformed.
func SrhSeg6Decode(buf []byte) ([]net.IP, error) {
	// Parse the IPv6 SRH header from the buffer
	srh := IPv6SrHdr{
		nextHdr:      buf[0],
		hdrLen:       buf[1],
		routingType:  buf[2],
		segmentsLeft: buf[3],
		firstSegment: buf[4],
		flags:        buf[5],
		reserved:     NativeEndian.Uint16(buf[6:8]),
	}

	// Move buffer forward after the SRH header
	buf = buf[8:]

	// Ensure the remaining buffer length is a multiple of 16 (for IPv6 segment addresses)
	if len(buf)%16 != 0 {
		err := fmt.Errorf("SrhSeg6Decode: error parsing Segment List (buf len: %d)", len(buf))
		return nil, err
	}

	// Extract each 16-byte IPv6 address and add it to the segment list
	for len(buf) > 0 {
		srh.Segments = append(srh.Segments, net.IP(buf[:16]))
		buf = buf[16:]
	}

	// Return the segment list
	return srh.Segments, nil
}
