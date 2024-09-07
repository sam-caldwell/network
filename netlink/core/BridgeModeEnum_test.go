package core

import (
	"testing"
	"unsafe"
)

func TestBridgeModeEnum(t *testing.T) {
	t.Run("size check", func(t *testing.T) {
		const expectedSizeInBytes = 1
		o := BridgeModeEnum(0)
		if unsafe.Sizeof(o) != expectedSizeInBytes {
			t.Fatalf("size check failed")
		}
	})
	t.Run("value check", func(t *testing.T) {
		if value := BridgeModeEnum(0); value != BridgeModeUnspec {
			t.Fatalf("value check failed 0")
		}
		if value := BridgeModeEnum(1); value != BridgeModeHairpin {
			t.Fatalf("value check failed 1")
		}
	})
}
