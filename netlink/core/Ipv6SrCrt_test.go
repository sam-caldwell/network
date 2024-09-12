package core

import (
	"testing"
	"unsafe"
)

func TestIpv6SrCrtConstants(t *testing.T) {
	t.Run("TestSize", func(t *testing.T) {
		const expectedSizeInBytes = 1
		var srcrt Ipv6SrCrt

		if size := unsafe.Sizeof(srcrt); size != uintptr(expectedSizeInBytes) {
			t.Errorf("Expected size %d but got %d", expectedSizeInBytes, size)
		}
	})

	t.Run("TestValues", func(t *testing.T) {
		tests := []struct {
			name     string
			value    Ipv6SrCrt
			expected Ipv6SrCrt
		}{
			{"Ipv6SrcrtStrict", Ipv6SrcrtStrict, 0x01},
			{"Ipv6SrcrtType0", Ipv6SrcrtType0, 0},
			{"Ipv6SrcrtType2", Ipv6SrcrtType2, 2},
			{"Ipv6SrcrtType4", Ipv6SrcrtType4, 4},
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
