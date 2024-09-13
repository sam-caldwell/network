package core

import (
	"net"
	"testing"
)

func TestIPv6SrHdrEqual(t *testing.T) {
	tests := []struct {
		name     string
		s1       IPv6SrHdr
		s2       IPv6SrHdr
		expected bool
	}{
		{
			name: "Equal headers",
			s1: IPv6SrHdr{
				nextHdr:      6,
				hdrLen:       8,
				routingType:  4,
				segmentsLeft: 2,
				firstSegment: 0,
				flags:        0,
				Segments:     []net.IP{net.ParseIP("2001:db8::1"), net.ParseIP("2001:db8::2")},
			},
			s2: IPv6SrHdr{
				nextHdr:      6,
				hdrLen:       8,
				routingType:  4,
				segmentsLeft: 2,
				firstSegment: 0,
				flags:        0,
				Segments:     []net.IP{net.ParseIP("2001:db8::1"), net.ParseIP("2001:db8::2")},
			},
			expected: true,
		},
		{
			name: "Different nextHdr",
			s1: IPv6SrHdr{
				nextHdr:      6,
				hdrLen:       8,
				routingType:  4,
				segmentsLeft: 2,
				firstSegment: 0,
				flags:        0,
				Segments:     []net.IP{net.ParseIP("2001:db8::1")},
			},
			s2: IPv6SrHdr{
				nextHdr:      17,
				hdrLen:       8,
				routingType:  4,
				segmentsLeft: 2,
				firstSegment: 0,
				flags:        0,
				Segments:     []net.IP{net.ParseIP("2001:db8::1")},
			},
			expected: false,
		},
		{
			name: "Different segments",
			s1: IPv6SrHdr{
				nextHdr:      6,
				hdrLen:       8,
				routingType:  4,
				segmentsLeft: 2,
				firstSegment: 0,
				flags:        0,
				Segments:     []net.IP{net.ParseIP("2001:db8::1")},
			},
			s2: IPv6SrHdr{
				nextHdr:      6,
				hdrLen:       8,
				routingType:  4,
				segmentsLeft: 2,
				firstSegment: 0,
				flags:        0,
				Segments:     []net.IP{net.ParseIP("2001:db8::2")},
			},
			expected: false,
		},
		{
			name: "Different number of segments",
			s1: IPv6SrHdr{
				nextHdr:      6,
				hdrLen:       8,
				routingType:  4,
				segmentsLeft: 2,
				firstSegment: 0,
				flags:        0,
				Segments:     []net.IP{net.ParseIP("2001:db8::1")},
			},
			s2: IPv6SrHdr{
				nextHdr:      6,
				hdrLen:       8,
				routingType:  4,
				segmentsLeft: 2,
				firstSegment: 0,
				flags:        0,
				Segments:     []net.IP{net.ParseIP("2001:db8::1"), net.ParseIP("2001:db8::2")},
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.s1.Equal(tt.s2)
			if result != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}
