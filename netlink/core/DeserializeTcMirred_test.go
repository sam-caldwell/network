package core

import (
	"bytes"
	"encoding/binary"
	"testing"
)

func TestDeserializeTcMirred(t *testing.T) {
	// Helper function to create a valid byte slice for TcMirred
	createValidTcMirredBytes := func() []byte {
		buf := new(bytes.Buffer)

		// Write values for TcGen (Index, Capab, Action, Refcnt, Bindcnt)
		_ = binary.Write(buf, NativeEndian, uint32(1)) // Index
		_ = binary.Write(buf, NativeEndian, uint32(2)) // Capab
		_ = binary.Write(buf, NativeEndian, int32(3))  // Action
		_ = binary.Write(buf, NativeEndian, int32(4))  // Refcnt
		_ = binary.Write(buf, NativeEndian, int32(5))  // Bindcnt

		// Write values for TcMirred-specific fields (Eaction, Ifindex)
		_ = binary.Write(buf, NativeEndian, int32(6))  // Eaction
		_ = binary.Write(buf, NativeEndian, uint32(7)) // Ifindex

		return buf.Bytes()
	}

	// Test case for valid input
	t.Run("valid input", func(t *testing.T) {
		// Create a valid byte slice
		validBytes := createValidTcMirredBytes()

		// Call DeserializeTcMirred
		tcMirred, err := DeserializeTcMirred(validBytes)

		// Verify no error occurred
		if err != nil {
			t.Fatalf("Expected no error, but got %v", err)
		}

		// Verify the deserialized values
		if tcMirred.TcGen.Index != 1 {
			t.Errorf("Expected Index to be 1, but got %d", tcMirred.TcGen.Index)
		}
		if tcMirred.TcGen.Capab != 2 {
			t.Errorf("Expected Capab to be 2, but got %d", tcMirred.TcGen.Capab)
		}
		if tcMirred.TcGen.Action != 3 {
			t.Errorf("Expected Action to be 3, but got %d", tcMirred.TcGen.Action)
		}
		if tcMirred.TcGen.Refcnt != 4 {
			t.Errorf("Expected Refcnt to be 4, but got %d", tcMirred.TcGen.Refcnt)
		}
		if tcMirred.TcGen.Bindcnt != 5 {
			t.Errorf("Expected Bindcnt to be 5, but got %d", tcMirred.TcGen.Bindcnt)
		}
		if tcMirred.Eaction != 6 {
			t.Errorf("Expected Eaction to be 6, but got %d", tcMirred.Eaction)
		}
		if tcMirred.Ifindex != 7 {
			t.Errorf("Expected Ifindex to be 7, but got %d", tcMirred.Ifindex)
		}
	})

	// Test case for invalid input (too short)
	t.Run("input too short", func(t *testing.T) {
		// Prepare an invalid byte slice (less than the expected size)
		invalidBytes := make([]byte, SizeOfTcMirred-10) // Too short

		// Call DeserializeTcMirred
		tcMirred, err := DeserializeTcMirred(invalidBytes)

		// Verify that an error occurred
		if err == nil {
			t.Fatalf("Expected an error due to short input, but got none")
		}

		// Verify that tcMirred is nil
		if tcMirred != nil {
			t.Errorf("Expected tcMirred to be nil due to short input, but got %v", tcMirred)
		}
	})
}
