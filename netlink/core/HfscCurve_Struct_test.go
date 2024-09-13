package core

import (
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
}
