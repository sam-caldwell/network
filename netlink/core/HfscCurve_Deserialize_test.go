package core

import (
	"encoding/binary"
	"testing"
)

func TestHfscCurveDeserialize(t *testing.T) {
	t.Run("valid input", func(t *testing.T) {
		// Prepare a byte slice with valid input (12 bytes total)
		buf := make([]byte, 12)

		// Populate the byte slice with values for m1, d, and m2 (using little-endian encoding)
		binary.LittleEndian.PutUint32(buf[0:4], 100)  // m1 = 100
		binary.LittleEndian.PutUint32(buf[4:8], 200)  // d = 200
		binary.LittleEndian.PutUint32(buf[8:12], 300) // m2 = 300

		// Call DeserializeHfscCurve
		curve := DeserializeHfscCurve(buf)

		// Check that the deserialized values match the expected values
		if curve.m1 != 100 {
			t.Errorf("Expected m1 to be 100, but got %d", curve.m1)
		}
		if curve.d != 200 {
			t.Errorf("Expected d to be 200, but got %d", curve.d)
		}
		if curve.m2 != 300 {
			t.Errorf("Expected m2 to be 300, but got %d", curve.m2)
		}
	})
	t.Run(ErrInputTooShort, func(t *testing.T) {
		// Prepare a byte slice with only 8 bytes (instead of 12)
		buf := make([]byte, 8)

		// Populate the byte slice
		binary.LittleEndian.PutUint32(buf[0:4], 100) // m1 = 100
		binary.LittleEndian.PutUint32(buf[4:8], 200) // d = 200

		// Expect a panic when trying to access beyond the buffer's length
		defer func() {
			if r := recover(); r == nil {
				t.Fatalf("Expected panic due to short buffer, but got none")
			}
		}()

		// Call DeserializeHfscCurve (this should panic)
		_ = DeserializeHfscCurve(buf)
	})
}
