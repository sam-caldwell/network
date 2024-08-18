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
		return fmt.Errorf("recvfrom error (%d): %v", n, err)
	}

	if nlmsg = (*unix.NlMsghdr)(unsafe.Pointer(&ack[0])); nlmsg.Type == unix.NLMSG_ERROR {
		// Extract the netlink error message
		nlmsgErr := (*unix.NlMsgerr)(unsafe.Pointer(&ack[unix.SizeofNlMsghdr]))
		if nlmsgErr.Error == 0 {
			return nil // This is actually an acknowledgment of success
		}
		return fmt.Errorf("\nfinal state:\n"+
			"netlink error:{\n"+
			"   len  : '%d',\n"+
			"   type : '%d',\n"+
			"   flags: '%d',\n"+
			"   seq  : '%d',\n"+
			"   pid' : '%d',\n"+
			"   code : '%d'\n"+
			"}\n",
			nlmsg.Len, nlmsg.Type, nlmsg.Flags, nlmsg.Seq, nlmsg.Pid, nlmsgErr.Error)
	}

	return nil
}
