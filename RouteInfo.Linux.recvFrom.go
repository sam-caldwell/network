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

	nlmsg = (*unix.NlMsghdr)(unsafe.Pointer(&ack[0]))
	if nlmsg.Type == unix.NLMSG_ERROR {
		// Extract the netlink error message
		nlmsgErr := (*unix.NlMsgerr)(unsafe.Pointer(&ack[unix.SizeofNlMsghdr]))
		if nlmsgErr.Error == 0 {
			// This is actually an acknowledgment of success
			return nil
		}
		return fmt.Errorf("netlink error: %v (error code: %d)", nlmsg, nlmsgErr.Error)
	}

	return nil
}
