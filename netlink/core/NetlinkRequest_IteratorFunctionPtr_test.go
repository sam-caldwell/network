package core

import (
	"testing"
)

func TestNetlinkRequest_IteratorFunction(t *testing.T) {
	var f IteratorFunctionPtr

	f = func(msg []byte) bool {
		return false
	}
	if f([]byte{}) {
		t.Fatal("failed")
	}
}
