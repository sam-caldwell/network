package core

import (
	"testing"
	"unsafe"
)

func TestCbID(t *testing.T) {
	t.Run("size field verification", func(t *testing.T) {
		_ = CbID{
			Idx: uint32(0),
			Val: uint32(0),
		}
	})
	t.Run("size check", func(t *testing.T) {
		const expectedSizeInBytes = 8
		var o CbID
		if unsafe.Sizeof(o) != expectedSizeInBytes {
			t.Fatalf("size check (struct) failed")
		}
	})
	t.Run("size check (Idx)", func(t *testing.T) {
		const expectedSizeInBytes = 4
		var o CbID
		if unsafe.Sizeof(o.Idx) != expectedSizeInBytes {
			t.Fatalf("size check (Idx) failed")
		}
	})
	t.Run("size check (Val)", func(t *testing.T) {
		const expectedSizeInBytes = 4
		var o CbID
		if unsafe.Sizeof(o.Val) != expectedSizeInBytes {
			t.Fatalf("size check (Val) failed")
		}
	})
}
