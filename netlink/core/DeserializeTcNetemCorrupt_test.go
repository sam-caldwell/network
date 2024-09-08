package core

import (
	"testing"
)

func TestDeserializeTcNetemCorrupt(t *testing.T) {
	// Helper function to create a valid byte slice for TcNetemCorrupt
	createValidTcNetemCorruptBytes := func() []byte {
		return []byte{
			0x01, 0x00, 0x00, 0x00, // Probability (uint32)
			0x02, 0x00, 0x00, 0x00, // Correlation (uint32)
		}
	}

	// Test case for valid input
	t.Run("valid input", func(t *testing.T) {
		// Create a valid byte slice
		validBytes := createValidTcNetemCorruptBytes()

		// Call DeserializeTcNetemCorrupt
		tcNetemCorrupt := DeserializeTcNetemCorrupt(validBytes)

		// Check that the struct is not nil
		if tcNetemCorrupt == nil {
			t.Fatalf("Expected tcNetemCorrupt to be non-nil, but got nil")
		}

		// Verify the deserialized values
		if tcNetemCorrupt.Probability != 1 {
			t.Errorf("Expected Probability to be 1, but got %d", tcNetemCorrupt.Probability)
		}
		if tcNetemCorrupt.Correlation != 2 {
			t.Errorf("Expected Correlation to be 2, but got %d", tcNetemCorrupt.Correlation)
		}
	})

	// Test case for invalid input (too short)
	t.Run("input too short", func(t *testing.T) {
		// Prepare an invalid byte slice (less than 8 bytes)
		invalidBytes := []byte{
			0x01, 0x00, 0x00, 0x00, // Probability (uint32)
		}

		// Call DeserializeTcNetemCorrupt
		tcNetemCorrupt := DeserializeTcNetemCorrupt(invalidBytes)

		// Check that the function returns nil for insufficient byte length
		if tcNetemCorrupt != nil {
			t.Errorf("Expected tcNetemCorrupt to be nil due to short input, but got %v", tcNetemCorrupt)
		}
	})
}
