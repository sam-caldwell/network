package core

import (
	"math"
	"testing"
	"unsafe"
)

func TestDefaultChange_constant(t *testing.T) {
	if DefaultChange != math.MaxUint32 {
		t.Fatal("value check failed")
	}
	if unsafe.Sizeof(DefaultChange) != unsafe.Sizeof(uint32(0)) {
		t.Fatal("size check failed")
	}
}
