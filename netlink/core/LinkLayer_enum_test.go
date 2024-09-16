package core

import (
	"testing"
	"unsafe"
)

func TestLinkLayer(t *testing.T) {
	t.Run("TestSize", func(t *testing.T) {
		const expectedSizeInBytes = 1
		var ll LinkLayer

		if size := unsafe.Sizeof(ll); size != uintptr(expectedSizeInBytes) {
			t.Errorf("Expected size %d but got %d", expectedSizeInBytes, size)
		}
	})

	t.Run("TestValues", func(t *testing.T) {
		tests := []struct {
			name     string
			value    LinkLayer
			expected LinkLayer
		}{
			{"LinkLayerUnspec", LinkLayerUnspec, 0},
			{"LinkLayerEthernet", LinkLayerEthernet, 1},
			{"LinkLayerAtm", LinkLayerAtm, 2},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				if test.value != test.expected {
					t.Errorf("Expected %d but got %d for %s", test.expected, test.value, test.name)
				}
			})
		}
	})

	t.Run("TestTcLinklayerMask", func(t *testing.T) {
		const expectedMask = 0x0F
		if TcLinklayerMask != expectedMask {
			t.Errorf("Expected mask %d but got %d", expectedMask, TcLinklayerMask)
		}
	})
}
