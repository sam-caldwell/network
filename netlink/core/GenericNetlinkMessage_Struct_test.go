package core

import (
	"testing"
	"unsafe"
)

func TestGenericNetlinkMessage(t *testing.T) {
	t.Run("GenericNetlinkMessage struct", func(t *testing.T) {
		t.Run("verify GenericNetlinkMessageSize", func(t *testing.T) {
			const expectedSizeInBytes = 2
			if GenericNetlinkMessageSize != expectedSizeInBytes {
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

}
