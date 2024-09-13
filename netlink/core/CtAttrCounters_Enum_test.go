package core

import (
	"testing"
	"unsafe"
)

func TestCtAttrCounters_Enum(t *testing.T) {
	t.Run("Test CtAttrCounters type", func(t *testing.T) {
		t.Run("size check", func(t *testing.T) {
			const expectedSizeInBytes = 1
			var o = CtAttrCounters(0)
			if sz := unsafe.Sizeof(o); sz != expectedSizeInBytes {
				t.Fatal("size mismatch")
			}
		})
		t.Run("value check", func(t *testing.T) {
			type TestData struct {
				actual   CtAttrCounters
				expected CtAttrCounters
			}
			testData := []TestData{
				{actual: CtaCountersUnspec, expected: 0},
				{actual: CtaCountersPackets, expected: 1},
				{actual: CtaCountersBytes, expected: 2},
				{actual: CtaCounters32Packets, expected: 3},
				{actual: CtaCounters32Bytes, expected: 4},
				{actual: CtaCountersPad, expected: 5},
				{actual: CtaCountersMax, expected: 5},
			}
			for _, row := range testData {
				if row.actual != row.expected {
					t.Fatalf("value mismatch.  actual: %d, expected: %d", row.actual, row.expected)
				}
			}
		})
	})
}
