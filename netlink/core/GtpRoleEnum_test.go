package core

import (
	"testing"
	"unsafe"
)

// TestGtpRoleEnum tests the size of the GtpRoleEnum type and the values of its enumerated constants.
func TestGtpRoleEnum(t *testing.T) {
	// Verify the size of the GtpRoleEnum type is 1 byte (uint8)
	if size := unsafe.Sizeof(GtpRoleGgsn); size != 1 {
		t.Errorf("Expected size of GtpRoleEnum to be 1 byte, but got %d", size)
	}

	// Test the values of the constants
	tests := []struct {
		name  string
		value GtpRoleEnum
		want  GtpRoleEnum
	}{
		{"GtpRoleGgsn", GtpRoleGgsn, 0},
		{"GtpRoleSgsn", GtpRoleSgsn, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.value != tt.want {
				t.Errorf("Expected %s to have value %d, but got %d", tt.name, tt.want, tt.value)
			}
		})
	}
}
