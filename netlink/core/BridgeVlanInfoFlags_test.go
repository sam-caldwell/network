package core

import (
	"testing"
	"unsafe"
)

func TestBridgeVlanInfoFlags(t *testing.T) {

	t.Run("Test Flag size", func(t *testing.T) {
		const expectedSizeInBytes = 2
		o := BridgeVlanInfo{}
		if sz := unsafe.Sizeof(o.Flags); sz != expectedSizeInBytes {
			t.Fatalf("size (Flags) mismatch: got %d, want %d", sz, expectedSizeInBytes)
		}
	})

	t.Run("Test BridgeVlanInfoFlags enum", func(t *testing.T) {
		t.Run("size test", func(t *testing.T) {
			const expectedSizeInBytes = 2
			o := BridgeVlanInfoFlags(0)
			if unsafe.Sizeof(o) != expectedSizeInBytes {
				t.Fatalf("size mismatch")
			}
		})
		t.Run("value test", func(t *testing.T) {
			testData := map[BridgeVlanInfoFlags]BridgeVlanInfoFlags{
				BridgeVlanInfoFlagsUnset: 0,
				BridgeVlanInfoMaster:     1,
				BridgeVlanInfoPvid:       2,
				BridgeVlanInfoUntagged:   4,
				BridgeVlanInfoRangeBegin: 8,
				BridgeVlanInfoRangeEnd:   16,
				BridgeVlanInfoBrentry:    32,
				BridgeVlanInfoOnlyOpts:   64,
			}
			for actual, expected := range testData {
				if actual != expected {
					t.Fatalf("value mismatch: expected %d, actual %d", expected, actual)
				}
			}
		})
	})
}
