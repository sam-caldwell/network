package core

import (
	"testing"
	"unsafe"
)

func TestCtAttrIpEnum(t *testing.T) {
	t.Run("size check", func(t *testing.T) {
		const expectedSizeInBytes = 1
		var o CtAttrIpEnum
		if sz := unsafe.Sizeof(o); sz != expectedSizeInBytes {
			t.Fatalf("size mismatch")
		}
	})
	t.Run("value check", func(t *testing.T) {
		testData := map[int]CtAttrIpEnum{
			0: CtAttrIpEnum(CtaIpUnspec),
			1: CtAttrIpEnum(CtaIpV4Src),
			2: CtAttrIpEnum(CtaIpV4Dst),
			3: CtAttrIpEnum(CtaIpV6Src),
			4: CtAttrIpEnum(CtaIpV6Dst),
		}
		for k, v := range testData {
			if k != int(v) {
				t.Fatalf("value mismatch.  Actual: %d, Expected: %d", k, v)
			}
		}
	})
}
