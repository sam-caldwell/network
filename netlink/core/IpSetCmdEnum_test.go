package core

import (
	"testing"
	"unsafe"
)

func TestIpSetCmdEnum(t *testing.T) {
	t.Run("TestSize", func(t *testing.T) {
		const expectedSizeInBytes = 1
		var cmd IpSetCmdEnum

		if size := unsafe.Sizeof(cmd); size != uintptr(expectedSizeInBytes) {
			t.Errorf("Expected size %d but got %d", expectedSizeInBytes, size)
		}
	})

	t.Run("TestValues", func(t *testing.T) {
		tests := []struct {
			name     string
			value    IpSetCmdEnum
			expected IpSetCmdEnum
		}{
			{"IpsetCmdProtocol", IpsetCmdProtocol, 1},
			{"IpsetCmdCreate", IpsetCmdCreate, 2},
			{"IpsetCmdDestroy", IpsetCmdDestroy, 3},
			{"IpsetCmdFlush", IpsetCmdFlush, 4},
			{"IpsetCmdRename", IpsetCmdRename, 5},
			{"IpsetCmdSwap", IpsetCmdSwap, 6},
			{"IpsetCmdList", IpsetCmdList, 7},
			{"IpsetCmdSave", IpsetCmdSave, 8},
			{"IpsetCmdAdd", IpsetCmdAdd, 9},
			{"IpsetCmdDel", IpsetCmdDel, 10},
			{"IpsetCmdTest", IpsetCmdTest, 11},
			{"IpsetCmdHeader", IpsetCmdHeader, 12},
			{"IpsetCmdType", IpsetCmdType, 13},
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
