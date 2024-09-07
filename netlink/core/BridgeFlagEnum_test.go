package core

import (
	"testing"
	"unsafe"
)

func TestBridgeFlagEnum(t *testing.T) {
	t.Run("size check", func(t *testing.T) {
		const expectedSizeInBytes = 1
		o := BridgeFlagEnum(0)
		if unsafe.Sizeof(o) != expectedSizeInBytes {
			t.Fatalf("size check failed")
		}
	})
	t.Run("value check", func(t *testing.T) {
		if value := BridgeFlagEnum(0); value != BridgeFlagsUnspec {
			t.Fatalf("value check failed 0")
		}
		if value := BridgeFlagEnum(1); value != BridgeFlagsMaster {
			t.Fatalf("value check failed 0")
		}
		if value := BridgeFlagEnum(2); value != BridgeFlagsSelf {
			t.Fatalf("value check failed 0")
		}
	})
}
