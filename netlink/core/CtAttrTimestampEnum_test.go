package core

import (
	"testing"
	"unsafe"
)

func TestCtAttrTimestampEnum(t *testing.T) {
	t.Run("size check", func(t *testing.T) {
		const expectedSizeInBytes = 8
		if sz := unsafe.Sizeof(CtAttrTimestampEnum(0)); sz != expectedSizeInBytes {
			t.Fatal("size check failed")
		}
	})
	t.Run("value check", func(t *testing.T) {
		testData := map[CtAttrTimestampEnum]CtAttrTimestampEnum{
			1: CtaTimestampStart,
			2: CtaTimestampStop,
		}
		for k, v := range testData {
			if k != v {
				t.Fatalf("value check mismatch.  Actual: %d, Expected: %d", k, v)
			}
		}
	})
}
