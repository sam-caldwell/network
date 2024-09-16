package core

import (
	"testing"
	"unsafe"
)

func TestIpSetCadtAttributeFlagsEnum(t *testing.T) {
	t.Run("size check", func(t *testing.T) {
		const expectedSizeInBytes = 1
		if sz := int(unsafe.Sizeof(IpSetCadtAttributeFlags(0))); sz != expectedSizeInBytes {
			t.Fatal("size mismatch")
		}
	})
	t.Run("value check", func(t *testing.T) {
		tests := []struct {
			name     string
			value    IpSetCadtAttributeFlags
			expected IpSetCadtAttributeFlags
		}{
			{"IpsetFlagBitBefore", IpsetFlagBitBefore, 0},
			{"IpsetFlagBefore", IpSetCadtAttributeFlags(IpsetFlagBefore), 1},
			{"IpsetFlagBitPhysdev", IpsetFlagBitPhysdev, 1},
			{"IpsetFlagPhysdev", IpSetCadtAttributeFlags(IpsetFlagPhysdev), 1 << 1},
			{"IpsetFlagBitNomatch", IpsetFlagBitNomatch, 2},
			{"IpsetFlagNomatch", IpSetCadtAttributeFlags(IpsetFlagNomatch), 1 << 2},
			{"IpsetFlagBitWithCounters", IpsetFlagBitWithCounters, 3},
			{"IpsetFlagWithCounters", IpSetCadtAttributeFlags(IpsetFlagWithCounters), 1 << 3},
			{"IpsetFlagBitWithComment", IpsetFlagBitWithComment, 4},
			{"IpsetFlagWithComment", IpSetCadtAttributeFlags(IpsetFlagWithComment), 1 << 4},
			{"IpsetFlagBitWithForceadd", IpsetFlagBitWithForceadd, 5},
			{"IpsetFlagWithForceadd", IpSetCadtAttributeFlags(IpsetFlagWithForceadd), 1 << 5},
			{"IpsetFlagBitWithSkbinfo", IpsetFlagBitWithSkbinfo, 6},
			{"IpsetFlagWithSkbinfo", IpSetCadtAttributeFlags(IpsetFlagWithSkbinfo), 1 << 6},
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
