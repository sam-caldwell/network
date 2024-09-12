package core

import (
	"golang.org/x/sys/unix"
	"testing"
)

// TestIfRtaEnum tests the values of IfRtaEnum constants.
func TestIfRtaEnum(t *testing.T) {
	tests := []struct {
		name     string
		value    IfRtaEnum
		expected IfRtaEnum
	}{
		{"IflaUnspec", IflaUnspec, IfRtaEnum(unix.IFLA_UNSPEC)},
		{"IflaAddress", IflaAddress, IfRtaEnum(unix.IFLA_ADDRESS)},
		{"IflaBroadcast", IflaBroadcast, IfRtaEnum(unix.IFLA_BROADCAST)},
		{"IflaIfname", IflaIfname, IfRtaEnum(unix.IFLA_IFNAME)},
		{"IflaMtu", IflaMtu, IfRtaEnum(unix.IFLA_MTU)},
		{"IflaLink", IflaLink, IfRtaEnum(unix.IFLA_LINK)},
		{"IflaQdisc", IflaQdisc, IfRtaEnum(unix.IFLA_QDISC)},
		{"IflaStats", IflaStats, IfRtaEnum(unix.IFLA_STATS)},
		{"IflaPermAddress", IflaPermAddress, IfRtaEnum(unix.IFLA_PERM_ADDRESS)},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.value != test.expected {
				t.Errorf("Expected %d but got %d for %s", test.expected, test.value, test.name)
			}
		})
	}
}
