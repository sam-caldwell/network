package core

import (
	"testing"
	"unsafe"
)

func TestNextSequenceNumber(t *testing.T) {
	const expectedSizeInBytes = int(unsafe.Sizeof(uint32(0)))
	if sz := int(unsafe.Sizeof(nextSequenceNumber)); sz != expectedSizeInBytes {
		t.Fatalf("Wrong size of next sequence number: %d", sz)
	}
}
