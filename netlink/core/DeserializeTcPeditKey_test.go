package core

import (
	"bytes"
	"encoding/binary"
	"testing"
)

func TestDeserializeTcPeditKey(t *testing.T) {
	// Helper function to create a valid byte slice for TcPeditKey
	createValidTcPeditKeyBytes := func() []byte {
		buf := new(bytes.Buffer)

		// Write values for TcPeditKey fields: Mask, Val, Off, At, OffMask, Shift (all uint32)
		_ = binary.Write(buf, binary.BigEndian, uint32(1)) // Mask
		_ = binary.Write(buf, binary.BigEndian, uint32(2)) // Val
		_ = binary.Write(buf, binary.BigEndian, uint32(3)) // Off
		_ = binary.Write(buf, binary.BigEndian, uint32(4)) // At
		_ = binary.Write(buf, binary.BigEndian, uint32(5)) // OffMask
		_ = binary.Write(buf, binary.BigEndian, uint32(6)) // Shift

		return buf.Bytes()
	}

	// Test case for valid input
	t.Run("valid input", func(t *testing.T) {
		// Create a valid byte slice
		validBytes := createValidTcPeditKeyBytes()

		// Call DeserializeTcPeditKey
		tcPeditKey, err := DeserializeTcPeditKey(validBytes)

		// Verify no error occurred
		if err != nil {
			t.Fatalf("Expected no error, but got %v", err)
		}

		// Verify the deserialized values
		if tcPeditKey.Mask != 1 {
			t.Errorf("Expected Mask to be 1, but got %d", tcPeditKey.Mask)
		}
		if tcPeditKey.Val != 2 {
			t.Errorf("Expected Val to be 2, but got %d", tcPeditKey.Val)
		}
		if tcPeditKey.Off != 3 {
			t.Errorf("Expected Off to be 3, but got %d", tcPeditKey.Off)
		}
		if tcPeditKey.At != 4 {
			t.Errorf("Expected At to be 4, but got %d", tcPeditKey.At)
		}
		if tcPeditKey.OffMask != 5 {
			t.Errorf("Expected OffMask to be 5, but got %d", tcPeditKey.OffMask)
		}
		if tcPeditKey.Shift != 6 {
			t.Errorf("Expected Shift to be 6, but got %d", tcPeditKey.Shift)
		}
	})

	// Test case for invalid input (too short)
	t.Run("input too short", func(t *testing.T) {
		// Prepare an invalid byte slice (less than the expected size)
		invalidBytes := make([]byte, SizeOfTcPeditKey-10) // Too short

		// Call DeserializeTcPeditKey
		tcPeditKey, err := DeserializeTcPeditKey(invalidBytes)

		// Verify that an error occurred
		if err == nil {
			t.Fatalf("Expected an error due to short input, but got none")
		}

		// Verify that tcPeditKey is nil
		if tcPeditKey != nil {
			t.Errorf("Expected tcPeditKey to be nil due to short input, but got %v", tcPeditKey)
		}
	})
}
