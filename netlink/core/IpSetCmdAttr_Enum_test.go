package core

import (
	"testing"
	"unsafe"
)

// TestIpSetCmdAttrEnumValues tests that the IpSetCmdAttr constants have the correct values.
func TestIpSetCmdAttrEnum(t *testing.T) {
	t.Run("size check", func(t *testing.T) {
		var attr IpSetCmdAttr
		expectedSize := 1 // uint8 should be 1 byte

		if size := unsafe.Sizeof(attr); size != uintptr(expectedSize) {
			t.Errorf("Expected size %d but got %d", expectedSize, size)
		}
	})
	t.Run("value check", func(t *testing.T) {
		tests := []struct {
			name     string
			value    IpSetCmdAttr
			expected IpSetCmdAttr
		}{
			{"IpsetAttrUnspec", IpsetAttrUnspec, 0},
			{"IpsetAttrProtocol", IpsetAttrProtocol, 1},
			{"IpsetAttrSetname", IpsetAttrSetname, 2},
			{"IpsetAttrTypename", IpsetAttrTypename, 3},
			{"IpsetAttrSetname2", IpsetAttrSetname2, 3}, // Same as Typename
			{"IpsetAttrRevision", IpsetAttrRevision, 4},
			{"IpsetAttrFamily", IpsetAttrFamily, 5},
			{"IpsetAttrFlags", IpsetAttrFlags, 6},
			{"IpsetAttrData", IpsetAttrData, 7},
			{"IpsetAttrAdt", IpsetAttrAdt, 8},
			{"IpsetAttrLineno", IpsetAttrLineno, 9},
			{"IpsetAttrProtocolMin", IpsetAttrProtocolMin, 10},
			{"IpsetAttrRevisionMin", IpsetAttrRevisionMin, 10}, // Same as ProtocolMin
			{"IpsetAttrIndex", IpsetAttrIndex, 11},
			{"IpsetAttrCmdMax", IpsetAttrCmdMax, 11}, // iota - 1 results in the same value as IpsetAttrIndex
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
