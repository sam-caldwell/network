package core

import (
	"testing"
)

func TestCtAttrCountersEnum_String(t *testing.T) {
	if v := CtAttrCountersEnum(CtaCountersPackets); v.String() != "CTA_COUNTERS_PACKETS" {
		t.Fatalf("value (CtaCountersPackets) mismatch: %v", v.String())
	}
	if v := CtAttrCountersEnum(CtaCountersBytes); v.String() != "CTA_COUNTERS_BYTES" {
		t.Fatalf("value (CtaCountersBytes) mismatch: %v", v.String())
	}
	if v := CtAttrCountersEnum(CtaCounters32Packets); v.String() != "CTA_COUNTERS32_PACKETS" {
		t.Fatalf("value (CtaCounters32Packets) mismatch: %v", v.String())
	}
	if v := CtAttrCountersEnum(CtaCounters32Bytes); v.String() != "CTA_COUNTERS32_BYTES" {
		t.Fatalf("value (CtaCounters32Bytes) mismatch: %v", v.String())
	}
	if v := CtAttrCountersEnum(CtaCountersPad); v.String() != "CTA_COUNTERS_PAD" {
		t.Fatalf("value (CtaCountersPad) mismatch: %v", v.String())
	}
}
