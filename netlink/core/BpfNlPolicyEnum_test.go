package core

import (
	"testing"
	"unsafe"
)

func TestBpfNlPolicyEnum(t *testing.T) {
	t.Run("size check", func(t *testing.T) {
		const expectedSizeInBytes = 1
		o := BpfNlPolicyEnum(0)
		if unsafe.Sizeof(o) != expectedSizeInBytes {
			t.Fatalf("size check failed")
		}
	})
	t.Run("value check", func(t *testing.T) {
		if value := BpfNlPolicyEnum(0); value != LwtBpfUnspec {
			t.Fatalf("value check failed 0")
		}
		if value := BpfNlPolicyEnum(1); value != LwtBpfIn {
			t.Fatalf("value check failed 0")
		}
		if value := BpfNlPolicyEnum(2); value != LwtBpfOut {
			t.Fatalf("value check failed 0")
		}
		if value := BpfNlPolicyEnum(3); value != LwtBpfXmit {
			t.Fatalf("value check failed 0")
		}
		if value := BpfNlPolicyEnum(4); value != LwtBpfXmitHeadroom {
			t.Fatalf("value check failed 0")
		}
		if value := BpfNlPolicyEnum(4); value != LwtBpfMax {
			t.Fatalf("value check failed 0")
		}
	})
}
