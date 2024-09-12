package core

import (
	"testing"
	"unsafe"
)

func TestLwTunnelEnum(t *testing.T) {
	t.Run("TestSize", func(t *testing.T) {
		const expectedSizeInBytes = 1
		var tunnelAttr LwTunnelEnum

		if size := unsafe.Sizeof(tunnelAttr); size != uintptr(expectedSizeInBytes) {
			t.Errorf("Expected size %d but got %d", expectedSizeInBytes, size)
		}
	})

	t.Run("TestValues", func(t *testing.T) {
		tests := []struct {
			name     string
			value    LwTunnelEnum
			expected LwTunnelEnum
		}{
			{"LwtunnelIp6Unspec", LwtunnelIp6Unspec, 0},
			{"LwtunnelIp6Id", LwtunnelIp6Id, 1},
			{"LwtunnelIp6Dst", LwtunnelIp6Dst, 2},
			{"LwtunnelIp6Src", LwtunnelIp6Src, 3},
			{"LwtunnelIp6Hoplimit", LwtunnelIp6Hoplimit, 4},
			{"LwtunnelIp6Tc", LwtunnelIp6Tc, 5},
			{"LwtunnelIp6Flags", LwtunnelIp6Flags, 6},
			{"LwtunnelIp6Pad", LwtunnelIp6Pad, 7},
			{"LwtunnelIp6Opts", LwtunnelIp6Opts, 8},
			{"LwtunnelIp6Max", LwtunnelIp6Max, LwtunnelIp6Opts}, // Note that LwtunnelIp6Max is 7
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				if test.value != test.expected {
					t.Errorf("Expected %d but got %d for %s", test.expected, test.value, test.name)
				}
			})
		}
	})
}
