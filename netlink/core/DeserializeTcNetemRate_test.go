package core

import (
	"testing"
)

func TestDeserializeTcNetemRate(t *testing.T) {
	// Helper function to create a valid byte slice for TcNetemRate
	createValidTcNetemRateBytes := func() []byte {
		return []byte{
			0x01, 0x00, 0x00, 0x00, // Rate (uint32)
			0x02, 0x00, 0x00, 0x00, // PacketOverhead (int32)
			0x03, 0x00, 0x00, 0x00, // CellSize (uint32)
			0x04, 0x00, 0x00, 0x00, // CellOverhead (int32)
		}
	}

	// Test case for valid input
	t.Run("valid input", func(t *testing.T) {
		// Create a valid byte slice
		validBytes := createValidTcNetemRateBytes()

		// Call DeserializeTcNetemRate
		tcNetemRate := DeserializeTcNetemRate(validBytes)

		// Check that the struct is not nil
		if tcNetemRate == nil {
			t.Fatalf("Expected tcNetemRate to be non-nil, but got nil")
		}

		// Verify the deserialized values
		if tcNetemRate.Rate != 1 {
			t.Errorf("Expected Rate to be 1, but got %d", tcNetemRate.Rate)
		}
		if tcNetemRate.PacketOverhead != 2 {
			t.Errorf("Expected PacketOverhead to be 2, but got %d", tcNetemRate.PacketOverhead)
		}
		if tcNetemRate.CellSize != 3 {
			t.Errorf("Expected CellSize to be 3, but got %d", tcNetemRate.CellSize)
		}
		if tcNetemRate.CellOverhead != 4 {
			t.Errorf("Expected CellOverhead to be 4, but got %d", tcNetemRate.CellOverhead)
		}
	})

	// Test case for invalid input (too short)
	t.Run("input too short", func(t *testing.T) {
		// Prepare an invalid byte slice (less than the expected size)
		invalidBytes := make([]byte, 8) // Too short

		// Call DeserializeTcNetemRate
		tcNetemRate := DeserializeTcNetemRate(invalidBytes)

		// Check that the function returns nil for insufficient byte length
		if tcNetemRate != nil {
			t.Errorf("Expected tcNetemRate to be nil due to short input, but got %v", tcNetemRate)
		}
	})
}
