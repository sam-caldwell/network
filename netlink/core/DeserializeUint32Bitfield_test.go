package core

import (
	"bytes"
	"encoding/binary"
	"reflect"
	"testing"
)

// TestDeserializeUint32Bitfield tests the DeserializeUint32Bitfield function.
func TestDeserializeUint32Bitfield(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		expected := Uint32Bitfield{
			Value:    0x12345678,
			Selector: 0x87654321,
		}

		// Serialize the expected struct into a byte slice
		buf := new(bytes.Buffer)
		_ = binary.Write(buf, NativeEndian, expected.Value)
		_ = binary.Write(buf, NativeEndian, expected.Selector)

		serializedData := buf.Bytes()

		// Deserialize the byte slice
		deserializedBitfield, err := DeserializeUint32Bitfield(serializedData)
		if err != nil {
			t.Fatalf("Unexpected error during deserialization: %v", err)
		}

		// Compare the deserialized struct with the expected struct
		if !reflect.DeepEqual(deserializedBitfield, &expected) {
			t.Fatalf("Deserialized object doesn't match expected.\nExpected: %+v\nGot: %+v", expected, deserializedBitfield)
		}
	})

	t.Run("data slice too short", func(t *testing.T) {
		shortData := make([]byte, SizeOfUint32Bitfield-1)
		_, err := DeserializeUint32Bitfield(shortData)
		if err == nil {
			t.Fatal("Expected an error for short input, but got nil")
		}

		expectedErr := "data slice too short"
		if err.Error() != expectedErr {
			t.Errorf("Expected error message %q, but got %q", expectedErr, err.Error())
		}
	})
}
