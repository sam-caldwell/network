package core

import (
	"testing"
)

func TestBEUint16Attr(t *testing.T) {
	tests := []struct {
		name string
		v    uint16
		want []byte
	}{
		{
			name: "Zero value",
			v:    0,
			want: []byte{0x00, 0x00},
		},
		{
			name: "Small number",
			v:    256, // 0x0100 in hex
			want: []byte{0x01, 0x00},
		},
		{
			name: "Max uint16",
			v:    65535, // 0xFFFF in hex
			want: []byte{0xFF, 0xFF},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BEUint16Attr(tt.v)
			if len(got) != 2 {
				t.Fatalf("expected length of result to be 2, got %v", len(got))
			}
			if got[0] != tt.want[0] || got[1] != tt.want[1] {
				t.Errorf("BEUint16Attr() = %v, want %v", got, tt.want)
			}
		})
	}
}
