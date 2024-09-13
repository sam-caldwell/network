package core

import (
	"golang.org/x/sys/unix"
	"testing"
	"unsafe"
)

func TestIfAddressMessage(t *testing.T) {

	t.Run("test IfAddressMessage struct", func(t *testing.T) {
		t.Run("ipv4 should be supported", func(t *testing.T) {
			t.Run("field check", func(t *testing.T) {
				i := IfAddressMessage{}
				i.Family = AfInet
				i.Prefixlen = uint8(0)
				i.Flags = uint8(0)
				i.Scope = uint8(0)
				i.Index = uint32(0)
			})
			t.Run("size check", func(t *testing.T) {
				const expectedSizeInBytes = IfAddressMessageSize
				if sz := int(unsafe.Sizeof(IfAddressMessage{})); sz != expectedSizeInBytes {
					t.Fatalf("size mismatch.  sz: %d", sz)
				}
			})
		})

		t.Run("ipv6 should be supported", func(t *testing.T) {
			t.Run("field check", func(t *testing.T) {
				i := IfAddressMessage{}
				i.Family = AfInet6
				i.Prefixlen = uint8(0)
				i.Flags = uint8(0)
				i.Scope = uint8(0)
				i.Index = uint32(0)
			})
			t.Run("size check", func(t *testing.T) {
				const expectedSizeInBytes = IfAddressMessageSize
				if sz := int(unsafe.Sizeof(IfAddressMessage{})); sz != expectedSizeInBytes {
					t.Fatalf("size mismatch.  sz: %d", sz)
				}
			})
		})
		t.Run("Test IfAddressMessage creation", func(t *testing.T) {
			// Define test cases with different InterfaceFamily values.
			testCases := []struct {
				name   string
				family InterfaceFamily
			}{
				{"IPv4", InterfaceFamily(unix.AF_INET)},
				{"IPv6", InterfaceFamily(unix.AF_INET6)},
				{"Unspecified", InterfaceFamily(unix.AF_UNSPEC)},
				// Add more interface families if needed.
			}

			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					msg := NewIfAddressMessage(tc.family)

					if msg == nil {
						t.Fatalf("Expected non-nil IfAddressMessage, got nil")
					}

					// Verify that the Family field is set correctly.
					if msg.Family != uint8(tc.family) {
						t.Errorf("Expected Family: %d, got: %d", tc.family, msg.Family)
					}

					// Verify that other fields are initialized to zero.
					if msg.Prefixlen != 0 {
						t.Errorf("Expected Prefixlen: 0, got: %d", msg.Prefixlen)
					}
					if msg.Flags != 0 {
						t.Errorf("Expected Flags: 0, got: %d", msg.Flags)
					}
					if msg.Scope != 0 {
						t.Errorf("Expected Scope: 0, got: %d", msg.Scope)
					}
					if msg.Index != 0 {
						t.Errorf("Expected Index: 0, got: %d", msg.Index)
					}
				})
			}
		})
	})
}
