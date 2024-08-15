//go:build linux

package network

import (
	"bytes"
	"encoding/binary"
	"golang.org/x/sys/unix"
	"net"
	"testing"
)

func TestRouteInfo_createIpRouteMessage(t *testing.T) {

	var expectedV4, expectedV6 bytes.Buffer
	t.Run("setup", func(t *testing.T) {
		// Expected IPv4 message
		_ = binary.Write(&expectedV4, binary.LittleEndian, unix.RtMsg{
			Family:   unix.AF_INET,
			Dst_len:  net.IPv4len * 8, // Length in bits
			Src_len:  0,
			Tos:      0,
			Table:    unix.RT_TABLE_MAIN,
			Protocol: unix.RTPROT_BOOT,
			Scope:    unix.RT_SCOPE_UNIVERSE,
			Type:     unix.RTN_UNICAST,
		})

		// Expected IPv6 message (corrected)
		_ = binary.Write(&expectedV6, binary.LittleEndian, unix.RtMsg{
			Family:   unix.AF_INET6,
			Dst_len:  net.IPv6len * 8, // Length in bits
			Src_len:  0,
			Tos:      0,
			Table:    unix.RT_TABLE_MAIN,
			Protocol: unix.RTPROT_BOOT,
			Scope:    unix.RT_SCOPE_UNIVERSE,
			Type:     unix.RTN_UNICAST,
		})
	})

	t.Run("test ipv4", func(t *testing.T) {
		var bufferV4 bytes.Buffer

		if err := createIpRouteMessage(false, &bufferV4); err != nil {
			t.Error(err)
		}

		if actual := bufferV4.Bytes(); !bytes.Equal(actual, expectedV4.Bytes()) {
			t.Fatalf("message is invalid.\n"+
				"  actual: %v\n"+
				"expected: %v\n", actual, expectedV4.Bytes())
		}
	})

	t.Run("test ipv6", func(t *testing.T) {
		var bufferV6 bytes.Buffer

		if err := createIpRouteMessage(true, &bufferV6); err != nil {
			t.Error(err)
		}

		if actual := bufferV6.Bytes(); !bytes.Equal(actual, expectedV6.Bytes()) {
			t.Fatalf("message is invalid.\n"+
				"  actual: %v\n"+
				"expected: %v\n", actual, expectedV6.Bytes())
		}
	})
}
