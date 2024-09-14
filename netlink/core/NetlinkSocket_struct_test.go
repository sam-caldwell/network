package core

import (
	"golang.org/x/sys/unix"
	"testing"
	"unsafe"
)

func TestNetlinkSocket_Struct(t *testing.T) {
	t.Run("size check", func(t *testing.T) {
		const expectedSizeInBytes = 36
		if sz := int(unsafe.Sizeof(NetlinkSocket{})); sz != expectedSizeInBytes {
			t.Fatalf("size mismatch: %d", sz)
		}
	})
	t.Run("field check", func(t *testing.T) {
		o := NetlinkSocket{
			fd:  int32(0),
			lsa: unix.SockaddrNetlink{},
		}
		o.Lock()
		o.Unlock()
	})
}
