package core

import (
	"testing"
)

func TestBEUint16Attr(t *testing.T) {
	t.Run("16-bit", func(t *testing.T) {
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
				got := BEUintAttr(tt.v)
				if len(got) != 2 {
					t.Fatalf("expected length of result to be 2, got %v", len(got))
				}
				if got[0] != tt.want[0] || got[1] != tt.want[1] {
					t.Errorf("BEUint16Attr() = %v, want %v", got, tt.want)
				}
			})
		}
	})
	t.Run("32-bit", func(t *testing.T) {
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
				got := BEUintAttr(tt.v)
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
	})
	t.Run("64-bit", func(t *testing.T) {
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
				v:    0x123456789ABCDEF8, // Correct representation of 1311768465173141112
				want: []byte{0x12, 0x34, 0x56, 0x78, 0x9A, 0xBC, 0xDE, 0xF8},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got := BEUintAttr(tt.v)
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
	})
}
