package core

import (
	"bytes"
	"encoding/binary"
	"testing"
)

func TestDeserializeTcHtbGlob(t *testing.T) {
	// Helper function to create a valid byte slice for TcHtbGlob
	createValidTcHtbGlobBytes := func() []byte {
		buf := new(bytes.Buffer)

		// Write values for TcHtbGlob fields: Version, Rate2Quantum, Defcls, Debug, DirectPkts (all uint32)
		_ = binary.Write(buf, binary.LittleEndian, uint32(1)) // Version
		_ = binary.Write(buf, binary.LittleEndian, uint32(2)) // Rate2Quantum
		_ = binary.Write(buf, binary.LittleEndian, uint32(3)) // Defcls
		_ = binary.Write(buf, binary.LittleEndian, uint32(4)) // Debug
		_ = binary.Write(buf, binary.LittleEndian, uint32(5)) // DirectPkts

		return buf.Bytes()
	}

	// Test case for valid input
	t.Run("valid input", func(t *testing.T) {
		// Create a valid byte slice
		validBytes := createValidTcHtbGlobBytes()

		// Call DeserializeTcHtbGlob
		tcHtbGlob, err := DeserializeTcHtbGlob(validBytes)

		// Verify no error occurred
		if err != nil {
			t.Fatalf("Expected no error, but got %v", err)
		}

		// Verify the deserialized values
		if tcHtbGlob.Version != 1 {
			t.Errorf("Expected Version to be 1, but got %d", tcHtbGlob.Version)
		}
		if tcHtbGlob.Rate2Quantum != 2 {
			t.Errorf("Expected Rate2Quantum to be 2, but got %d", tcHtbGlob.Rate2Quantum)
		}
		if tcHtbGlob.Defcls != 3 {
			t.Errorf("Expected Defcls to be 3, but got %d", tcHtbGlob.Defcls)
		}
		if tcHtbGlob.Debug != 4 {
			t.Errorf("Expected Debug to be 4, but got %d", tcHtbGlob.Debug)
		}
		if tcHtbGlob.DirectPkts != 5 {
			t.Errorf("Expected DirectPkts to be 5, but got %d", tcHtbGlob.DirectPkts)
		}
	})

	// Test case for invalid input (too short)
	t.Run("input too short", func(t *testing.T) {
		// Prepare an invalid byte slice (less than the expected size)
		invalidBytes := make([]byte, SizeOfTcHtbGlob-10) // Too short

		// Call DeserializeTcHtbGlob
		tcHtbGlob, err := DeserializeTcHtbGlob(invalidBytes)

		// Verify that an error occurred
		if err == nil {
			t.Fatalf("Expected an error due to short input, but got none")
		}

		// Verify that tcHtbGlob is nil
		if tcHtbGlob != nil {
			t.Errorf("Expected tcHtbGlob to be nil due to short input, but got %v", tcHtbGlob)
		}
	})
}
