//go:build linux

package network

import (
	"bytes"
	"encoding/binary"
	"golang.org/x/sys/unix"
	"net"
	"testing"
)

func TestRouteInfo_createNetlinkRtAttr(t *testing.T) {
	t.Run("test with v4", func(t *testing.T) {
		var actual bytes.Buffer
		var expected bytes.Buffer
		_ = binary.Write(&expected, binary.LittleEndian, unix.RtAttr{
			Len:  unix.SizeofRtAttr + net.IPv4len,
			Type: unix.RTA_DST,
		})
		ip := net.ParseIP("192.168.1.1")
		_ = binary.Write(&expected, binary.LittleEndian, ip)

		_ = createNetlinkRtAttr(&actual, unix.RTA_DST, net.IPv4len, ip)

		if !bytes.Equal(actual.Bytes(), expected.Bytes()) {
			t.Errorf("createNetlinkRtAttr() mismatch\n"+
				"  actual:%v\n"+
				"expected:%v",
				actual, expected)
		}
	})

	t.Run("test with v6", func(t *testing.T) {
		var actual bytes.Buffer
		var expected bytes.Buffer
		ip := net.ParseIP("fe80::1fa8:4cd:8210:5e87/128")
		_ = binary.Write(&expected, binary.LittleEndian, unix.RtAttr{
			Len:  unix.SizeofRtAttr + net.IPv6len,
			Type: unix.RTA_DST,
		})
		_ = binary.Write(&expected, binary.LittleEndian, ip)

		_ = createNetlinkRtAttr(&actual, unix.RTA_DST, net.IPv6len, ip)

		if !bytes.Equal(actual.Bytes(), expected.Bytes()) {
			t.Errorf("createNetlinkRtAttr() mismatch\n"+
				"  actual:%v\n"+
				"expected:%v",
				actual, expected)
		}
	})

}
