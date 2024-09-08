package core

import (
	"testing"
)

func TestDeserializeTcCsum(t *testing.T) {
	// Test case for valid input
	t.Run("valid input", func(t *testing.T) {
		// Prepare a valid byte slice (24 bytes)
		validBytes := []byte{
			// TcGen fields
			0x01, 0x00, 0x00, 0x00, // Index (uint32)
			0x02, 0x00, 0x00, 0x00, // Capab (uint32)
			0x03, 0x00, 0x00, 0x00, // Action (int32)
			0x04, 0x00, 0x00, 0x00, // Refcnt (int32)
			0x05, 0x00, 0x00, 0x00, // Bindcnt (int32)
			// UpdateFlags field
			0x06, 0x00, 0x00, 0x00, // UpdateFlags (uint32)
		}

		// Call DeserializeTcCsum
		tcCsum, err := DeserializeTcCsum(validBytes)

		// Verify no error occurred
		if err != nil {
			t.Fatalf("Expected no error, but got %v", err)
		}

		// Verify the deserialized values
		if tcCsum.TcGen.Index != 1 {
			t.Errorf("Expected Index to be 1, but got %d", tcCsum.TcGen.Index)
		}
		if tcCsum.TcGen.Capab != 2 {
			t.Errorf("Expected Capab to be 2, but got %d", tcCsum.TcGen.Capab)
		}
		if tcCsum.TcGen.Action != 3 {
			t.Errorf("Expected Action to be 3, but got %d", tcCsum.TcGen.Action)
		}
		if tcCsum.TcGen.Refcnt != 4 {
			t.Errorf("Expected Refcnt to be 4, but got %d", tcCsum.TcGen.Refcnt)
		}
		if tcCsum.TcGen.Bindcnt != 5 {
			t.Errorf("Expected Bindcnt to be 5, but got %d", tcCsum.TcGen.Bindcnt)
		}
		if tcCsum.UpdateFlags != 6 {
			t.Errorf("Expected UpdateFlags to be 6, but got %d", tcCsum.UpdateFlags)
		}
	})

	// Test case for invalid input (too short)
	t.Run("input too short", func(t *testing.T) {
		// Prepare an invalid byte slice (less than 24 bytes)
		invalidBytes := []byte{
			0x01, 0x00, 0x00, 0x00, // Index (uint32)
			0x02, 0x00, 0x00, 0x00, // Capab (uint32)
		}

		// Call DeserializeTcCsum
		tcCsum, err := DeserializeTcCsum(invalidBytes)

		// Verify that an error occurred
		if err == nil {
			t.Fatalf("Expected an error due to short input, but got none")
		}

		// Verify that tcCsum is nil
		if tcCsum != nil {
			t.Errorf("Expected tcCsum to be nil due to short input, but got %v", tcCsum)
		}
	})
}
