package core

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestCtAttrCountersEnum(t *testing.T) {
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
	t.Run("Test .FromInt() method", func(t *testing.T) {
		var c CtAttrCounters
		for i := range []int{0, 1, 2, 4, 128, 255} {

			if c.FromInt(i); int(c) != i {
				t.Errorf("CtAttrCounters.FromInt(0,0)=%d, want %d", c, i)
			}

		}
	})
	t.Run("Test .String() method", func(t *testing.T) {
		if v := CtAttrCounters(CtaCountersPackets); v.String() != "CTA_COUNTERS_PACKETS" {
			t.Fatalf("value (CtaCountersPackets) mismatch: %v", v.String())
		}
		if v := CtAttrCounters(CtaCountersBytes); v.String() != "CTA_COUNTERS_BYTES" {
			t.Fatalf("value (CtaCountersBytes) mismatch: %v", v.String())
		}
		if v := CtAttrCounters(CtaCounters32Packets); v.String() != "CTA_COUNTERS32_PACKETS" {
			t.Fatalf("value (CtaCounters32Packets) mismatch: %v", v.String())
		}
		if v := CtAttrCounters(CtaCounters32Bytes); v.String() != "CTA_COUNTERS32_BYTES" {
			t.Fatalf("value (CtaCounters32Bytes) mismatch: %v", v.String())
		}
		if v := CtAttrCounters(CtaCountersPad); v.String() != "CTA_COUNTERS_PAD" {
			t.Fatalf("value (CtaCountersPad) mismatch: %v", v.String())
		}
	})
	t.Run("test .ToInt() method", func(t *testing.T) {
		for v := range []int{0, 1, 128, 255} {
			t.Run(fmt.Sprintf("%s(%d)", t.Name(), v), func(t *testing.T) {
				o := CtAttrCounters(v)
				if int(o) != v {
					t.Fatalf("value mismatch on %d", v)
				}
			})
		}
	})
}
