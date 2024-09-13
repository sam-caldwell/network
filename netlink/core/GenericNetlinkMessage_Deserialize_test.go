package core

import (
	"testing"
)

func TestGenericNetlinkMessage_Deserialize(t *testing.T) {
	t.Run("deserialize function", func(t *testing.T) {
		// Subtest 1: Happy path with valid input
		t.Run("valid input", func(t *testing.T) {
			// Prepare a byte slice with valid input
			input := []byte{0x01, 0x02}

			// Call DeserializeGenlMsg
			msg, err := DeserializeGenericNetlinkMessage(input)

			// Check for no errors
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			// Check the fields of the returned GenericNetlinkMessage
			if msg.Command != 0x01 {
				t.Errorf("Expected Command to be 0x01, but got %d", msg.Command)
			}
			if msg.Version != 0x02 {
				t.Errorf("Expected Version to be 0x02, but got %d", msg.Version)
			}
		})

		t.Run("Sad path with input too short", func(t *testing.T) {
			// Prepare a byte slice with too short length
			input := []byte{0x01}

			// Call DeserializeGenlMsg
			_, err := DeserializeGenericNetlinkMessage(input)

			// Expect an error due to small byte slice
			if err == nil {
				t.Fatalf("Expected error due to small byte slice, but got none")
			}
		})
	})
}
