package core

import (
	"bytes"
	"encoding/binary"
	"testing"
)

// TestDeserializeVfRate tests the DeserializeVfRate function.
func TestDeserializeVfRate(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		// Create a sample VfRate structure for testing
		expectedRate := VfRate{
			Vf:        0x12345678,
			MinTxRate: 0x87654321,
			MaxTxRate: 0xABCDEF01,
		}

		// Serialize the expected VfRate into a byte slice
		buf := new(bytes.Buffer)
		_ = binary.Write(buf, NativeEndian, expectedRate.Vf)
		_ = binary.Write(buf, NativeEndian, expectedRate.MinTxRate)
		_ = binary.Write(buf, NativeEndian, expectedRate.MaxTxRate)

		serializedData := buf.Bytes()

		// Deserialize the byte slice
		deserializedRate, err := DeserializeVfRate(serializedData)
		if err != nil {
			t.Fatalf("Unexpected error during deserialization: %v", err)
		}

		// Check if the deserialized structure matches the expected structure
		if deserializedRate.Vf != expectedRate.Vf || deserializedRate.MinTxRate != expectedRate.MinTxRate ||
			deserializedRate.MaxTxRate != expectedRate.MaxTxRate {

			t.Fatalf("Deserialized object doesn't match expected.\n"+
				"Expected: %+v\n"+
				"Got: %+v",
				expectedRate, deserializedRate)
		}
	})

	t.Run("invalid length", func(t *testing.T) {
		// Create a byte slice that is too small to contain a full VfRate structure
		shortData := make([]byte, SizeOfVfRate-1)

		// Try deserializing the short byte slice
		_, err := DeserializeVfRate(shortData)
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
