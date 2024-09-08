package core

import (
	"bytes"
	"encoding/binary"
	"testing"
)

func TestDeserializeTcf(t *testing.T) {
	// Helper function to create a valid byte slice for Tcf
	createValidTcfBytes := func() []byte {
		buf := new(bytes.Buffer)

		// Write values for Tcf fields: Install, LastUse, Expires, FirstUse (all uint64)
		binary.Write(buf, NativeEndian, uint64(1)) // Install
		binary.Write(buf, NativeEndian, uint64(2)) // LastUse
		binary.Write(buf, NativeEndian, uint64(3)) // Expires
		binary.Write(buf, NativeEndian, uint64(4)) // FirstUse

		return buf.Bytes()
	}

	// Test case for valid input
	t.Run("valid input", func(t *testing.T) {
		// Create a valid byte slice
		validBytes := createValidTcfBytes()

		// Call DeserializeTcf
		tcf, err := DeserializeTcf(validBytes)

		// Verify no error occurred
		if err != nil {
			t.Fatalf("Expected no error, but got %v", err)
		}

		// Verify the deserialized values
		if tcf.Install != 1 {
			t.Errorf("Expected Install to be 1, but got %d", tcf.Install)
		}
		if tcf.LastUse != 2 {
			t.Errorf("Expected LastUse to be 2, but got %d", tcf.LastUse)
		}
		if tcf.Expires != 3 {
			t.Errorf("Expected Expires to be 3, but got %d", tcf.Expires)
		}
		if tcf.FirstUse != 4 {
			t.Errorf("Expected FirstUse to be 4, but got %d", tcf.FirstUse)
		}
	})

	// Test case for invalid input (too short)
	t.Run("input too short", func(t *testing.T) {
		// Prepare an invalid byte slice (less than the size of Tcf)
		invalidBytes := []byte{
			0x01, 0x00, 0x00, 0x00, // Too short
		}

		// Call DeserializeTcf
		tcf, err := DeserializeTcf(invalidBytes)

		// Verify that an error occurred
		if err == nil {
			t.Fatalf("Expected an error due to short input, but got none")
		}

		// Verify that tcf is nil
		if tcf != nil {
			t.Errorf("Expected tcf to be nil due to short input, but got %v", tcf)
		}
	})
}
