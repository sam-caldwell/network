package core

import (
	"testing"
	"unsafe"
)

func TestNetlinkMessageHeader_Size(t *testing.T) {

	const expectedSizeInBytes = 0x10

	if sz := int(unsafe.Sizeof(NetlinkMessageHeader{})); sz != expectedSizeInBytes {

		t.Fatalf("size mismatch: got %d, want %d", sz, expectedSizeInBytes)

	}

}
