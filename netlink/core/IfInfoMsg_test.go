package core

import (
	"encoding/binary"
	"golang.org/x/sys/unix"
	"testing"
	"unsafe"
)

func TestInfoMsg(t *testing.T) {
	t.Run("IfInfoMsg struct", func(t *testing.T) {
		t.Run("sizeOfIfInfoMsg check", func(t *testing.T) {
			if SizeOfIfInfoMsg != 16 {
				t.Fatal("SizeOfIfInfoMsg mismatch")
			}
		})
		t.Run("check size", func(t *testing.T) {
			const expectedSizeInBytes = SizeOfIfInfoMsg
			if sz := int(unsafe.Sizeof(IfInfoMsg{})); sz != expectedSizeInBytes {
				t.Fatal("size mismatch")
			}
		})
		t.Run("field check", func(t *testing.T) {
			_ = IfInfoMsg{
				unix.IfInfomsg{
					Family: uint8(0),
					Type:   uint16(0),
					Index:  int32(0),
					Flags:  uint32(0),
					Change: uint32(0),
				},
			}
		})
	})

	t.Run("DeserializeIfInfoMsg()", func(t *testing.T) {
		// Subtest 1: Valid input
		t.Run("valid input", func(t *testing.T) {
			// Prepare a byte slice with valid input (16 bytes total)
			buf := make([]byte, SizeOfIfInfoMsg)

			// Populate the byte slice with values for IfInfomsg fields (using little-endian encoding)
			buf[0] = 0x02 // Family
			// Skip byte 1
			binary.LittleEndian.PutUint16(buf[2:4], 0x01)  // Type = 1
			binary.LittleEndian.PutUint32(buf[4:8], 100)   // Index = 100
			binary.LittleEndian.PutUint32(buf[8:12], 200)  // Flags = 200
			binary.LittleEndian.PutUint32(buf[12:16], 300) // Change = 300

			// Call DeserializeIfInfoMsg
			msg, err := DeserializeIfInfoMsg(buf)

			// Check for no errors
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			// Check the deserialized values
			if msg.Family != 0x02 {
				t.Errorf("Expected Family to be 0x02, but got %d", msg.Family)
			}
			if msg.Type != 0x01 {
				t.Errorf("Expected Type to be 0x01, but got %d", msg.Type)
			}
			if msg.Index != 100 {
				t.Errorf("Expected Index to be 100, but got %d", msg.Index)
			}
			if msg.Flags != 200 {
				t.Errorf("Expected Flags to be 200, but got %d", msg.Flags)
			}
			if msg.Change != 300 {
				t.Errorf("Expected Change to be 300, but got %d", msg.Change)
			}
		})

		// Subtest 2: Input too short
		t.Run(ErrInputTooShort, func(t *testing.T) {
			// Prepare a byte slice with insufficient length (less than 16 bytes)
			buf := make([]byte, 8)

			// Call DeserializeIfInfoMsg
			_, err := DeserializeIfInfoMsg(buf)

			// Expect an error due to short byte slice
			if err == nil {
				t.Fatalf("Expected error due to short byte slice, but got none")
			}

			// Ensure the error message is correct
			expectedError := "IfInfoMsg to short"
			if err.Error() != expectedError {
				t.Errorf("Expected error message '%s', but got '%s'", expectedError, err.Error())
			}
		})
	})
}
