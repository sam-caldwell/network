package core

import (
	"golang.org/x/sys/unix"
	"testing"
	"unsafe"
)

func TestInfoMsg_Struct(t *testing.T) {
	t.Run("IfInfoMsg struct", func(t *testing.T) {
		t.Run("sizeOfIfInfoMsg check", func(t *testing.T) {
			if IfInfoMsgSize != 16 {
				t.Fatal("IfInfoMsgSize mismatch")
			}
		})
		t.Run("check size", func(t *testing.T) {
			const expectedSizeInBytes = IfInfoMsgSize
			if sz := int(unsafe.Sizeof(IfInfoMsg{})); sz != expectedSizeInBytes {
				t.Fatal("size mismatch")
			}
		})
		t.Run("field check", func(t *testing.T) {
			_ = IfInfoMsg{
				unix.IfInfomsg{
					Family: uint8(0),
					Type:   uint16(0),
					Index:  int32(0),
					Flags:  uint32(0),
					Change: uint32(0),
				},
			}
		})
	})
}
