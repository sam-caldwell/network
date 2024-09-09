package core

import (
	"bytes"
	"encoding/binary"
	"testing"
)

// TestDeserializeVfMac tests the DeserializeVfMac function.
func TestDeserializeVfMac(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		// Create a sample VfMac structure for testing
		expectedMac := VfMac{
			Vf:  0x12345678,
			Mac: [32]byte{0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0xAA, 0xBB, 0xCC, 0xDD, 0xEE, 0xFF},
		}

		// Serialize the expected VfMac into a byte slice
		buf := new(bytes.Buffer)
		_ = binary.Write(buf, NativeEndian, expectedMac.Vf)
		buf.Write(expectedMac.Mac[:])

		serializedData := buf.Bytes()

		// Deserialize the byte slice
		deserializedMac, err := DeserializeVfMac(serializedData)
		if err != nil {
			t.Fatalf("Unexpected error during deserialization: %v", err)
		}

		// Check if the deserialized structure matches the expected structure
		if deserializedMac.Vf != expectedMac.Vf || !bytes.Equal(deserializedMac.Mac[:], expectedMac.Mac[:]) {
			t.Fatalf("Deserialized object doesn't match expected.\nExpected: %+v\nGot: %+v", expectedMac, deserializedMac)
		}
	})

	t.Run("input too short", func(t *testing.T) {
		// Create a byte slice that's too short
		shortData := make([]byte, SizeOfVfMac-1)

		// Try deserializing the short byte slice
		_, err := DeserializeVfMac(shortData)
		if err == nil {
			t.Fatal("Expected an error for short input, but got nil")
		}

		// Verify the expected error message
		expectedErr := "input too short to deserialize VfMac"
		if err.Error() != expectedErr {
			t.Errorf("Expected error message %q, but got %q", expectedErr, err.Error())
		}
	})
}
