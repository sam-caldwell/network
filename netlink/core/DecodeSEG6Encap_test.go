package core

import (
	"bytes"
	"encoding/binary"
	"net"
	"testing"
)

func TestDecodeSEG6Encap(t *testing.T) {
	// Subtest 1: Happy path with valid SRH and IPv6 segments
	t.Run("valid SRH with two segments", func(t *testing.T) {
		// Prepare a sample byte buffer for the test.
		// This represents a Segment Routing Header (SRH) with 2 IPv6 segments.
		// Mode is set to 1, followed by 12-byte SRH and 2 segments (32 bytes total for 2 IPv6 addresses).
		buf := make([]byte, 44)

		// Set the mode (first 4 bytes, set to 1)
		binary.BigEndian.PutUint32(buf[0:4], 1)

		// SRH: Fill with dummy values
		buf[4] = 0x3B                                  // Next header
		buf[5] = 0x04                                  // Header length (4 * 8 = 32 bytes, two IPv6 segments)
		buf[6] = 0x04                                  // Routing type (SRH)
		buf[7] = 0x02                                  // Segments Left
		buf[8] = 0x02                                  // First Segment
		buf[9] = 0x00                                  // Flags
		binary.BigEndian.PutUint16(buf[10:12], 0x0000) // Reserved

		// Add two IPv6 segments (32 bytes total, each 16 bytes)
		ip1 := net.ParseIP("2001:0db8::1").To16()
		ip2 := net.ParseIP("2001:0db8::2").To16()

		copy(buf[12:28], ip1)
		copy(buf[28:44], ip2)

		// Call DecodeSEG6Encap with the buffer
		mode, segments, err := DecodeSEG6Encap(buf)

		// Check for no errors
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		// Check that the mode is correctly parsed
		if mode != 1 {
			t.Errorf("Expected mode to be 1, but got %d", mode)
		}

		// Check that two segments were decoded
		if len(segments) != 2 {
			t.Errorf("Expected 2 segments, but got %d", len(segments))
		}

		// Check that the segments match the expected values
		if !bytes.Equal(segments[0], ip1) {
			t.Errorf("Expected first segment to be %v, but got %v", ip1, segments[0])
		}

		if !bytes.Equal(segments[1], ip2) {
			t.Errorf("Expected second segment to be %v, but got %v", ip2, segments[1])
		}
	})

	// Subtest 2: Sad path - Invalid buffer length (not a multiple of 16 for IPv6 segments)
	t.Run("invalid buffer length", func(t *testing.T) {
		// Create a buffer with incorrect length (not a multiple of 16 for IPv6 addresses)
		buf := make([]byte, 30) // should be at least 16 bytes per IPv6 address

		// Set mode and SRH (irrelevant because of invalid segment length)
		binary.BigEndian.PutUint32(buf[0:4], 1)
		buf[4] = 0x3B                                  // Next header
		buf[5] = 0x04                                  // Header length
		buf[6] = 0x04                                  // Routing type (SRH)
		buf[7] = 0x01                                  // Segments Left
		buf[8] = 0x01                                  // First Segment
		buf[9] = 0x00                                  // Flags
		binary.BigEndian.PutUint16(buf[10:12], 0x0000) // Reserved

		// Call DecodeSEG6Encap with the buffer
		_, _, err := DecodeSEG6Encap(buf)

		// Expect an error due to the invalid buffer length
		if err == nil {
			t.Fatalf("Expected error due to invalid buffer length, but got none")
		}
	})
}
