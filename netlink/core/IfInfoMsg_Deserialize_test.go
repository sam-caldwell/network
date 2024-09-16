package core

import (
	"encoding/binary"
	"golang.org/x/sys/unix"
	"testing"
)

func TestInfoMsgDeserialize(t *testing.T) {
	t.Run("DeserializeIfInfoMsg()", func(t *testing.T) {
		// Subtest 1: Valid input
		t.Run("valid input", func(t *testing.T) {
			// Prepare a byte slice with valid input (16 bytes total)
			buf := make([]byte, IfInfoMsgSize)

			// Populate the byte slice with values for IfInfomsg fields (using little-endian encoding)
			buf[0] = 0x02 // Family
			// Skip byte 1
			binary.NativeEndian.PutUint16(buf[2:4], 0x01)  // Type = 1
			binary.NativeEndian.PutUint32(buf[4:8], 100)   // Index = 100
			binary.NativeEndian.PutUint32(buf[8:12], 200)  // Flags = 200
			binary.NativeEndian.PutUint32(buf[12:16], 300) // Change = 300

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
			expectedError := ErrInputTooShort
			if err.Error() != expectedError {
				t.Errorf("Expected error message '%s', but got '%s'", expectedError, err.Error())
			}
		})

		t.Run("serialize-deserialize agreement", func(t *testing.T) {
			var (
				byteData []byte
				err      error
			)
			data := IfInfoMsg{
				IfInfomsg: unix.IfInfomsg{
					Family: unix.AF_INET,
					Type:   1,
					Index:  2,
					Flags:  3,
					Change: 4,
				},
			}
			t.Run("serialize input", func(t *testing.T) {
				if byteData, err = data.Serialize(); err != nil {
					t.Fatalf("Unexpected error: %v", err)
				}
				if sz := len(byteData); sz != IfInfoMsgSize {
					t.Fatalf("serialized data error.  sz:%d", sz)
				}
			})
			t.Run("deserialize input using function", func(t *testing.T) {
				actual, err := DeserializeIfInfoMsg(byteData)
				if err != nil {
					t.Fatalf("Unexpected error: %v", err)
				}
				if actual.Family != data.Family {
					t.Fatalf("Expected Family to be %d, but got %d", data.Family, actual.Family)
				}
				if actual.Type != data.Type {
					t.Fatalf("Expected Type to be %d, but got %d", data.Family, actual.Family)
				}
				if actual.Index != data.Index {
					t.Fatalf("Expected Index to be %d, but got %d", data.Family, actual.Family)
				}
				if actual.Flags != data.Flags {
					t.Fatalf("Expected Flags to be %d, but got %d", data.Family, actual.Family)
				}
				if actual.Change != data.Change {
					t.Fatalf("Expected Change to be %d, but got %d", data.Family, actual.Family)
				}
			})
			t.Run("deserialize input using method", func(t *testing.T) {
				var actual IfInfoMsg
				if actual.Deserialize(byteData) != nil {
					t.Fatalf("Unexpected error: %v", err)
				}
				if actual.Family != data.Family {
					t.Fatalf("Expected Family to be %d, but got %d", data.Family, actual.Family)
				}
				if actual.Type != data.Type {
					t.Fatalf("Expected Type to be %d, but got %d", data.Family, actual.Family)
				}
				if actual.Index != data.Index {
					t.Fatalf("Expected Index to be %d, but got %d", data.Family, actual.Family)
				}
				if actual.Flags != data.Flags {
					t.Fatalf("Expected Flags to be %d, but got %d", data.Family, actual.Family)
				}
				if actual.Change != data.Change {
					t.Fatalf("Expected Change to be %d, but got %d", data.Family, actual.Family)
				}
			})
		})
	})
}
