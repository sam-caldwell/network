package core

import (
	"bytes"
	"encoding/binary"
	"testing"
)

// TestDeserializeVfLinkState tests the DeserializeVfLinkState function.
func TestDeserializeVfLinkState(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		// Create a sample VfLinkState structure for testing
		expectedState := VfLinkState{
			Vf:        0x12345678,
			LinkState: 0x87654321,
		}

		// Serialize the expected VfLinkState into a byte slice
		buf := new(bytes.Buffer)
		_ = binary.Write(buf, NativeEndian, expectedState.Vf)
		_ = binary.Write(buf, NativeEndian, expectedState.LinkState)

		serializedData := buf.Bytes()

		// Deserialize the byte slice
		deserializedState, err := DeserializeVfLinkState(serializedData)
		if err != nil {
			t.Fatalf("Unexpected error during deserialization: %v", err)
		}

		// Check if the deserialized structure matches the expected structure
		if deserializedState.Vf != expectedState.Vf || deserializedState.LinkState != expectedState.LinkState {
			t.Fatalf("Deserialized object doesn't match expected.\nExpected: %+v\nGot: %+v", expectedState, deserializedState)
		}
	})

	t.Run("invalid length", func(t *testing.T) {
		// Create a short byte slice that is too small to contain a full VfLinkState structure
		shortData := make([]byte, SizeOfVfLinkState-1)

		// Try deserializing the short byte slice
		_, err := DeserializeVfLinkState(shortData)
		if err == nil {
			t.Fatal("Expected an error for short input, but got nil")
		}

		// Verify the expected error message
		expectedErr := "invalid length for VfLinkState"
		if err.Error() != expectedErr {
			t.Errorf("Expected error message %q, but got %q", expectedErr, err.Error())
		}
	})
}
