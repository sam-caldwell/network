package core

import (
	"testing"
	"unsafe"
)

func TestBridgeVlanInfo(t *testing.T) {

	t.Run("check structure", func(t *testing.T) {
		_ = BridgeVlanInfo{
			Flags: BridgeVlanInfoFlags(0),
			Vid:   VlanIdType(0),
		}
	})

	t.Run("test struct size", func(t *testing.T) {
		const expectedSizeInBytes = 4
		o := BridgeVlanInfo{}
		if sz := unsafe.Sizeof(o); sz != expectedSizeInBytes {
			t.Fatalf("struct size mismatch: got %d, want %d", sz, expectedSizeInBytes)
		}
	})

	t.Run("test Vid size", func(t *testing.T) {
		const expectedSizeInBytes = 2
		o := BridgeVlanInfo{}
		if sz := unsafe.Sizeof(o.Vid); sz != expectedSizeInBytes {
			t.Fatalf("size (Vid) mismatch: got %d, want %d", sz, expectedSizeInBytes)
		}
	})

}
