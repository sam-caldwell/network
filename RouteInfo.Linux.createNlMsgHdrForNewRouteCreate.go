//go:build linux

package network

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"golang.org/x/sys/unix"
)

// createNlMsgHdrForNewRouteCreate - Netlink message header
func createNlMsgHdrForNewRouteCreate(buffer *bytes.Buffer) (err error) {

	hdr := unix.NlMsghdr{
		Len:   uint32(unix.SizeofNlMsghdr + unix.SizeofRtMsg),
		Type:  unix.RTM_NEWROUTE,
		Flags: unix.NLM_F_CREATE | unix.NLM_F_REQUEST | unix.NLM_F_ACK,
		Seq:   1,
		Pid:   uint32(unix.Getpid()),
	}

	if err := binary.Write(buffer, binary.LittleEndian, hdr); err != nil {
		return fmt.Errorf("binary write error: %v", err)
	}

	return nil

}
