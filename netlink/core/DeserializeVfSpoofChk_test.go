package core

import (
	"bytes"
	"encoding/binary"
	"testing"
)

// TestDeserializeVfSpoofChk tests the DeserializeVfSpoofChk function.
func TestDeserializeVfSpoofChk(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		// Create a sample VfSpoofchk structure for testing
		expectedMsg := VfSpoofchk{
			Vf:      0x12345678,
			Setting: 0x87654321,
		}

		// Serialize the expected VfSpoofchk into a byte slice
		buf := new(bytes.Buffer)
		_ = binary.Write(buf, NativeEndian, expectedMsg.Vf)
		_ = binary.Write(buf, NativeEndian, expectedMsg.Setting)

		serializedData := buf.Bytes()

		// Deserialize the byte slice
		deserializedMsg, err := DeserializeVfSpoofChk(serializedData)
		if err != nil {
			t.Fatalf("Unexpected error during deserialization: %v", err)
		}

		// Check if the deserialized structure matches the expected structure
		if deserializedMsg.Vf != expectedMsg.Vf || deserializedMsg.Setting != expectedMsg.Setting {
			t.Fatalf("Deserialized object doesn't match expected.\nExpected: %+v\nGot: %+v", expectedMsg, deserializedMsg)
		}
	})

	t.Run("invalid length", func(t *testing.T) {
		// Create a byte slice that is too small to contain a full VfSpoofchk structure
		shortData := make([]byte, SizeOfVfSpoofchk-1)

		// Try deserializing the short byte slice
		_, err := DeserializeVfSpoofChk(shortData)
		if err == nil {
			t.Fatal("Expected an error for short input, but got nil")
		}

		// Verify the expected error message
		expectedErr := "input too short"
		if err.Error() != expectedErr {
			t.Errorf("Expected error message %q, but got %q", expectedErr, err.Error())
		}
	})
}
