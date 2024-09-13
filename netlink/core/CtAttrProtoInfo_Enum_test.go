package core

import (
	"testing"
	"unsafe"
)

func TestCtAttrProtoInfoEnum(t *testing.T) {
	t.Run("size check", func(t *testing.T) {
		const expectedSizeInBytes = 1
		if sz := unsafe.Sizeof(CtAttrProtoInfo(0)); sz != expectedSizeInBytes {
			t.Fatal("size mismatch")
		}
	})

	t.Run("value check", func(t *testing.T) {
		if v := CtAttrProtoInfo(0); v != CtaProtoInfoUnspec {
			t.Fatalf("CtaProtoInfoUnspec mismatch")
		}
		if v := CtAttrProtoInfo(1); v != CtaProtoInfoTcp {
			t.Fatalf("CtaProtoInfoTcp mismatch")
		}
		if v := CtAttrProtoInfo(2); v != CtaProtoInfoDccp {
			t.Fatalf("CtaProtoInfoDccp mismatch")
		}
		if v := CtAttrProtoInfo(3); v != CtaProtoInfoSctp {
			t.Fatalf("CtaProtoInfoSctp mismatch")
		}
	})

}
