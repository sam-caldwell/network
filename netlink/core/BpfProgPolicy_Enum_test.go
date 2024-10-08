package core

import (
	"testing"
	"unsafe"
)

func TestBpfProgPolicyEnum(t *testing.T) {
	t.Run("size check", func(t *testing.T) {
		const expectedSizeInBytes = 1
		o := BpfProgPolicy(0)
		if unsafe.Sizeof(o) != expectedSizeInBytes {
			t.Fatalf("size check failed")
		}
	})
	t.Run("value check", func(t *testing.T) {
		if value := BpfProgPolicy(0); value != LwtBpfProgUnspec {
			t.Fatalf("value check failed 0")
		}
		if value := BpfProgPolicy(1); value != LwtBpfProgFd {
			t.Fatalf("value check failed 1")
		}
		if value := BpfProgPolicy(2); value != LwtBpfProgName {
			t.Fatalf("value check failed 2")
		}
		if value := BpfProgPolicy(2); value != LwtBpfProgMax {
			t.Fatalf("value check failed 2")
		}
	})
}
