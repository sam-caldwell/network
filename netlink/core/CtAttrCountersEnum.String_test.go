package core

import (
	"testing"
)

func TestCtAttrCountersEnum_String(t *testing.T) {
	if v := CtAttrCountersEnum(CtaCountersUnspec); v.String() != "CtaCountersUnspec" {
		t.Fatalf("value (CtaCountersUnspec) mismatch: %v", v.String())
	}
	if v := CtAttrCountersEnum(CtaCountersPackets); v.String() != "CtaCountersPackets" {
		t.Fatalf("value (CtaCountersPackets) mismatch: %v", v.String())
	}
	if v := CtAttrCountersEnum(CtaCountersBytes); v.String() != "CtaCountersBytes" {
		t.Fatalf("value (CtaCountersBytes) mismatch: %v", v.String())
	}
	if v := CtAttrCountersEnum(CtaCounters32Packets); v.String() != "CtaCounters32Packets" {
		t.Fatalf("value (CtaCounters32Packets) mismatch: %v", v.String())
	}
	if v := CtAttrCountersEnum(CtaCounters32Bytes); v.String() != "CtaCounters32Bytes" {
		t.Fatalf("value (CtaCounters32Bytes) mismatch: %v", v.String())
	}
	if v := CtAttrCountersEnum(CtaCountersPad); v.String() != "CtaCountersPad" {
		t.Fatalf("value (CtaCountersPad) mismatch: %v", v.String())
	}
}
