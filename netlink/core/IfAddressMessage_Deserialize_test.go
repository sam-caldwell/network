package core

import (
	"bytes"
	"encoding/binary"
	"github.com/sam-caldwell/convert"
	"testing"
)

func TestDeserializeIfAddress(t *testing.T) {

	t.Run("Deserialize method", func(t *testing.T) {
		var (
			testData     IfAddressMessage
			expectedData IfAddressMessage
			actualData   []byte
		)
		t.Run("setup actualData", func(t *testing.T) {
			actualData = []byte{
				0x02,
				0x10,
				0x11,
				0x12,
				0x20, 0x21, 0x22, 0x23,
			}
		})
		t.Run("setup expectedData", func(t *testing.T) {
			expectedData.Family = AfInet
			expectedData.Prefixlen = uint8(0x10)
			expectedData.Flags = uint8(0x11)
			expectedData.Scope = uint8(0x12)
			expectedData.Index = convert.BytesToUint32([4]byte{0x23, 0x22, 0x21, 0x20})
		})
		t.Run("run deserialize method", func(t *testing.T) {
			if err := testData.Deserialize(actualData); err != nil {
				t.Fatal(err)
			}
		})
		t.Run("compare outcome", func(t *testing.T) {
			if testData.Family != expectedData.Family {
				t.Fatalf("Family mismatch")
			}
			if testData.Prefixlen != expectedData.Prefixlen {
				t.Fatalf("Prefixlen mismatch")
			}
			if testData.Flags != expectedData.Flags {
				t.Fatalf("Flags mismatch")
			}
			if testData.Scope != expectedData.Scope {
				t.Fatalf("Scope mismatch")
			}
			if testData.Index != expectedData.Index {
				t.Fatalf("Index mismatch")
			}
		})
	})

	t.Run("DeserializeIfAddressMessage()", func(t *testing.T) {
		// Subtest 1: Valid input
		t.Run("valid input", func(t *testing.T) {
			// Prepare a byte slice with valid input (8 bytes total)
			buf := make([]byte, IfAddressMessageSize)

			// Populate the byte slice with values for IfAddrmsg fields
			buf[0] = 0x02                                  // Family
			buf[1] = 0x20                                  // Prefixlen
			buf[2] = 0x01                                  // Flags
			buf[3] = 0x10                                  // Scope
			binary.LittleEndian.PutUint32(buf[4:8], 12345) // Index

			// Call DeserializeIfAddressMessage
			msg, err := DeserializeIfAddressMessage(buf)

			// Check for no errors
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			// Check the deserialized values
			if msg.Family != 0x02 {
				t.Errorf("Expected Family to be 0x02, but got %d", msg.Family)
			}
			if msg.Prefixlen != 0x20 {
				t.Errorf("Expected Prefixlen to be 0x20, but got %d", msg.Prefixlen)
			}
			if msg.Flags != 0x01 {
				t.Errorf("Expected Flags to be 0x01, but got %d", msg.Flags)
			}
			if msg.Scope != 0x10 {
				t.Errorf("Expected Scope to be 0x10, but got %d", msg.Scope)
			}
			if msg.Index != 12345 {
				t.Errorf("Expected Index to be 12345, but got %d", msg.Index)
			}
		})

		// Subtest 2: Input too short
		t.Run(ErrInputTooShort, func(t *testing.T) {
			// Prepare a byte slice with insufficient length (only 4 bytes)
			buf := make([]byte, 4)

			// Call DeserializeIfAddressMessage
			_, err := DeserializeIfAddressMessage(buf)

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
	})

	t.Run("serialize-deserialize agreement", func(t *testing.T) {
		var (
			expectedData IfAddressMessage
			actualData   []byte
		)
		t.Run("setup actualData", func(t *testing.T) {
			actualData = []byte{
				0x02,
				0x10,
				0x11,
				0x12,
				0x20, 0x21, 0x22, 0x23,
			}
		})
		t.Run("deserialize", func(t *testing.T) {
			if err := expectedData.Deserialize(actualData); err != nil {
				t.Fatal(err)
			}
		})
		t.Run("serialize", func(t *testing.T) {
			if serializedData, err := expectedData.Serialize(); err != nil {
				t.Fatal(err)
			} else {
				if !bytes.Equal(actualData, serializedData) {
					t.Fatalf("serialized message did not match deserialized message\n"+
						"  actual: %v\n"+
						"expected: %v", actualData, serializedData)
				}
			}
		})
	})
}
