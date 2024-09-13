package core

import (
	"golang.org/x/sys/unix"
	"testing"
)

func TestEncapType(t *testing.T) {
	tests := []struct {
		name           string
		ifInfoMsgType  uint16
		expectedResult string
	}{
		{"TypeEther", unix.ARPHRD_ETHER, "ether"},
		{"TypeLoopback", unix.ARPHRD_LOOPBACK, "loopback"},
		{"TypeNone", 65534, "none"},
		{"TypeVoid", 65535, "void"},
		{"UnknownType", 9999, "unknown9999"},
		{"TypeGeneric", 0, "generic"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ifInfoMsg := IfInfoMsg{
				IfInfomsg: unix.IfInfomsg{
					Type: tt.ifInfoMsgType,
				},
			}

			result := ifInfoMsg.EncapType()

			if result != tt.expectedResult {
				t.Errorf("EncapType() = %v, want %v", result, tt.expectedResult)
			}
		})
	}
}
