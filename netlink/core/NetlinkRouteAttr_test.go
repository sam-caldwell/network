package core

import (
	"testing"
	"unsafe"
)

func TestNetlinkRouteAttr(t *testing.T) {
	t.Run("test structure", func(t *testing.T) {
		t.Run("struct size check", func(t *testing.T) {
			const expectedSizeInBytes = int(unsafe.Sizeof(RtAttr{}) + unsafe.Sizeof([]byte{}))
			if sz := int(unsafe.Sizeof(NetlinkRouteAttr{})); sz != expectedSizeInBytes {
				t.Errorf("wrong size, expected %d, got %d", expectedSizeInBytes, sz)
			}
		})
		t.Run("struct fields", func(t *testing.T) {
			_ = NetlinkRouteAttr{
				Attr:  RtAttr{},
				Value: []byte{},
			}
		})
	})
}
