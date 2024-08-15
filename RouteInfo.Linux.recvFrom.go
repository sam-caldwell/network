//go:build linux

package network

import (
	"fmt"
	"golang.org/x/sys/unix"
	"unsafe"
)

// handleAckResponse - Handle ACK response (optional but recommended)
func handleAckResponse(socketFileDescriptor int) (err error) {

	var nlmsg *unix.NlMsghdr

	ack := make([]byte, 4096)

	if n, _, err := unix.Recvfrom(socketFileDescriptor, ack, 0); err != nil || n < unix.SizeofNlMsghdr {
		return fmt.Errorf("recvfrom error: %v", err)
	}

	if nlmsg = (*unix.NlMsghdr)(unsafe.Pointer(&ack[0])); nlmsg.Type == unix.NLMSG_ERROR {
		return fmt.Errorf("netlink error: %v", nlmsg)
	}

	return nil

}
