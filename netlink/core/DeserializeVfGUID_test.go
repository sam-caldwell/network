package core

import (
	"bytes"
	"encoding/binary"
	"testing"
)

// TestDeserializeVfGUID tests the DeserializeVfGUID function.
func TestDeserializeVfGUID(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		// Create a sample VfGUID structure for testing
		expectedGUID := VfGUID{
			Vf:   0x12345678,
			Rsvd: 0x87654321,
			GUID: 0x123456789ABCDEF0,
		}

		// Serialize the expected VfGUID into a byte slice
		buf := new(bytes.Buffer)
		_ = binary.Write(buf, NativeEndian, expectedGUID.Vf)
		_ = binary.Write(buf, NativeEndian, expectedGUID.Rsvd)
		_ = binary.Write(buf, NativeEndian, expectedGUID.GUID)

		serializedData := buf.Bytes()

		// Deserialize the byte slice
		deserializedGUID, err := DeserializeVfGUID(serializedData)
		if err != nil {
			t.Fatalf("Unexpected error during deserialization: %v", err)
		}

		// Check if the deserialized structure matches the expected structure
		if deserializedGUID.Vf != expectedGUID.Vf || deserializedGUID.Rsvd != expectedGUID.Rsvd || deserializedGUID.GUID != expectedGUID.GUID {
			t.Fatalf("Deserialized object doesn't match expected.\nExpected: %+v\nGot: %+v", expectedGUID, deserializedGUID)
		}
	})

	t.Run("byte slice too short", func(t *testing.T) {
		// Create a short byte slice that is too small to contain a full VfGUID structure
		shortData := make([]byte, SizeOfVfGUID-1)

		// Try deserializing the short byte slice
		_, err := DeserializeVfGUID(shortData)
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
