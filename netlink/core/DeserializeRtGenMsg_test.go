package core

import (
	"testing"
)

func TestDeserializeRtGenMsg(t *testing.T) {
	t.Run("valid input", func(t *testing.T) {
		// Prepare a byte slice with valid input (1 byte for Family)
		buf := []byte{0x02} // Example value for Family

		// Call DeserializeRtGenMsg
		rtGenMsg := DeserializeRtGenMsg(buf)

		// Check the deserialized value
		if rtGenMsg.Family != 0x02 {
			t.Errorf("Expected Family to be 0x02, but got %d", rtGenMsg.Family)
		}
	})

	t.Run("empty input", func(t *testing.T) {
		// Prepare an empty byte slice
		buf := []byte{}

		// Call DeserializeRtGenMsg
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Expected panic due to empty input, but got none")
			}
		}()

		_ = DeserializeRtGenMsg(buf)
	})
}
