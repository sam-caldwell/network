package core

import (
	"testing"
)

func TestCtAttrCounters_String(t *testing.T) {
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
}
