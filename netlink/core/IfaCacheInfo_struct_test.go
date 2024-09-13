package core

import (
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
	t.Run("test IfaCacheinfo size", func(t *testing.T) {
		const expectedSizeInBytes = 16
		if sz := int(unsafe.Sizeof(IfaCacheInfo{})); sz != expectedSizeInBytes {
			t.Fatalf("expected IfaCacheInfo size mismatch")
		}
	})
	t.Run("test IfaCacheinfo type", func(t *testing.T) {
		_ = IfaCacheInfo{
			unix.IfaCacheinfo{},
		}
	})

}
