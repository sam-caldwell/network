package core

import (
	"golang.org/x/sys/unix"
	"testing"
)

func TestNewIfInfomsg(t *testing.T) {
	tests := []struct {
		name           string
		family         InterfaceFamily
		expectedFamily uint8
	}{
		{"AF_INET", AfInet, uint8(unix.AF_INET)},
		{"AF_INET6", AfInet6, uint8(unix.AF_INET6)},
		{"AF_UNSPEC", AfUnspec, uint8(unix.AF_UNSPEC)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ifInfoMsg := NewIfInfomsg(tt.family)

			if ifInfoMsg.Family != tt.expectedFamily {
				t.Errorf("NewIfInfomsg() Family = %v, want %v", ifInfoMsg.Family, tt.expectedFamily)
			}
		})
	}
}
