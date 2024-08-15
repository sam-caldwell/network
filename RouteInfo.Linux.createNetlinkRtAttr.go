//go:build linux

package network

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"golang.org/x/sys/unix"
	"net"
)

// createNetlinkRtAttr - create netlink rt Attribute
func createNetlinkRtAttr(buffer *bytes.Buffer, attrType uint16, attrIpLen uint16, ipValue net.IP) error {

	rtAttr := unix.RtAttr{
		Len:  unix.SizeofRtAttr + attrIpLen,
		Type: attrType,
	}

	if err := binary.Write(buffer, binary.LittleEndian, rtAttr); err != nil {
		return fmt.Errorf("binary write error: %v", err)
	}

	buffer.Write(ipValue)

	return nil

}
