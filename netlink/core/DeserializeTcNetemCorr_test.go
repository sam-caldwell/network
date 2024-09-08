package core

import (
	"testing"
)

func TestDeserializeTcNetemCorr(t *testing.T) {
	// Helper function to create a valid byte slice for TcNetemCorr
	createValidTcNetemCorrBytes := func() []byte {
		return []byte{
			0x01, 0x00, 0x00, 0x00, // DelayCorr (uint32)
			0x02, 0x00, 0x00, 0x00, // LossCorr (uint32)
			0x03, 0x00, 0x00, 0x00, // DupCorr (uint32)
		}
	}

	// Test case for valid input
	t.Run("valid input", func(t *testing.T) {
		// Create a valid byte slice
		validBytes := createValidTcNetemCorrBytes()

		// Call DeserializeTcNetemCorr
		tcNetemCorr := DeserializeTcNetemCorr(validBytes)

		// Verify the deserialized values
		if tcNetemCorr.DelayCorr != 1 {
			t.Errorf("Expected DelayCorr to be 1, but got %d", tcNetemCorr.DelayCorr)
		}
		if tcNetemCorr.LossCorr != 2 {
			t.Errorf("Expected LossCorr to be 2, but got %d", tcNetemCorr.LossCorr)
		}
		if tcNetemCorr.DupCorr != 3 {
			t.Errorf("Expected DupCorr to be 3, but got %d", tcNetemCorr.DupCorr)
		}
	})

	// Test case for invalid input (too short)
	t.Run("input too short", func(t *testing.T) {
		// Prepare an invalid byte slice (less than the expected size)
		invalidBytes := []byte{
			0x01, 0x00, 0x00, 0x00, // DelayCorr (uint32)
			0x02, 0x00, 0x00, 0x00, // LossCorr (uint32)
		}

		// Call DeserializeTcNetemCorr and expect it to panic or return an incomplete structure
		defer func() {
			if r := recover(); r == nil {
				t.Fatalf("Expected a panic due to short input, but it did not occur")
			}
		}()
		_ = DeserializeTcNetemCorr(invalidBytes) // This should trigger a panic
	})
}
