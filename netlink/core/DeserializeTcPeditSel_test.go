package core

import (
	"testing"
)

func TestDeserializeTcPeditSel(t *testing.T) {
	// Helper function to create a valid byte slice for TcPeditSel
	createValidTcPeditSelBytes := func() []byte {
		// Manually construct the byte slice to represent TcPeditSel fields explicitly.
		// TcPeditSel: TcGen (Index, Capab, Action, Refcnt, Bindcnt), NKeys, Flags
		return []byte{
			0x00, 0x00, 0x00, 0x01, // Index (uint32)
			0x00, 0x00, 0x00, 0x02, // Capab (uint32)
			0x00, 0x00, 0x00, 0x03, // Action (int32)
			0x00, 0x00, 0x00, 0x04, // Refcnt (int32)
			0x00, 0x00, 0x00, 0x05, // Bindcnt (int32)
			0x02,       // NKeys (uint8)
			0x01,       // Flags (uint8)
			0x00, 0x00, //Padding
		}
	}

	// Test case for valid input
	t.Run("valid input", func(t *testing.T) {
		// Create a valid byte slice
		validBytes := createValidTcPeditSelBytes()

		// Call DeserializeTcPeditSel
		tcPeditSel, err := DeserializeTcPeditSel(validBytes)

		// Verify no error occurred
		if err != nil {
			t.Fatalf("Expected no error, but got %v", err)
		}

		// Verify the deserialized values
		if tcPeditSel.TcGen.Index != 1 {
			t.Errorf("Expected Index to be 1, but got %d", tcPeditSel.TcGen.Index)
		}
		if tcPeditSel.TcGen.Capab != 2 {
			t.Errorf("Expected Capab to be 2, but got %d", tcPeditSel.TcGen.Capab)
		}
		if tcPeditSel.TcGen.Action != 3 {
			t.Errorf("Expected Action to be 3, but got %d", tcPeditSel.TcGen.Action)
		}
		if tcPeditSel.TcGen.Refcnt != 4 {
			t.Errorf("Expected Refcnt to be 4, but got %d", tcPeditSel.TcGen.Refcnt)
		}
		if tcPeditSel.TcGen.Bindcnt != 5 {
			t.Errorf("Expected Bindcnt to be 5, but got %d", tcPeditSel.TcGen.Bindcnt)
		}
		if tcPeditSel.NKeys != 2 {
			t.Errorf("Expected NKeys to be 2, but got %d", tcPeditSel.NKeys)
		}
		if tcPeditSel.Flags != 1 {
			t.Errorf("Expected Flags to be 1, but got %d", tcPeditSel.Flags)
		}
	})

	// Test case for invalid input (too short)
	t.Run("input too short", func(t *testing.T) {
		// Prepare an invalid byte slice (less than the expected size)
		invalidBytes := []byte{
			0x00, 0x00, 0x00, 0x01, // Index (uint32)
			0x00, 0x00, 0x00, 0x02, // Capab (uint32)
			// Too short, missing the rest of the fields
		}

		// Call DeserializeTcPeditSel
		tcPeditSel, err := DeserializeTcPeditSel(invalidBytes)

		// Verify that an error occurred
		if err == nil {
			t.Fatalf("Expected error due to short input, but got none")
		}

		// Verify that tcPeditSel is nil
		if tcPeditSel != nil {
			t.Errorf("Expected tcPeditSel to be nil due to short input, but got %v", tcPeditSel)
		}
	})
}
