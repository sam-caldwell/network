package core

import (
	"bytes"
	"encoding/binary"
	"golang.org/x/sys/unix"
	"testing"
	"unsafe"
)

func TestIfaCacheInfo(t *testing.T) {
	//
	// type IfaCacheinfo struct {
	//    Prefered uint32
	//    Valid    uint32
	//    Cstamp   uint32
	//    Tstamp   uint32
	// }
	//
	t.Run("test SizeOfIfaCacheInfo constant", func(t *testing.T) {
		const expectedSizeInBytes = int(unsafe.Sizeof(IfaCacheInfo{}))
		if SizeOfIfaCacheInfo != int(unsafe.Sizeof(IfaCacheInfo{})) {
			t.Fatalf("expected SizeOfIfaCacheInfo mismatch")
		}
	})

	t.Run("test IfaCacheInfo size", func(t *testing.T) {
		const expectedSizeInBytes = 16
		if sz := int(unsafe.Sizeof(IfaCacheInfo{})); sz != expectedSizeInBytes {
			t.Fatalf("expected IfaCacheInfo size mismatch")
		}
	})

	t.Run("test Len() method", func(t *testing.T) {
		const expectedSizeInBytes = SizeOfIfaCacheInfo
		var o IfaCacheInfo
		if sz := o.Len(); sz != expectedSizeInBytes {
			t.Fatalf("Len() method failed")
		}
	})

	t.Run("test Serialize() method", func(t *testing.T) {
		testInput := IfaCacheInfo{
			IfaCacheinfo: unix.IfaCacheinfo{
				Prefered: uint32(01),
				Valid:    uint32(02),
				Cstamp:   uint32(03),
				Tstamp:   uint32(04),
			},
		}
		expected := []byte{
			0x01, 0x00, 0x00, 0x00,
			0x02, 0x00, 0x00, 0x00,
			0x03, 0x00, 0x00, 0x00,
			0x04, 0x00, 0x00, 0x00,
		}
		data, err := testInput.Serialize()
		if err != nil {
			t.Fatal(err)
		}
		if !bytes.Equal(data, expected) {
			t.Fatalf("Serialize() method failed")
		}
	})
	t.Run("test Deserialize() method", func(t *testing.T) {
		var actual IfaCacheInfo
		expected := IfaCacheInfo{
			IfaCacheinfo: unix.IfaCacheinfo{
				Prefered: uint32(01),
				Valid:    uint32(02),
				Cstamp:   uint32(03),
				Tstamp:   uint32(04),
			},
		}
		testInput := []byte{
			0x01, 0x00, 0x00, 0x00,
			0x02, 0x00, 0x00, 0x00,
			0x03, 0x00, 0x00, 0x00,
			0x04, 0x00, 0x00, 0x00,
		}
		if err := actual.Deserialize(testInput); err != nil {
			t.Fatal(err)
		}
		if actual.Prefered != expected.Prefered {
			t.Fatalf("Deserialize() method failed (Prefered)")
		}
		if actual.Valid != expected.Valid {
			t.Fatalf("Deserialize() method failed (Valid)")
		}
		if actual.Cstamp != expected.Cstamp {
			t.Fatalf("Deserialize() method failed (Cstamp)")
		}
		if actual.Tstamp != expected.Tstamp {
			t.Fatalf("Deserialize() method failed (Tstamp)")
		}
	})

	t.Run("Test IfaCacheInfo Deserializer function", func(t *testing.T) {
		// Subtest 1: Valid input
		t.Run("valid input", func(t *testing.T) {
			// Prepare a byte slice with valid input (16 bytes total)
			buf := make([]byte, SizeOfIfaCacheInfo)

			// Populate the byte slice with values for IfaCacheInfo fields (using little-endian encoding)
			binary.LittleEndian.PutUint32(buf[0:4], 100)   // Prefered = 100
			binary.LittleEndian.PutUint32(buf[4:8], 200)   // Valid = 200
			binary.LittleEndian.PutUint32(buf[8:12], 300)  // Cstamp = 300
			binary.LittleEndian.PutUint32(buf[12:16], 400) // Tstamp = 400

			// Call DeserializeIfaCacheInfo
			info, err := DeserializeIfaCacheInfo(buf)

			// Check for no errors
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			// Check the deserialized values
			if info.Prefered != 100 {
				t.Errorf("Expected Prefered to be 100, but got %d", info.Prefered)
			}
			if info.Valid != 200 {
				t.Errorf("Expected Valid to be 200, but got %d", info.Valid)
			}
			if info.Cstamp != 300 {
				t.Errorf("Expected Cstamp to be 300, but got %d", info.Cstamp)
			}
			if info.Tstamp != 400 {
				t.Errorf("Expected Tstamp to be 400, but got %d", info.Tstamp)
			}
		})

		// Subtest 2: Invalid input (too short)
		t.Run(ErrInputTooShort, func(t *testing.T) {
			// Prepare a byte slice with insufficient length
			buf := make([]byte, 12) // Less than the required 16 bytes

			// Call DeserializeIfaCacheInfo
			_, err := DeserializeIfaCacheInfo(buf)

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
}
