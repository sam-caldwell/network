package core

import (
	"testing"
	"unsafe"
)

func TestNetlinkRequestSize(t *testing.T) {
	if NetlinkRequestSize != int(unsafe.Sizeof(NetlinkRequest{})) {
		t.Fatalf("NetlinkRequestSize failed")
	}
}
