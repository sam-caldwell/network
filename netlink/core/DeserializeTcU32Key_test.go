package core

import (
	"encoding/binary"
	"testing"
)

// TestDeserializeTcU32Key tests the DeserializeTcU32Key function.
func TestDeserializeTcU32Key(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		// Prepare a byte slice for TcU32Key
		data := make([]byte, SizeOfTcU32Key)

		// Fill the byte slice with sample values for Mask, Val, Off, and OffMask
		binary.BigEndian.PutUint32(data[0:], 0x01020304)     // Mask (BigEndian)
		binary.BigEndian.PutUint32(data[4:], 0x05060708)     // Val (BigEndian)
		binary.LittleEndian.PutUint32(data[8:], 0x090A0B0C)  // Off (LittleEndian)
		binary.LittleEndian.PutUint32(data[12:], 0x0D0E0F10) // OffMask (LittleEndian)

		// Call DeserializeTcU32Key
		result, err := DeserializeTcU32Key(data)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		// Verify the result for Mask
		if result.Mask != 0x01020304 {
			t.Errorf("Expected Mask 0x01020304, got 0x%08x", result.Mask)
		}

		// Verify the result for Val
		if result.Val != 0x05060708 {
			t.Errorf("Expected Val 0x05060708, got 0x%08x", result.Val)
		}

		// Verify the result for Off
		if result.Off != 0x090A0B0C {
			t.Errorf("Expected Off 0x090A0B0C, got 0x%08x", result.Off)
		}

		// Verify the result for OffMask
		if result.OffMask != 0x0D0E0F10 {
			t.Errorf("Expected OffMask 0x0D0E0F10, got 0x%08x", result.OffMask)
		}
	})

	t.Run("tests DeserializeTcU32Key with short input.", func(t *testing.T) {
		// Prepare short data that is less than the expected size
		shortData := make([]byte, SizeOfTcU32Key-1)

		// Call DeserializeTcU32Key
		_, err := DeserializeTcU32Key(shortData)
		if err == nil {
			t.Fatal("Expected an error for short input, but got nil")
		}

		expectedErr := "byte slice too short"
		if err.Error()[:len(expectedErr)] != expectedErr {
			t.Errorf("Expected error message %q, but got %q", expectedErr, err.Error())
		}
	})
}
