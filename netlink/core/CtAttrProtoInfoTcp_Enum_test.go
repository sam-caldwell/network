package core

import (
	"testing"
	"unsafe"
)

func TestCtAttrProtoInfoTcpEnum(t *testing.T) {
	t.Run("size check", func(t *testing.T) {
		const expectedSizeInBytes = 1
		if sz := unsafe.Sizeof(CtAttrProtoInfoTcp(0)); sz != expectedSizeInBytes {
			t.Errorf("Expected size of %d, got %d.", expectedSizeInBytes, sz)
		}
	})
	t.Run("value check", func(t *testing.T) {
		testData := map[CtAttrProtoInfoTcp]CtAttrProtoInfoTcp{
			0: CtaProtoInfoTcpUnspec,
			1: CtaProtoInfoTcpState,
			2: CtaProtoInfoTcpWscaleOriginal,
			3: CtaProtoInfoTcpWscaleReply,
			4: CtaProtoInfoTcpFlagsOriginal,
			5: CtaProtoInfoTcpFlagsReply,
		}
		for k, v := range testData {
			if k != v {
				t.Fatalf("value mismatch. actual: %d, expected: %d", k, v)
			}
		}
	})
}
