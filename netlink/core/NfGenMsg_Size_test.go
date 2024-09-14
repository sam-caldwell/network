package core

import (
	"testing"
	"unsafe"
)

func TestNfGenMsg_Size(t *testing.T) {
	if sz := int(unsafe.Sizeof(NfGenMsg{})); sz != NfGenMsgSize {
		t.Fatalf("NfGenMsgSize got %d, want %d", sz, NfGenMsgSize)
	}
}
