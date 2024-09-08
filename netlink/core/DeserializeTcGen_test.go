package core

import (
	"testing"
)

func TestDeserializeTcGen(t *testing.T) {
	// Helper function to create a valid byte slice for TcGen
	createValidTcGenBytes := func() []byte {
		return []byte{
			0x01, 0x00, 0x00, 0x00, // Index (uint32)
			0x02, 0x00, 0x00, 0x00, // Capab (uint32)
			0x03, 0x00, 0x00, 0x00, // Action (int32)
			0x04, 0x00, 0x00, 0x00, // Refcnt (int32)
			0x05, 0x00, 0x00, 0x00, // Bindcnt (int32)
		}
	}

	// Test case for valid input
	t.Run("valid input", func(t *testing.T) {
		// Create a valid byte slice
		validBytes := createValidTcGenBytes()

		// Call DeserializeTcGen
		tcGen, err := DeserializeTcGen(validBytes)

		// Verify no error occurred
		if err != nil {
			t.Fatalf("Expected no error, but got %v", err)
		}

		// Verify the deserialized values
		if tcGen.Index != 1 {
			t.Errorf("Expected Index to be 1, but got %d", tcGen.Index)
		}
		if tcGen.Capab != 2 {
			t.Errorf("Expected Capab to be 2, but got %d", tcGen.Capab)
		}
		if tcGen.Action != 3 {
			t.Errorf("Expected Action to be 3, but got %d", tcGen.Action)
		}
		if tcGen.Refcnt != 4 {
			t.Errorf("Expected Refcnt to be 4, but got %d", tcGen.Refcnt)
		}
		if tcGen.Bindcnt != 5 {
			t.Errorf("Expected Bindcnt to be 5, but got %d", tcGen.Bindcnt)
		}
	})

	// Test case for invalid input (too short)
	t.Run("input too short", func(t *testing.T) {
		// Prepare an invalid byte slice (less than the expected size)
		invalidBytes := []byte{
			0x01, 0x00, 0x00, 0x00, // Index (uint32)
		}

		// Call DeserializeTcGen
		tcGen, err := DeserializeTcGen(invalidBytes)

		// Verify that an error occurred
		if err == nil {
			t.Fatalf("Expected an error due to short input, but got none")
		}

		// Verify that tcGen is nil
		if tcGen != nil {
			t.Errorf("Expected tcGen to be nil due to short input, but got %v", tcGen)
		}
	})
}
