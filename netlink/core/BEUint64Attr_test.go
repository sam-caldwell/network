package core

import (
	"testing"
)

func TestBEUint64Attr(t *testing.T) {
	tests := []struct {
		name string
		v    uint64
		want []byte
	}{
		{
			name: "Zero value",
			v:    0,
			want: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		},
		{
			name: "Small number",
			v:    256, // 0x0000000000000100 in hex
			want: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00},
		},
		{
			name: "Large number",
			v:    18446744073709551615, // 0xFFFFFFFFFFFFFFFF in hex (max uint64)
			want: []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF},
		},
		{
			name: "Middle value",
			v:    1311768465173141112, // 0x123456789ABCDEF8 in hex
			want: []byte{0x12, 0x34, 0x56, 0x78, 0x9A, 0xBC, 0xDE, 0xF8},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BEUint64Attr(tt.v)
			if len(got) != 8 {
				t.Fatalf("expected length of result to be 8, got %v", len(got))
			}
			for i := range tt.want {
				if got[i] != tt.want[i] {
					t.Errorf("BEUint64Attr() = %v, want %v", got, tt.want)
					break
				}
			}
		})
	}
}
