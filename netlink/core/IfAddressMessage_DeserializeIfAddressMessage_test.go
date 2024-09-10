package core

import (
	"encoding/binary"
	"testing"
)

func TestDeserializeIfAddressMessage(t *testing.T) {
	// Subtest 1: Valid input
	t.Run("valid input", func(t *testing.T) {
		// Prepare a byte slice with valid input (8 bytes total)
		buf := make([]byte, SizeOfIfAddressMessage)

		// Populate the byte slice with values for IfAddrmsg fields
		buf[0] = 0x02                                  // Family
		buf[1] = 0x20                                  // Prefixlen
		buf[2] = 0x01                                  // Flags
		buf[3] = 0x10                                  // Scope
		binary.LittleEndian.PutUint32(buf[4:8], 12345) // Index

		// Call DeserializeIfAddressMessage
		msg, err := DeserializeIfAddressMessage(buf)

		// Check for no errors
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		// Check the deserialized values
		if msg.Family != 0x02 {
			t.Errorf("Expected Family to be 0x02, but got %d", msg.Family)
		}
		if msg.Prefixlen != 0x20 {
			t.Errorf("Expected Prefixlen to be 0x20, but got %d", msg.Prefixlen)
		}
		if msg.Flags != 0x01 {
			t.Errorf("Expected Flags to be 0x01, but got %d", msg.Flags)
		}
		if msg.Scope != 0x10 {
			t.Errorf("Expected Scope to be 0x10, but got %d", msg.Scope)
		}
		if msg.Index != 12345 {
			t.Errorf("Expected Index to be 12345, but got %d", msg.Index)
		}
	})

	// Subtest 2: Input too short
	t.Run("input too short", func(t *testing.T) {
		// Prepare a byte slice with insufficient length (only 4 bytes)
		buf := make([]byte, 4)

		// Call DeserializeIfAddressMessage
		_, err := DeserializeIfAddressMessage(buf)

		// Expect an error due to short byte slice
		if err == nil {
			t.Fatalf("Expected error due to short byte slice, but got none")
		}

		// Ensure the error message is correct
		expectedError := "input too short"
		if err.Error() != expectedError {
			t.Errorf("Expected error message '%s', but got '%s'", expectedError, err.Error())
		}
	})
}
