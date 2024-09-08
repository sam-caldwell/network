package core

import (
	"bytes"
	"encoding/binary"
	"testing"
)

func TestDeserializeTcMsg(t *testing.T) {
	// Helper function to create a valid byte slice for TcMsg
	createValidTcMsgBytes := func() []byte {
		buf := new(bytes.Buffer)

		// Write values for TcMsg fields: Family, Pad, Ifindex, Handle, Parent, Info
		binary.Write(buf, NativeEndian, uint8(1))         // Family
		binary.Write(buf, NativeEndian, [3]byte{0, 0, 0}) // Pad (3 bytes)
		binary.Write(buf, NativeEndian, int32(2))         // Ifindex
		binary.Write(buf, NativeEndian, uint32(3))        // Handle
		binary.Write(buf, NativeEndian, uint32(4))        // Parent
		binary.Write(buf, NativeEndian, uint32(5))        // Info

		return buf.Bytes()
	}

	// Test case for valid input
	t.Run("valid input", func(t *testing.T) {
		// Create a valid byte slice
		validBytes := createValidTcMsgBytes()

		// Call DeserializeTcMsg
		tcMsg, err := DeserializeTcMsg(validBytes)

		// Verify no error occurred
		if err != nil {
			t.Fatalf("Expected no error, but got %v", err)
		}

		// Verify the deserialized values
		if tcMsg.Family != 1 {
			t.Errorf("Expected Family to be 1, but got %d", tcMsg.Family)
		}
		if tcMsg.Ifindex != 2 {
			t.Errorf("Expected Ifindex to be 2, but got %d", tcMsg.Ifindex)
		}
		if tcMsg.Handle != 3 {
			t.Errorf("Expected Handle to be 3, but got %d", tcMsg.Handle)
		}
		if tcMsg.Parent != 4 {
			t.Errorf("Expected Parent to be 4, but got %d", tcMsg.Parent)
		}
		if tcMsg.Info != 5 {
			t.Errorf("Expected Info to be 5, but got %d", tcMsg.Info)
		}
	})

	// Test case for invalid input (too short)
	t.Run("input too short", func(t *testing.T) {
		// Prepare an invalid byte slice (less than the expected size)
		invalidBytes := make([]byte, SizeOfTcMsg-10) // Too short

		// Call DeserializeTcMsg
		tcMsg, err := DeserializeTcMsg(invalidBytes)

		// Verify that an error occurred
		if err == nil {
			t.Fatalf("Expected an error due to short input, but got none")
		}

		// Verify that tcMsg is nil
		if tcMsg != nil {
			t.Errorf("Expected tcMsg to be nil due to short input, but got %v", tcMsg)
		}
	})
}
