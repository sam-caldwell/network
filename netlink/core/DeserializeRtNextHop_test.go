package core

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestDeserializeRtNextHop(t *testing.T) {
	// Prepare a valid byte slice (32 bytes)
	validTestData := []byte{
		// First 8 bytes for RtNexthop
		0x08, 0x00, // Len (2 bytes, little-endian)
		0x01,                   // Flags
		0x02,                   // Hops
		0x04, 0x03, 0x02, 0x01, // Ifindex (4 bytes, little-endian)
		// Fill the rest with 0s for 32-byte total length (padding or Children)
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	}

	t.Run("input too short", func(t *testing.T) {
		// Prepare an invalid byte slice (less than 32 bytes)
		buf := make([]byte, 6) // Too short

		fmt.Printf("Length of short byte slice: %d bytes\n", len(buf)) // Debugging slice size

		// Call DeserializeRtNextHop
		rtNextHop := DeserializeRtNextHop(buf)

		// Check that the function returns nil for insufficient byte length
		if rtNextHop != nil {
			t.Fatalf("Expected rtNextHop to be nil due to short input, but got %v", rtNextHop)
		}
	})

	t.Run("valid input", func(t *testing.T) {
		var rtNextHop *RtNexthop
		t.Run("setup", func(t *testing.T) {
			rtNextHop = DeserializeRtNextHop(validTestData)
			// Print the actual byte slice length for debugging
			fmt.Printf("Length of valid byte slice: %d bytes\n", len(validTestData))
		})

		t.Run("verify result is not nil", func(t *testing.T) {
			if rtNextHop == nil {
				t.Fatalf("Expected rtNextHop to be non-nil")
			}
		})

		t.Run("check the size of unix.RtNexthop", func(t *testing.T) {
			// Check the size of unix.RtNexthop struct
			fmt.Printf("Size of RtNexthop: %d bytes\n", unsafe.Sizeof(RtNexthop{}))
			expected := unsafe.Sizeof(RtNexthop{})
			if actual := unsafe.Sizeof(*rtNextHop); actual != expected {
				t.Fatalf("size is not correct. actual: %d, expected: %d", actual, expected)
			}
		})

		t.Run("Check the deserialized Len using the Len field", func(t *testing.T) {
			const expectedLen = 32
			if rtNextHop.RtNexthop.Len != 8 {
				t.Fatalf("Expected Len to be 0x08, but got %d", rtNextHop.RtNexthop.Len)
			}
			// Check the deserialized Len using the Len() method from the interface
			if rtNextHop.Len() != expectedLen {
				t.Fatalf("Expected Len from Len() to be 0x08, but got %d", rtNextHop.Len())
			}
		})

		t.Run("Check the deserialized Flag", func(t *testing.T) {
			if rtNextHop.Flags != 0x01 {
				t.Fatalf("Expected Flags to be 0x01, but got %d", rtNextHop.Flags)
			}
		})

		t.Run("Check the deserialized Hops", func(t *testing.T) {
			if rtNextHop.Hops != 0x02 {
				t.Fatalf("Expected Hops to be 0x02, but got %d", rtNextHop.Hops)
			}
		})

		t.Run("Check the Ifindex value", func(t *testing.T) {
			expectedIfindex := int32(0x01020304)
			if rtNextHop.Ifindex != expectedIfindex {
				t.Fatalf("Expected Ifindex to be 0x%08x, but got 0x%08x", expectedIfindex, rtNextHop.Ifindex)
			}
		})
	})
}
