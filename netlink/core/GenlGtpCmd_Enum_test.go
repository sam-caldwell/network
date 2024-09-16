package core

import (
	"testing"
	"unsafe"
)

// TestGenlGtpCmdEnum tests the size of the GenlGtpCmd type
// and the values of its enumerated constants.
func TestGenlGtpCmdEnum(t *testing.T) {
	// Verify the size of the GenlGtpCmd type is 1 byte (uint8)
	if size := unsafe.Sizeof(GenlGtpCmdNewpdp); size != 1 {
		t.Errorf("Expected size of GenlGtpCmd to be 1 byte, but got %d", size)
	}

	// Test the values of the constants
	tests := []struct {
		name  string
		value GenlGtpCmd
		want  GenlGtpCmd
	}{
		{"GenlGtpCmdNewpdp", GenlGtpCmdNewpdp, 0},
		{"GenlGtpCmdDelpdp", GenlGtpCmdDelpdp, 1},
		{"GenlGtpCmdGetpdp", GenlGtpCmdGetpdp, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.value != tt.want {
				t.Errorf("Expected %s to have value %d, but got %d", tt.name, tt.want, tt.value)
			}
		})
	}
}
