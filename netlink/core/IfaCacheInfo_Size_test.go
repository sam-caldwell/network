package core

import (
	"testing"
	"unsafe"
)

func TestIfaCacheInfo_Size(t *testing.T) {

	const expectedSizeInBytes = int(unsafe.Sizeof(IfaCacheInfo{}))
	if IfaCacheInfoSize != int(unsafe.Sizeof(IfaCacheInfo{})) {
		t.Fatalf("expected IfaCacheInfoSize mismatch")
	}
}
