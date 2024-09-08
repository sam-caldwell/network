package core

import (
	"testing"
)

func TestDeserializeTcNetemReorder(t *testing.T) {

	// Test case for valid input
	t.Run("valid input", func(t *testing.T) {
		validInputBytes := []byte{
			0x01, 0x00, 0x00, 0x00, // Probability (uint32)
			0x02, 0x00, 0x00, 0x00, // Correlation (uint32)
		}

		// Call DeserializeTcNetemReorder
		tcNetemReorder, err := DeserializeTcNetemReorder(validInputBytes)
		if err != nil {
			t.Errorf("DeserializeTcNetemReorder failed: %v", err)
		}

		// Check that the struct is not nil
		if tcNetemReorder == nil {
			t.Fatalf("Expected tcNetemReorder to be non-nil, but got nil")
		}

		// Verify the deserialized values
		if tcNetemReorder.Probability != 1 {
			t.Errorf("Expected Probability to be 1, but got %d", tcNetemReorder.Probability)
		}

		if tcNetemReorder.Correlation != 2 {
			t.Errorf("Expected Correlation to be 2, but got %d", tcNetemReorder.Correlation)
		}

	})

	// Test case for invalid input (too short)
	t.Run("input too short", func(t *testing.T) {
		// Prepare an invalid byte slice (less than the expected size)
		invalidBytes := []byte{
			0x01, 0x00, 0x00, 0x00, // Probability (uint32)
		}

		// Call DeserializeTcNetemReorder
		tcNetemReorder, err := DeserializeTcNetemReorder(invalidBytes)
		if err == nil {
			t.Errorf("Expected DeserializeTcNetemReorder to return an error, but got nil")
		}
		if tcNetemReorder != nil {
			t.Errorf("Expected tcNetemReorder to be nil, but got %v", tcNetemReorder)
		}
	})
}
