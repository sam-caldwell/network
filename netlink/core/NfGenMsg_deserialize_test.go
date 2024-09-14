package core

import (
	"encoding/binary"
	"testing"
)

func TestDeserializeNfgenmsg(t *testing.T) {

	t.Run("valid input", func(t *testing.T) {
		// Prepare a byte slice with valid input (4 bytes total)
		buf := make([]byte, NfGenMsgSize)

		// Populate the byte slice with values
		buf[0] = 0x01                                // NfgenFamily
		buf[1] = 0x02                                // Version
		binary.BigEndian.PutUint16(buf[2:4], 0x0304) // ResId

		// Call DeserializeNfgenmsg
		nfgenmsg, err := DeserializeNfgenmsg(buf)

		// Check for no errors
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		// Check the deserialized values
		if nfgenmsg.NfgenFamily != 0x01 {
			t.Errorf("Expected NfgenFamily to be 0x01, but got %d", nfgenmsg.NfgenFamily)
		}
		if nfgenmsg.Version != 0x02 {
			t.Errorf("Expected Version to be 0x02, but got %d", nfgenmsg.Version)
		}
		if nfgenmsg.ResId != 0x0304 {
			t.Errorf("Expected ResId to be 0x0304, but got %d", nfgenmsg.ResId)
		}
	})

	t.Run(ErrInputTooShort, func(t *testing.T) {
		// Prepare a byte slice with insufficient length (less than 4 bytes)
		buf := make([]byte, 3)

		// Call DeserializeNfgenmsg
		_, err := DeserializeNfgenmsg(buf)

		// Expect an error due to short byte slice
		if err == nil {
			t.Fatalf("Expected error due to short byte slice, but got none")
		}

		// Ensure the error message is correct
		expectedError := ErrInputTooShort
		if err.Error() != expectedError {
			t.Errorf("Expected error message '%s', but got '%s'", expectedError, err.Error())
		}
	})
}
