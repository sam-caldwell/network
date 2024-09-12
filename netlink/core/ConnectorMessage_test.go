package core

import (
	"testing"
	"unsafe"
)

func TestCnMsg_struct(t *testing.T) {
	t.Run("field check", func(t *testing.T) {
		_ = ConnectorMessage{
			ID: CbID{
				Idx: uint32(0),
				Val: uint32(0),
			},
			Seq:    uint32(0),
			Ack:    uint32(0),
			Length: uint16(0),
			Flags:  uint16(0),
		}
	})
	t.Run("size check", func(t *testing.T) {
		var o ConnectorMessage
		t.Run("size of ID field", func(t *testing.T) {
			if unsafe.Sizeof(o.ID) != unsafe.Sizeof(CbID{}) {
				t.Fatalf("ID value size mismatch")
			}
		})
		t.Run("size of Seq field", func(t *testing.T) {
			if unsafe.Sizeof(o.Seq) != 4 {
				t.Fatalf("Seq value size mismatch")
			}
		})
		t.Run("size of Ack field", func(t *testing.T) {
			if unsafe.Sizeof(o.Ack) != 4 {
				t.Fatalf("Ack value size mismatch")
			}
		})
		t.Run("size of Length field", func(t *testing.T) {
			if unsafe.Sizeof(o.Length) != 2 {
				t.Fatalf("Length value size mismatch")
			}
		})
		t.Run("size of Flags field", func(t *testing.T) {
			if unsafe.Sizeof(o.Flags) != 2 {
				t.Fatalf("Flags value size mismatch")
			}
		})
	})
}
