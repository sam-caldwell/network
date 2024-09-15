package core

import (
	"testing"
	"unsafe"
)

func TestMacVlanMacAddrMode(t *testing.T) {
	t.Run("TestSize", func(t *testing.T) {
		const expectedSizeInBytes = 1
		var mode MacVlanMacAddrMode

		if size := unsafe.Sizeof(mode); size != uintptr(expectedSizeInBytes) {
			t.Errorf("Expected size %d but got %d", expectedSizeInBytes, size)
		}
	})

	t.Run("TestValues", func(t *testing.T) {
		tests := []struct {
			name     string
			value    MacVlanMacAddrMode
			expected MacVlanMacAddrMode
		}{
			{"MacvlanMacaddrAdd", MacvlanMacaddrAdd, 0},
			{"MacvlanMacaddrDel", MacvlanMacaddrDel, 1},
			{"MacvlanMacaddrFlush", MacvlanMacaddrFlush, 2},
			{"MacvlanMacaddrSet", MacvlanMacaddrSet, 3},
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
