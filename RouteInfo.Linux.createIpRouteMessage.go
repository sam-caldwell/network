//go:build linux

package network

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"golang.org/x/sys/unix"
	"net"
)

// createIpv6RouteMessage - Define the IPv6 route message
func createIpRouteMessage(useV6 bool, buffer *bytes.Buffer) (err error) {
	const (
		v4Len = net.IPv4len * 8 // Length in bits
		v6Len = net.IPv6len * 8 // Length in bits
	)

	var destinationLength uint8 = v4Len

	if useV6 {
		destinationLength = v6Len
	}

	routeMessage := unix.RtMsg{
		Family:   unix.AF_INET,
		Dst_len:  destinationLength, // Length in bits
		Src_len:  0,
		Tos:      0,
		Table:    unix.RT_TABLE_MAIN,
		Protocol: unix.RTPROT_BOOT,
		Scope:    unix.RT_SCOPE_UNIVERSE,
		Type:     unix.RTN_UNICAST,
	}

	if err := binary.Write(buffer, binary.LittleEndian, routeMessage); err != nil {
		return fmt.Errorf("binary write error: %v", err)
	}

	return nil

}
