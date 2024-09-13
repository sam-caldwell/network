package core

import (
	"golang.org/x/sys/unix"
	"testing"
	"unsafe"
)

// TestInterfaceFamily tests the values of InterfaceFamily constants.
func TestInterface(t *testing.T) {
	t.Run("test type size", func(t *testing.T) {
		const expectedSizeInBytes = 1
		if sz := int(unsafe.Sizeof(InterfaceFamily(0))); sz != expectedSizeInBytes {
			t.Errorf("got %d; want %d", sz, expectedSizeInBytes)
		}
	})
	t.Run("test enumerated values", func(t *testing.T) {
		tests := []struct {
			name     string
			value    InterfaceFamily
			expected InterfaceFamily
		}{
			{"AfUnspec", InterfaceFamily(AfUnspec), InterfaceFamily(unix.AF_UNSPEC)},
			{"AfUnix", InterfaceFamily(AfUnix), InterfaceFamily(unix.AF_UNIX)},
			{"AfInet", InterfaceFamily(AfInet), InterfaceFamily(unix.AF_INET)},
			{"AfAx25", InterfaceFamily(AfAx25), InterfaceFamily(unix.AF_X25)},
			{"AfIpx", InterfaceFamily(AfIpx), InterfaceFamily(unix.AF_IPX)},
			{"AfAppleTalk", InterfaceFamily(AfAppleTalk), InterfaceFamily(unix.AF_APPLETALK)},
			{"AfNetRom", InterfaceFamily(AfNetRom), InterfaceFamily(unix.AF_NETROM)},
			{"AfBridge", InterfaceFamily(AfBridge), InterfaceFamily(unix.AF_BRIDGE)},
			{"AfAal5", InterfaceFamily(AfAal5), InterfaceFamily(8)}, // Not in unix package, manually set to 8
			{"AfX25", InterfaceFamily(AfX25), InterfaceFamily(unix.AF_X25)},
			{"AfInet6", InterfaceFamily(AfInet6), InterfaceFamily(unix.AF_INET6)},
			{"AfMax", InterfaceFamily(AfMax), InterfaceFamily(unix.AF_MAX)},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				if test.value != test.expected {
					t.Errorf("Expected %d but got %d for %s", test.expected, test.value, test.name)
				}
			})
		}
	})
	t.Run("test Equal method", func(t *testing.T) {
		tests := []struct {
			name     string
			lhs      InterfaceFamily
			rhs      InterfaceFamily
			expected bool
		}{
			{"Equal values (AfUnspec)", AfUnspec, AfUnspec, true},
			{"Equal values (AfInet)", AfInet, AfInet, true},
			{"Different values (AfInet vs AfInet6)", AfInet, AfInet6, false},
			{"Different values (AfAppleTalk vs AfNetRom)", AfAppleTalk, AfNetRom, false},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				if got := test.lhs.Equal(test.rhs); got != test.expected {
					t.Errorf("Expected %v but got %v for %s", test.expected, got, test.name)
				}
			})
		}
	})
	t.Run("test ToUint8() method", func(t *testing.T) {
		tests := []struct {
			name     string
			input    InterfaceFamily
			expected uint8
		}{
			{"AfUnspec", AfUnspec, uint8(AfUnspec)},
			{"AfInet", AfInet, uint8(AfInet)},
			{"AfAppleTalk", AfAppleTalk, uint8(AfAppleTalk)},
			{"AfInet6", AfInet6, uint8(AfInet6)},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				if got := test.input.ToUint8(); got != test.expected {
					t.Errorf("Expected %d but got %d for %s", test.expected, got, test.name)
				}
			})
		}
	})
}
