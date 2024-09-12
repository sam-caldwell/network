package core

import (
	"testing"
	"unsafe"
)

// TestGreEnum tests the size of the GreEnum type and the values of its enumerated constants.
func TestGreEnum(t *testing.T) {
	// Verify the size of the GreEnum type is 2 bytes (uint16)
	if size := unsafe.Sizeof(GreCsum); size != 2 {
		t.Errorf("Expected size of GreEnum to be 2 bytes, but got %d", size)
	}

	// Test the values of the constants
	tests := []struct {
		name  string
		value GreEnum
		want  GreEnum
	}{
		{"GreCsum", GreCsum, 0x8000},
		{"GreRouting", GreRouting, 0x4000},
		{"GreKey", GreKey, 0x2000},
		{"GreSeq", GreSeq, 0x1000},
		{"GreStrict", GreStrict, 0x0800},
		{"GreRec", GreRec, 0x0700},
		{"GreAck", GreAck, 0x0080},
		{"GreFlags", GreFlags, 0x00F8},
		{"GreVersion", GreVersion, 0x0007},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.value != tt.want {
				t.Errorf("Expected %s to have value 0x%x, but got 0x%x", tt.name, tt.want, tt.value)
			}
		})
	}
}
