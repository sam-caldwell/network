package core

import (
	"bytes"
	"encoding/binary"
	"testing"
)

func TestGenericNetlinkMessage_Serialize(t *testing.T) {
	// Create an instance of GenericNetlinkMessage with some values
	msg := &GenericNetlinkMessage{
		Command: 0x10, // Example command value
		Version: 0x01, // Example version value
	}

	// Call the Serialize method
	serializedBytes, err := msg.Serialize()
	if err != nil {
		t.Fatalf("Expected no error during serialization, but got: %v", err)
	}

	// Create a buffer to manually serialize the values for comparison
	expectedBuf := new(bytes.Buffer)

	// Write Command and Version manually
	if err := binary.Write(expectedBuf, NativeEndian, msg.Command); err != nil {
		t.Fatalf("Failed to write command manually: %v", err)
	}
	if err := binary.Write(expectedBuf, NativeEndian, msg.Version); err != nil {
		t.Fatalf("Failed to write version manually: %v", err)
	}

	// Convert the expected buffer to bytes
	expectedBytes := expectedBuf.Bytes()

	// Compare the length of serialized bytes
	if len(serializedBytes) != len(expectedBytes) {
		t.Errorf("Expected serialized length to be %d, but got %d",
			len(expectedBytes), len(serializedBytes))
	}

	// Compare the serialized output with the expected output
	for i := range serializedBytes {
		if serializedBytes[i] != expectedBytes[i] {
			t.Errorf("Byte mismatch at index %d: expected 0x%x, got 0x%x",
				i, expectedBytes[i], serializedBytes[i])
		}
	}
}
