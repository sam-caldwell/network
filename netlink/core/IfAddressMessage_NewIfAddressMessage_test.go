package core

import (
	"golang.org/x/sys/unix"
	"testing"
)

func TestNewIfAddressMessage(t *testing.T) {
	tests := []struct {
		name           string
		family         InterfaceFamily
		expectedFamily uint8
	}{
		{
			name:           "AF_UNSPEC",
			family:         AfUnspec,
			expectedFamily: uint8(unix.AF_UNSPEC),
		},
		{
			name:           "AF_INET",
			family:         AfInet,
			expectedFamily: uint8(unix.AF_INET),
		},
		{
			name:           "AF_INET6",
			family:         AfInet6,
			expectedFamily: uint8(unix.AF_INET6),
		},
		{
			name:           "AF_APPLETALK",
			family:         AfAppleTalk,
			expectedFamily: uint8(unix.AF_APPLETALK),
		},
		{
			name:           "AF_X25",
			family:         AfX25,
			expectedFamily: uint8(unix.AF_X25),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a new IfAddressMessage with the specified family
			ifAddrMsg := NewIfAddressMessage(tt.family)

			// Verify that the Family field in the IfAddrmsg struct is set correctly
			if ifAddrMsg.IfAddrmsg.Family != tt.expectedFamily {
				t.Errorf("Expected Family: %d, got: %d", tt.expectedFamily, ifAddrMsg.IfAddrmsg.Family)
			}
		})
	}
}
