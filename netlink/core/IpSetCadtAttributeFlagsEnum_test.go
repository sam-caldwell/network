package core

import (
	"testing"
	"unsafe"
)

func TestIpSetCadtAttributeFlagsEnum(t *testing.T) {
	t.Run("size check", func(t *testing.T) {
		const expectedSizeInBytes = 1
		if sz := int(unsafe.Sizeof(IpSetCadtAttributeFlagsEnum(0))); sz != expectedSizeInBytes {
			t.Fatal("size mismatch")
		}
	})
	t.Run("value check", func(t *testing.T) {
		tests := []struct {
			name     string
			value    IpSetCadtAttributeFlagsEnum
			expected IpSetCadtAttributeFlagsEnum
		}{
			{"IpsetFlagBitBefore", IpsetFlagBitBefore, 0},
			{"IpsetFlagBefore", IpSetCadtAttributeFlagsEnum(IpsetFlagBefore), 1},
			{"IpsetFlagBitPhysdev", IpsetFlagBitPhysdev, 1},
			{"IpsetFlagPhysdev", IpSetCadtAttributeFlagsEnum(IpsetFlagPhysdev), 1 << 1},
			{"IpsetFlagBitNomatch", IpsetFlagBitNomatch, 2},
			{"IpsetFlagNomatch", IpSetCadtAttributeFlagsEnum(IpsetFlagNomatch), 1 << 2},
			{"IpsetFlagBitWithCounters", IpsetFlagBitWithCounters, 3},
			{"IpsetFlagWithCounters", IpSetCadtAttributeFlagsEnum(IpsetFlagWithCounters), 1 << 3},
			{"IpsetFlagBitWithComment", IpsetFlagBitWithComment, 4},
			{"IpsetFlagWithComment", IpSetCadtAttributeFlagsEnum(IpsetFlagWithComment), 1 << 4},
			{"IpsetFlagBitWithForceadd", IpsetFlagBitWithForceadd, 5},
			{"IpsetFlagWithForceadd", IpSetCadtAttributeFlagsEnum(IpsetFlagWithForceadd), 1 << 5},
			{"IpsetFlagBitWithSkbinfo", IpsetFlagBitWithSkbinfo, 6},
			{"IpsetFlagWithSkbinfo", IpSetCadtAttributeFlagsEnum(IpsetFlagWithSkbinfo), 1 << 6},
			{"IpsetFlagCadtMax", IpsetFlagCadtMax, 15},
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
