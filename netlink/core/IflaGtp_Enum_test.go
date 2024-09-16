package core

import (
	"testing"
	"unsafe"
)

// TestIflaGtpEnum tests the size of the IflaGtp type and the values of its enumerated constants.
func TestIflaGtpEnum(t *testing.T) {
	// Verify the size of the IflaGtp type is 1 byte (uint8)
	if size := unsafe.Sizeof(IflaGtpUnspec); size != 1 {
		t.Errorf("Expected size of IflaGtp to be 1 byte, but got %d", size)
	}

	// Test the values of the constants
	tests := []struct {
		name  string
		value IflaGtp
		want  IflaGtp
	}{
		{"IflaGtpUnspec", IflaGtpUnspec, 0},
		{"IflaGtpFd0", IflaGtpFd0, 1},
		{"IflaGtpFd1", IflaGtpFd1, 2},
		{"IflaGtpPdpHashsize", IflaGtpPdpHashsize, 3},
		{"IflaGtpRole", IflaGtpRole, 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.value != tt.want {
				t.Errorf("Expected %s to have value %d, but got %d", tt.name, tt.want, tt.value)
			}
		})
	}
}
