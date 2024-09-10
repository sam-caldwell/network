package core

import (
	"net"
	"testing"
)

func TestSrhSeg6Encode(t *testing.T) {
	t.Run("valid segment list", func(t *testing.T) {
		// Define a valid list of IPv6 segments
		segments := []net.IP{
			net.ParseIP("2001:db8::1").To16(),
			net.ParseIP("2001:db8::2").To16(),
		}

		// Expected result: 8-byte header + 16 bytes for each segment
		expectedLength := 8 + len(segments)*16

		// Call the SrhSeg6Encode function
		result, err := SrhSeg6Encode(segments)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		// Check if the result has the expected length
		if len(result) != expectedLength {
			t.Fatalf("Expected length: %d, got: %d", expectedLength, len(result))
		}

		// Validate the header fields
		if result[0] != 0 {
			t.Errorf("Expected nextHdr: 0, got: %d", result[0])
		}
		if result[1] != uint8(16*len(segments)>>3) {
			t.Errorf("Expected hdrLen: %d, got: %d", uint8(16*len(segments)>>3), result[1])
		}
		if result[2] != byte(Ipv6SrcrtType4) {
			t.Errorf("Expected routingType: %d, got: %d", Ipv6SrcrtType4, result[2])
		}
		if result[3] != uint8(len(segments)-1) {
			t.Errorf("Expected segmentsLeft: %d, got: %d", len(segments)-1, result[3])
		}
		if result[4] != uint8(len(segments)-1) {
			t.Errorf("Expected firstSegment: %d, got: %d", len(segments)-1, result[4])
		}
		if result[5] != 0 {
			t.Errorf("Expected flags: 0, got: %d", result[5])
		}

		// Check if the segment addresses are correctly appended
		for i, segment := range segments {
			start := 8 + i*16
			end := start + 16
			if !net.IP(result[start:end]).Equal(segment) {
				t.Errorf("Expected segment %d: %v, got: %v", i, segment, net.IP(result[start:end]))
			}
		}
	})

	t.Run("empty segment list", func(t *testing.T) {
		// Call SrhSeg6Encode with an empty segment list
		_, err := SrhSeg6Encode([]net.IP{})
		if err == nil {
			t.Fatal("Expected error for empty segment list, got nil")
		}

		// Verify the error message
		expectedErr := "SrhSeg6Encode: No Segments"
		if err.Error() != expectedErr {
			t.Errorf("Expected error: %q, got: %q", expectedErr, err.Error())
		}
	})
}
