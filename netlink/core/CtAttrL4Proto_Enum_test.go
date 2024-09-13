package core

import (
	"testing"
	"unsafe"
)

func TestCtAttrL4Proto_Enum(t *testing.T) {
	t.Run("size check", func(t *testing.T) {
		const expectedSizeInBytes = 1
		o := CtAttrL4Proto(0)
		if sz := unsafe.Sizeof(o); sz != expectedSizeInBytes {
			t.Fatalf("size mismatch failed")
		}
	})
	t.Run("value check", func(t *testing.T) {
		if o := CtaProtoNum; int(o) != 1 {
			t.Fatalf("CtaProtNum mismatch")
		}
		if o := CtaProtoSrcPort; int(o) != 2 {
			t.Fatalf("CtaProtNum mismatch")
		}
		if o := CtaProtoDstPort; int(o) != 3 {
			t.Fatalf("CtaProtNum mismatch")
		}
	})
}
