//go:build linux

package network

import (
	"golang.org/x/sys/unix"
	"testing"
	"unsafe"
)

func TestHandleAckResponse(t *testing.T) {
	// Create a netlink socket
	fd, err := unix.Socket(unix.AF_NETLINK, unix.SOCK_RAW, unix.NETLINK_ROUTE)
	if err != nil {
		t.Fatalf("failed to create socket: %v", err)
	}
	defer func() { _ = unix.Close(fd) }()

	// Bind the socket to a local address
	localAddr := &unix.SockaddrNetlink{
		Family: unix.AF_NETLINK,
		Pid:    uint32(unix.Getpid()),
	}

	if err := unix.Bind(fd, localAddr); err != nil {
		t.Fatalf("failed to bind socket: %v", err)
	}

	// Create a netlink message header for a dummy request
	nlmsg := unix.NlMsghdr{
		Len:   unix.SizeofNlMsghdr,
		Type:  unix.RTM_GETROUTE, // Dummy request type, could be anything
		Flags: unix.NLM_F_REQUEST | unix.NLM_F_ACK,
		Seq:   1,
		Pid:   uint32(unix.Getpid()),
	}

	// Send the message to the kernel
	if err := unix.Sendto(fd, (*(*[unix.SizeofNlMsghdr]byte)(unsafe.Pointer(&nlmsg)))[:], 0, &unix.SockaddrNetlink{Family: unix.AF_NETLINK}); err != nil {
		t.Fatalf("failed to send message: %v", err)
	}

	// Now call handleAckResponse to receive and handle the ACK
	if err := handleAckResponse(fd); err != nil {
		t.Fatalf("handleAckResponse() failed: '%v'", err)
	}
}
