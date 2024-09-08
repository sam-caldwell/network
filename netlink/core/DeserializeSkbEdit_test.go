package core

import (
	"testing"
)

func TestDeserializeSkbEdit(t *testing.T) {
	// Test case for valid input
	t.Run("valid input", func(t *testing.T) {
		// Prepare a valid byte slice (20 bytes)
		buf := []byte{
			// Index (4 bytes, little-endian)
			0x01, 0x00, 0x00, 0x00,
			// Capab (4 bytes, little-endian)
			0x02, 0x00, 0x00, 0x00,
			// Action (4 bytes, little-endian)
			0x03, 0x00, 0x00, 0x00,
			// Refcnt (4 bytes, little-endian)
			0x04, 0x00, 0x00, 0x00,
			// Bindcnt (4 bytes, little-endian)
			0x05, 0x00, 0x00, 0x00,
		}

		// Call DeserializeSkbEdit
		skbEdit, err := DeserializeSkbEdit(buf)

		// Verify no error occurred
		if err != nil {
			t.Fatalf("Expected no error, but got %v", err)
		}

		// Verify the deserialized values
		if skbEdit.Index != 1 {
			t.Errorf("Expected Index to be 1, but got %d", skbEdit.Index)
		}
		if skbEdit.Capab != 2 {
			t.Errorf("Expected Capab to be 2, but got %d", skbEdit.Capab)
		}
		if skbEdit.Action != 3 {
			t.Errorf("Expected Action to be 3, but got %d", skbEdit.Action)
		}
		if skbEdit.Refcnt != 4 {
			t.Errorf("Expected Refcnt to be 4, but got %d", skbEdit.Refcnt)
		}
		if skbEdit.Bindcnt != 5 {
			t.Errorf("Expected Bindcnt to be 5, but got %d", skbEdit.Bindcnt)
		}
	})

	// Test case for invalid input (too short)
	t.Run("input too short", func(t *testing.T) {
		// Prepare an invalid byte slice (less than 20 bytes)
		buf := make([]byte, 10) // Too short

		// Call DeserializeSkbEdit
		skbEdit, err := DeserializeSkbEdit(buf)

		// Verify the function returns an error
		if err == nil {
			t.Fatalf("Expected error due to short input, but got none")
		}

		// Verify that skbEdit is nil
		if skbEdit != nil {
			t.Errorf("Expected skbEdit to be nil due to short input, but got %v", skbEdit)
		}
	})
}
