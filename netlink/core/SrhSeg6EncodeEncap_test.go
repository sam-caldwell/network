package core

import (
	"net"
	"testing"
)

func TestSrhSeg6EncodeEncap(t *testing.T) {
	t.Run("valid segment list", func(t *testing.T) {
		// Define a valid list of IPv6 segments
		segments := []net.IP{
			net.ParseIP("2001:db8::1").To16(),
			net.ParseIP("2001:db8::2").To16(),
		}

		// Expected result: 12-byte header + 16 bytes for each segment
		expectedLength := 12 + len(segments)*16
		mode := 1 // Sample mode value

		// Call the SrhSeg6EncodeEncap function
		result, err := SrhSeg6EncodeEncap(mode, segments)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		// Check if the result has the expected length
		if len(result) != expectedLength {
			t.Fatalf("Expected length: %d, got: %d", expectedLength, len(result))
		}

		// Validate the encoded mode
		expectedMode := uint32(mode)
		encodedMode := NativeEndian.Uint32(result[0:4])
		if encodedMode != expectedMode {
			t.Errorf("Expected mode: %d, got: %d", expectedMode, encodedMode)
		}

		// Validate the SRH header fields
		if result[4] != 0 {
			t.Errorf("Expected nextHdr: 0, got: %d", result[4])
		}
		if result[5] != uint8(16*len(segments)>>3) {
			t.Errorf("Expected hdrLen: %d, got: %d", uint8(16*len(segments)>>3), result[5])
		}
		if result[6] != byte(Ipv6SrcrtType4) {
			t.Errorf("Expected routingType: %d, got: %d", Ipv6SrcrtType4, result[6])
		}
		if result[7] != uint8(len(segments)-1) {
			t.Errorf("Expected segmentsLeft: %d, got: %d", len(segments)-1, result[7])
		}
		if result[8] != uint8(len(segments)-1) {
			t.Errorf("Expected firstSegment: %d, got: %d", len(segments)-1, result[8])
		}
		if result[9] != 0 {
			t.Errorf("Expected flags: 0, got: %d", result[9])
		}

		// Check if the segment addresses are correctly appended
		for i, segment := range segments {
			start := 12 + i*16
			end := start + 16
			if !net.IP(result[start:end]).Equal(segment) {
				t.Errorf("Expected segment %d: %v, got: %v", i, segment, net.IP(result[start:end]))
			}
		}
	})

	t.Run("empty segment list", func(t *testing.T) {
		// Call SrhSeg6EncodeEncap with an empty segment list
		mode := 1
		_, err := SrhSeg6EncodeEncap(mode, []net.IP{})
		if err == nil {
			t.Fatal("Expected error for empty segment list, got nil")
		}

		// Verify the error message
		expectedErr := "SrhSeg6EncodeEncap: No Segment in srh"
		if err.Error() != expectedErr {
			t.Errorf("Expected error: %q, got: %q", expectedErr, err.Error())
		}
	})
}
