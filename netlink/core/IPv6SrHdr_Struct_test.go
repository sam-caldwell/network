package core

import (
	"net"
	"testing"
	"unsafe"
)

func TestIPv6SrHdr(t *testing.T) {
	t.Run("TestSize", func(t *testing.T) {
		// 6 bytes for fixed fields + 2 bytes for reserved (total 8 bytes) + 24 bytes overhead for []net.IP
		const expectedSizeInBytes = 6 + 2 + unsafe.Sizeof([]net.IP{})
		var srh IPv6SrHdr

		if size := unsafe.Sizeof(srh); size != uintptr(expectedSizeInBytes) {
			t.Errorf("Expected size %d but got %d", expectedSizeInBytes, size)
		}
	})

	t.Run("TestValues", func(t *testing.T) {
		segments := []net.IP{
			net.ParseIP("2001:db8::1"),
			net.ParseIP("2001:db8::2"),
		}

		// Initializing the structure
		srh := IPv6SrHdr{
			nextHdr:      6,   // e.g., TCP (6)
			hdrLen:       12,  // Header length
			routingType:  4,   // Segment Routing
			segmentsLeft: 2,   // Two segments to visit
			firstSegment: 1,   // Starting at segment 1
			flags:        0x0, // No flags set
			reserved:     0,   // Reserved field
			Segments:     segments,
		}

		t.Run("TestNextHdr", func(t *testing.T) {
			expected := uint8(6)
			if srh.nextHdr != expected {
				t.Errorf("Expected nextHdr %d but got %d", expected, srh.nextHdr)
			}
		})

		t.Run("TestHdrLen", func(t *testing.T) {
			expected := uint8(12)
			if srh.hdrLen != expected {
				t.Errorf("Expected hdrLen %d but got %d", expected, srh.hdrLen)
			}
		})

		t.Run("TestRoutingType", func(t *testing.T) {
			expected := uint8(4)
			if srh.routingType != expected {
				t.Errorf("Expected routingType %d but got %d", expected, srh.routingType)
			}
		})

		t.Run("TestSegmentsLeft", func(t *testing.T) {
			expected := uint8(2)
			if srh.segmentsLeft != expected {
				t.Errorf("Expected segmentsLeft %d but got %d", expected, srh.segmentsLeft)
			}
		})

		t.Run("TestFirstSegment", func(t *testing.T) {
			expected := uint8(1)
			if srh.firstSegment != expected {
				t.Errorf("Expected firstSegment %d but got %d", expected, srh.firstSegment)
			}
		})

		t.Run("TestFlags", func(t *testing.T) {
			expected := uint8(0x0)
			if srh.flags != expected {
				t.Errorf("Expected flags %d but got %d", expected, srh.flags)
			}
		})

		t.Run("TestReserved", func(t *testing.T) {
			expected := uint16(0)
			if srh.reserved != expected {
				t.Errorf("Expected reserved %d but got %d", expected, srh.reserved)
			}
		})

		t.Run("TestSegments", func(t *testing.T) {
			expectedSegments := []net.IP{
				net.ParseIP("2001:db8::1"),
				net.ParseIP("2001:db8::2"),
			}

			for i, segment := range srh.Segments {
				if !segment.Equal(expectedSegments[i]) {
					t.Errorf("Expected segment %s but got %s", expectedSegments[i], segment)
				}
			}
		})
	})
}
