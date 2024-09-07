package core

import (
	"testing"
	"unsafe"
)

func TestCnMsgOp_struct(t *testing.T) {
	t.Run("test CnMsgOp fields", func(t *testing.T) {
		_ = CnMsgOp{
			CnMsg: CnMsg{
				ID: CbID{
					Idx: uint32(0),
					Val: uint32(0),
				},
				Seq:    uint32(0),
				Ack:    uint32(0),
				Length: uint16(0),
				Flags:  uint16(0),
			},
			Op: uint32(0),
		}
	})
	t.Run("size check", func(t *testing.T) {
		var o CnMsgOp
		t.Run("Size of CnMsgOp", func(t *testing.T) {
			if unsafe.Sizeof(o.CnMsg) != unsafe.Sizeof(CnMsg{}) {
				t.Fatalf("CnMsg size verification")
			}
			if unsafe.Sizeof(o.Op) != unsafe.Sizeof(uint32(0)) {
				t.Fatalf("Op size verification")
			}
		})
	})
}
