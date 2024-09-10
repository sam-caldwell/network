package core

import (
	"encoding/binary"
	"testing"
)

func TestDeserializeIfaCacheInfo(t *testing.T) {
	// Subtest 1: Valid input
	t.Run("valid input", func(t *testing.T) {
		// Prepare a byte slice with valid input (16 bytes total)
		buf := make([]byte, SizeOfIfaCacheInfo)

		// Populate the byte slice with values for IfaCacheInfo fields (using little-endian encoding)
		binary.LittleEndian.PutUint32(buf[0:4], 100)   // Prefered = 100
		binary.LittleEndian.PutUint32(buf[4:8], 200)   // Valid = 200
		binary.LittleEndian.PutUint32(buf[8:12], 300)  // Cstamp = 300
		binary.LittleEndian.PutUint32(buf[12:16], 400) // Tstamp = 400

		// Call DeserializeIfaCacheInfo
		info, err := DeserializeIfaCacheInfo(buf)

		// Check for no errors
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		// Check the deserialized values
		if info.Prefered != 100 {
			t.Errorf("Expected Prefered to be 100, but got %d", info.Prefered)
		}
		if info.Valid != 200 {
			t.Errorf("Expected Valid to be 200, but got %d", info.Valid)
		}
		if info.Cstamp != 300 {
			t.Errorf("Expected Cstamp to be 300, but got %d", info.Cstamp)
		}
		if info.Tstamp != 400 {
			t.Errorf("Expected Tstamp to be 400, but got %d", info.Tstamp)
		}
	})

	// Subtest 2: Invalid input (too short)
	t.Run("input too short", func(t *testing.T) {
		// Prepare a byte slice with insufficient length
		buf := make([]byte, 12) // Less than the required 16 bytes

		// Call DeserializeIfaCacheInfo
		_, err := DeserializeIfaCacheInfo(buf)

		// Expect an error due to short byte slice
		if err == nil {
			t.Fatalf("Expected error due to short byte slice, but got none")
		}

		// Ensure the error message is correct
		expectedError := "data slice too short"
		if err.Error() != expectedError {
			t.Errorf("Expected error message '%s', but got '%s'", expectedError, err.Error())
		}
	})
}
