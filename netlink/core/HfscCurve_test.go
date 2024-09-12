package core

import (
	"encoding/binary"
	"testing"
	"unsafe"
)

func TestCurve(t *testing.T) {
	t.Run("test struct", func(t *testing.T) {
		t.Run("test struct fields", func(t *testing.T) {
			_ = Curve{
				m1: uint32(0),
				d:  uint32(0),
				m2: uint32(0),
			}
		})
		t.Run("default value test", func(t *testing.T) {
			var c Curve
			if c.m1 != 0 {
				t.Fatalf("m1 expects 0")
			}
			if c.d != 0 {
				t.Fatalf("d expects 0")
			}
			if c.m2 != 0 {
				t.Fatalf("m2 expects 0")
			}
		})

		t.Run("size check", func(t *testing.T) {
			const expectedSizeInBytes = 12
			if actual := unsafe.Sizeof(Curve{}); actual != expectedSizeInBytes {
				t.Errorf("Curve size mismatch. got %d, want %d", actual, expectedSizeInBytes)
			}
		})

		t.Run("field (m1) size check", func(t *testing.T) {
			const expectedSizeInBytes = 4
			var c Curve
			if sz := unsafe.Sizeof(c.m1); sz != expectedSizeInBytes {
				t.Fatalf("m1 size mismatch. got %d, want %d", sz, expectedSizeInBytes)
			}
		})

		t.Run("field (d) size check", func(t *testing.T) {
			const expectedSizeInBytes = 4
			var c Curve
			if sz := unsafe.Sizeof(c.d); sz != expectedSizeInBytes {
				t.Fatalf("d size mismatch. got %d, want %d", sz, expectedSizeInBytes)
			}
		})

		t.Run("field (m2) size check", func(t *testing.T) {
			const expectedSizeInBytes = 4
			var c Curve
			if sz := unsafe.Sizeof(c.m2); sz != expectedSizeInBytes {
				t.Fatalf("m2 size mismatch. got %d, want %d", sz, expectedSizeInBytes)
			}
		})
	})
	t.Run("Test .Attrs() method", func(t *testing.T) {
		c := Curve{
			m1: 1,
			d:  2,
			m2: 3,
		}
		if m1, d, m2 := c.Attrs(); m1 != 1 || d != 2 || m2 != 3 {
			t.Fatalf("attrs failed")
		}
	})
	t.Run("test .Set() method", func(t *testing.T) {
		var c Curve
		t.Run("Pre-test", func(t *testing.T) {
			if c.m1 != 0 || c.m2 != 0 || c.d != 0 {
				t.Fatal("initial state is wrong")
			}
		})
		t.Run("Use .Set() method", func(t *testing.T) {
			c.Set(1, 2, 3)
		})
		t.Run("verify outcome", func(t *testing.T) {
			if c.m1 != 1 {
				t.Fatalf("outcome (%d) verification failed", c.m1)
			}
			if c.d != 2 {
				t.Fatalf("outcome (%d) verification failed", c.d)
			}
			if c.m2 != 3 {
				t.Fatalf("outcome (%d) verification failed", c.m2)
			}
		})
	})
	t.Run("DeserializeHfscCurve", func(t *testing.T) {
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
	})
}
