package core

import (
	"bytes"
	"encoding/binary"
	"github.com/sam-caldwell/convert"
	"golang.org/x/sys/unix"
	"testing"
	"unsafe"
)

func TestIfAddressMessage(t *testing.T) {
	t.Run("test SizeOfIfAddressMessage", func(t *testing.T) {
		const expectedSizeInBytes = 8
		if SizeOfIfAddressMessage != expectedSizeInBytes {
			t.Fatalf("SizeOfIfAddressMessage mismatch")
		}
	})
	t.Run("test IfAddressMessage struct", func(t *testing.T) {
		t.Run("ipv4 should be supported", func(t *testing.T) {
			t.Run("field check", func(t *testing.T) {
				i := IfAddressMessage{}
				i.Family = AfInet
				i.Prefixlen = uint8(0)
				i.Flags = uint8(0)
				i.Scope = uint8(0)
				i.Index = uint32(0)
			})
			t.Run("size check", func(t *testing.T) {
				const expectedSizeInBytes = SizeOfIfAddressMessage
				if sz := int(unsafe.Sizeof(IfAddressMessage{})); sz != expectedSizeInBytes {
					t.Fatalf("size mismatch.  sz: %d", sz)
				}
			})
		})

		t.Run("ipv6 should be supported", func(t *testing.T) {
			t.Run("field check", func(t *testing.T) {
				i := IfAddressMessage{}
				i.Family = AfInet6
				i.Prefixlen = uint8(0)
				i.Flags = uint8(0)
				i.Scope = uint8(0)
				i.Index = uint32(0)
			})
			t.Run("size check", func(t *testing.T) {
				const expectedSizeInBytes = SizeOfIfAddressMessage
				if sz := int(unsafe.Sizeof(IfAddressMessage{})); sz != expectedSizeInBytes {
					t.Fatalf("size mismatch.  sz: %d", sz)
				}
			})
		})
		t.Run("Test IfAddressMessage creation", func(t *testing.T) {
			// Define test cases with different InterfaceFamily values.
			testCases := []struct {
				name   string
				family InterfaceFamily
			}{
				{"IPv4", InterfaceFamily(unix.AF_INET)},
				{"IPv6", InterfaceFamily(unix.AF_INET6)},
				{"Unspecified", InterfaceFamily(unix.AF_UNSPEC)},
				// Add more interface families if needed.
			}

			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					msg := NewIfAddressMessage(tc.family)

					if msg == nil {
						t.Fatalf("Expected non-nil IfAddressMessage, got nil")
					}

					// Verify that the Family field is set correctly.
					if msg.Family != uint8(tc.family) {
						t.Errorf("Expected Family: %d, got: %d", tc.family, msg.Family)
					}

					// Verify that other fields are initialized to zero.
					if msg.Prefixlen != 0 {
						t.Errorf("Expected Prefixlen: 0, got: %d", msg.Prefixlen)
					}
					if msg.Flags != 0 {
						t.Errorf("Expected Flags: 0, got: %d", msg.Flags)
					}
					if msg.Scope != 0 {
						t.Errorf("Expected Scope: 0, got: %d", msg.Scope)
					}
					if msg.Index != 0 {
						t.Errorf("Expected Index: 0, got: %d", msg.Index)
					}
				})
			}
		})

		t.Run("len method", func(t *testing.T) {
			const expectedSizeInBytes = SizeOfIfAddressMessage
			var msg IfAddressMessage
			if msg.Len() != expectedSizeInBytes {
				t.Fatalf("Len() method did not return expected size")
			}
		})

		t.Run("serialize method", func(t *testing.T) {
			var (
				err          error
				testData     IfAddressMessage
				expectedData []byte
				actualData   []byte
			)
			t.Run("setup testData", func(t *testing.T) {
				testData.Family = AfInet
				testData.Prefixlen = uint8(0x10)
				testData.Flags = uint8(0x11)
				testData.Scope = uint8(0x12)
				testData.Index = convert.BytesToUint32([4]byte{0x23, 0x22, 0x21, 0x20})
			})
			t.Run("setup expectedData", func(t *testing.T) {
				expectedData = []byte{
					0x02,
					0x10,
					0x11,
					0x12,
					0x20, 0x21, 0x22, 0x23,
				}
			})
			t.Run("run serialize method", func(t *testing.T) {
				actualData, err = testData.Serialize()
				if err != nil {
					t.Fatal(err)
				}
			})
			t.Run("compare outcome", func(t *testing.T) {
				if !bytes.Equal(actualData, expectedData) {
					t.Fatalf("serialize method failed.\n"+
						"  actual:%v\n"+
						"expected:%v",
						actualData, expectedData)
				}
			})
		})

		t.Run("deserialize method", func(t *testing.T) {
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
				buf := make([]byte, SizeOfIfAddressMessage)

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
	})
}
