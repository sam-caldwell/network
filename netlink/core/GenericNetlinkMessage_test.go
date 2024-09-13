package core

import (
	"bytes"
	"encoding/binary"
	"testing"
	"unsafe"
)

func TestGenericNetlinkMessage(t *testing.T) {
	t.Run("GenericNetlinkMessage struct", func(t *testing.T) {
		t.Run("verify SizeOfGenericNetlinkMessage", func(t *testing.T) {
			const expectedSizeInBytes = 2
			if SizeOfGenericNetlinkMessage != expectedSizeInBytes {
				t.Fatal("SizeOfGenlMsg size mismatch")
			}
		})
		t.Run("size check", func(t *testing.T) {
			// Ensure the size of the GenericNetlinkMessage struct matches the expected size (2 bytes)
			const expectedSizeInBytes = 2 // In C: __u8 cmd (1 byte) + __u8 version (1 byte)
			actualSize := unsafe.Sizeof(GenericNetlinkMessage{})
			if actualSize != uintptr(expectedSizeInBytes) {
				t.Errorf("Expected GenericNetlinkMessage size to be %d bytes, but got %d bytes",
					expectedSizeInBytes, actualSize)
			}
		})
		t.Run("test the correctness of the GenericNetlinkMessage struct fields", func(t *testing.T) {
			// Create a new GenericNetlinkMessage with specific values
			msg := GenericNetlinkMessage{
				Command: 0x10, // Example command value
				Version: 0x01, // Example version value
			}

			// Check if the Command field is correctly set
			if msg.Command != 0x10 {
				t.Errorf("Expected Command to be 0x10, but got 0x%x", msg.Command)
			}

			// Check if the Version field is correctly set
			if msg.Version != 0x01 {
				t.Errorf("Expected Version to be 0x01, but got 0x%x", msg.Version)
			}
		})
	})
	t.Run("test Len() method", func(t *testing.T) {
		// Create an instance of GenericNetlinkMessage with some values
		msg := &GenericNetlinkMessage{
			Command: 0x10, // Example command value
			Version: 0x01, // Example version value
		}

		// Call Len() to get the length
		length := msg.Len()

		// Calculate the expected length using binary.Size directly
		expectedLength := binary.Size(*msg)

		// Check if the calculated length matches the expected length
		if length != expectedLength {
			t.Errorf("Expected length to be %d, but got %d", expectedLength, length)
		}
	})
	t.Run("serialize method", func(t *testing.T) {
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
	})
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
