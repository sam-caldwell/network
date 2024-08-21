package core

import (
	"testing"
)

func TestIfAddressMessage_struct(t *testing.T) {

	t.Run("ipv4 should be supported", func(t *testing.T) {
		i := &IfAddressMessage{}
		i.Family = AfInet
		i.Prefixlen = uint8(0)
		i.Flags = uint8(0)
		i.Scope = uint8(0)
		i.Index = uint32(0)
	})

	t.Run("ipv6 should be supported", func(t *testing.T) {
		i := &IfAddressMessage{}
		i.Family = AfInet6
		i.Prefixlen = uint8(0)
		i.Flags = uint8(0)
		i.Scope = uint8(0)
		i.Index = uint32(0)
	})
}
