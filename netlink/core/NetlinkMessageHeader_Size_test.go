package core

import (
	"golang.org/x/sys/unix"
	"testing"
	"unsafe"
)

func TestNetlinkMessageHeader_Size(t *testing.T) {
	const expectedSizeInBytes = 16
	if sz := int(unsafe.Sizeof(unix.NlMsghdr{})); sz != expectedSizeInBytes {
		t.Fatalf("size mismatch: got %d, want %d", sz, expectedSizeInBytes)
	}
}
