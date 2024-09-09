package core

import (
	"bytes"
	"encoding/binary"
	"golang.org/x/sys/unix"
	"testing"
)

func TestDeserializeUnixRtAttr(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		// Prepare a sample unix.RtAttr structure for testing
		expectedAttr := unix.RtAttr{
			Len:  16,
			Type: 2,
		}

		// Serialize the expected unix.RtAttr into a byte slice
		buf := new(bytes.Buffer)
		_ = binary.Write(buf, NativeEndian, expectedAttr.Len)
		_ = binary.Write(buf, NativeEndian, expectedAttr.Type)

		serializedData := buf.Bytes()

		// Deserialize the byte slice
		deserializedAttr, err := DeserializeUnixRtAttr(serializedData)
		if err != nil {
			t.Fatalf("Unexpected error during deserialization: %v", err)
		}

		// Check if the deserialized structure matches the expected structure
		if deserializedAttr.Len != expectedAttr.Len || deserializedAttr.Type != expectedAttr.Type {
			t.Fatalf("Deserialized object doesn't match expected.\nExpected: %+v\nGot: %+v", expectedAttr, deserializedAttr)
		}
	})

	t.Run("byte slice too short", func(t *testing.T) {
		// Create a short byte slice that is too small to contain a full RtAttr structure
		shortData := make([]byte, SizeOfUnixRtAttr-1)

		// Try deserializing the short byte slice
		_, err := DeserializeUnixRtAttr(shortData)
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
