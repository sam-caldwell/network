package core

import (
	"testing"
	"unsafe"
)

func TestCnMsg_Len(t *testing.T) {
	var o CnMsgOp
	if o.Len() != int(unsafe.Sizeof(CnMsgOp{})) {
		t.Fatalf("CnMsgOp.Len() method size mismatch")
	}
}
