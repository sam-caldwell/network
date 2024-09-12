package core

import (
	"testing"
	"unsafe"
)

func TestLwTunnelEncapsulation(t *testing.T) {
	t.Run("TestSize", func(t *testing.T) {
		const expectedSizeInBytes = 1
		var encapType LwTunnelEncapsulation

		if size := unsafe.Sizeof(encapType); size != uintptr(expectedSizeInBytes) {
			t.Errorf("Expected size %d but got %d", expectedSizeInBytes, size)
		}
	})

	t.Run("TestValues", func(t *testing.T) {
		tests := []struct {
			name     string
			value    LwTunnelEncapsulation
			expected LwTunnelEncapsulation
		}{
			{"LwtunnelEncapNone", LwtunnelEncapNone, 0},
			{"LwtunnelEncapMpls", LwtunnelEncapMpls, 1},
			{"LwtunnelEncapIp", LwtunnelEncapIp, 2},
			{"LwtunnelEncapIla", LwtunnelEncapIla, 3},
			{"LwtunnelEncapIp6", LwtunnelEncapIp6, 4},
			{"LwtunnelEncapSeg6", LwtunnelEncapSeg6, 5},
			{"LwtunnelEncapBpf", LwtunnelEncapBpf, 6},
			{"LwtunnelEncapSeg6Local", LwtunnelEncapSeg6Local, 7},
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
