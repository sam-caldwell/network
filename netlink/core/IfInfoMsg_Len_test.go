package core

import (
	"testing"
)

func TestInfoMsg_Len(t *testing.T) {
	var msg IfInfoMsg
	if sz := msg.Len(); sz != IfInfoMsgSize {
		t.Errorf("msg.Len()=%d; want %d", sz, IfInfoMsgSize)
	}
}
