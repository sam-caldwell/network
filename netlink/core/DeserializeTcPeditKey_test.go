package core

import (
	"testing"
)

func TestDeserializeTcPeditKey(t *testing.T) {
	// Helper function to create a valid byte slice for TcPeditKey
	createValidTcPeditKeyBytes := func() []byte {
		// Manually construct the byte slice to represent TcPeditKey fields explicitly.
		// TcPeditKey: Mask, Val, Off, At, OffMask, Shift (all uint32, BigEndian)
		return []byte{
			0x00, 0x00, 0x00, 0x01, // Mask (uint32)
			0x00, 0x00, 0x00, 0x02, // Val (uint32)
			0x00, 0x00, 0x00, 0x03, // Off (uint32)
			0x00, 0x00, 0x00, 0x04, // At (uint32)
			0x00, 0x00, 0x00, 0x05, // OffMask (uint32)
			0x00, 0x00, 0x00, 0x06, // Shift (uint32)
		}
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
		invalidBytes := []byte{
			0x00, 0x00, 0x00, 0x01, // Mask (uint32)
			0x00, 0x00, 0x00, 0x02, // Val (uint32)
			// Too short, missing the rest of the fields
		}

		// Call DeserializeTcPeditKey
		tcPeditKey, err := DeserializeTcPeditKey(invalidBytes)

		// Verify that an error occurred
		if err == nil {
			t.Fatalf("Expected error due to short input, but got none")
		}

		// Verify that tcPeditKey is nil
		if tcPeditKey != nil {
			t.Errorf("Expected tcPeditKey to be nil due to short input, but got %v", tcPeditKey)
		}
	})
}
