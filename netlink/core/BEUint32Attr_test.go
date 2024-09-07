package core

import (
	"testing"
)

func TestBEUint32Attr(t *testing.T) {
	tests := []struct {
		name string
		v    uint32
		want []byte
	}{
		{
			name: "Zero value",
			v:    0,
			want: []byte{0x00, 0x00, 0x00, 0x00},
		},
		{
			name: "Small number",
			v:    256, // 0x00000100 in hex
			want: []byte{0x00, 0x00, 0x01, 0x00},
		},
		{
			name: "Large number",
			v:    4294967295, // 0xFFFFFFFF in hex (max uint32)
			want: []byte{0xFF, 0xFF, 0xFF, 0xFF},
		},
		{
			name: "Middle value",
			v:    305419896, // 0x12345678 in hex
			want: []byte{0x12, 0x34, 0x56, 0x78},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BEUint32Attr(tt.v)
			if len(got) != 4 {
				t.Fatalf("expected length of result to be 4, got %v", len(got))
			}
			for i := range tt.want {
				if got[i] != tt.want[i] {
					t.Errorf("BEUint32Attr() = %v, want %v", got, tt.want)
					break
				}
			}
		})
	}
}
