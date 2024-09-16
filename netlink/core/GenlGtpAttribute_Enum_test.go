package core

import (
	"testing"
	"unsafe"
)

// TestGenlGtpAttributeEnum tests the size of the GenlGtpAttributeEnum type
// and the values of its enumerated constants.
func TestGenlGtpAttributeEnum(t *testing.T) {
	// Verify the size of the GenlGtpAttributeEnum type is 1 byte (uint8)
	if size := unsafe.Sizeof(GenlGtpAttrUnspec); size != 1 {
		t.Errorf("Expected size of GenlGtpAttributeEnum to be 1 byte, but got %d", size)
	}

	// Test the values of the constants
	tests := []struct {
		name  string
		value GenlGtpAttributeEnum
		want  GenlGtpAttributeEnum
	}{
		{"GenlGtpAttrUnspec", GenlGtpAttrUnspec, 0},
		{"GenlGtpAttrLink", GenlGtpAttrLink, 1},
		{"GenlGtpAttrVersion", GenlGtpAttrVersion, 2},
		{"GenlGtpAttrTid", GenlGtpAttrTid, 3},
		{"GenlGtpAttrPeerAddress", GenlGtpAttrPeerAddress, 4},
		{"GenlGtpAttrMsAddress", GenlGtpAttrMsAddress, 5},
		{"GenlGtpAttrFlow", GenlGtpAttrFlow, 6},
		{"GenlGtpAttrNetNsFd", GenlGtpAttrNetNsFd, 7},
		{"GenlGtpAttrITei", GenlGtpAttrITei, 8},
		{"GenlGtpAttrOTei", GenlGtpAttrOTei, 9},
		{"GenlGtpAttrPad", GenlGtpAttrPad, 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.value != tt.want {
				t.Errorf("Expected %s to have value %d, but got %d", tt.name, tt.want, tt.value)
			}
		})
	}
}
