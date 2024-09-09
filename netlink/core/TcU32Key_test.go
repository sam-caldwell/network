package core

import (
	"testing"
	"unsafe"
)

func TestTcU32KeyStruct(t *testing.T) {
	t.Run("field check", func(t *testing.T) {
		_ = TcU32Key{
			Mask:    uint32(0),
			Val:     uint32(0),
			Off:     int32(0),
			OffMask: int32(0),
		}
	})
	t.Run("Size check", func(t *testing.T) {
		if sz := int(unsafe.Sizeof(TcU32Key{})); sz != 16 {
			t.Fatal("size check failed")
		}
	})
}
