package core

import (
	"testing"
	"unsafe"
)

func TestIpSetCadtSpecificAttributes(t *testing.T) {
	t.Run("TestValues", func(t *testing.T) {
		tests := []struct {
			name     string
			value    IpSetCadtSpecificAttributes
			expected IpSetCadtSpecificAttributes
		}{
			{"IpsetAttrIp", IpsetAttrIp, 1},
			{"IpsetAttrIpFrom", IpsetAttrIpFrom, 1},
			{"IpsetAttrIpTo", IpsetAttrIpTo, 2},
			{"IpsetAttrCidr", IpsetAttrCidr, 3},
			{"IpsetAttrPort", IpsetAttrPort, 4},
			{"IpsetAttrPortFrom", IpsetAttrPortFrom, 4},
			{"IpsetAttrPortTo", IpsetAttrPortTo, 5},
			{"IpsetAttrTimeout", IpsetAttrTimeout, 6},
			{"IpsetAttrProto", IpsetAttrProto, 7},
			{"IpsetAttrCadtFlags", IpsetAttrCadtFlags, 8},
			{"IpsetAttrCadtLineno", IpsetAttrCadtLineno, 9},
			{"IpsetAttrMark", IpsetAttrMark, 10},
			{"IpsetAttrMarkmask", IpsetAttrMarkmask, 11},
			{"IpsetAttrBitmask", IpsetAttrBitmask, 12},
			{"IpsetAttrCadtMax", IpsetAttrCadtMax, 16}, // iota calculation
			{"IpsetAttrInitval", IpsetAttrInitval, 17},
			{"IpsetAttrHashsize", IpsetAttrHashsize, 18},
			{"IpsetAttrMaxelem", IpsetAttrMaxelem, 19},
			{"IpsetAttrNetmask", IpsetAttrNetmask, 20},
			{"IpsetAttrBucketsize", IpsetAttrBucketsize, 21},
			{"IpsetAttrResize", IpsetAttrResize, 22},
			{"IpsetAttrSize", IpsetAttrSize, 23},
			{"IpsetAttrElements", IpsetAttrElements, 24},
			{"IpsetAttrReferences", IpsetAttrReferences, 25},
			{"IpsetAttrMemsize", IpsetAttrMemsize, 26},
			{"IpsetAttrCreateMax", IpsetAttrCreateMax, 26},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				if test.value != test.expected {
					t.Errorf("Expected %d but got %d for %s", test.expected, test.value, test.name)
				}
			})
		}
	})

	t.Run("TestSize", func(t *testing.T) {
		const expectedSizeInBytes = unsafe.Sizeof(int(0))
		var attr IpSetCadtSpecificAttributes
		if size := unsafe.Sizeof(attr); size != expectedSizeInBytes {
			t.Errorf("Expected size %d but got %d", expectedSizeInBytes, size)
		}
	})
}
