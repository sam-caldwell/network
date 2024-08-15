//go:build linux

package network

import (
	"bytes"
	"encoding/binary"
	"golang.org/x/sys/unix"
	"testing"
)

func TestRouteInfo_sendMessage(t *testing.T) {
	// Create a netlink socket
	fd, err := unix.Socket(unix.AF_NETLINK, unix.SOCK_RAW, unix.NETLINK_ROUTE)
	if err != nil {
		t.Fatalf("failed to create socket: %v", err)
	}
	defer unix.Close(fd)

	// Bind the socket to a local address
	localAddr := &unix.SockaddrNetlink{
		Family: unix.AF_NETLINK,
		Pid:    uint32(unix.Getpid()),
	}

	if err := unix.Bind(fd, localAddr); err != nil {
		t.Fatalf("failed to bind socket: %v", err)
	}

	// Prepare a message in the buffer
	var buffer bytes.Buffer

	// Add a dummy netlink message header
	nlmsg := unix.NlMsghdr{
		Len:   unix.SizeofNlMsghdr,
		Type:  unix.RTM_NEWROUTE, // Dummy message type
		Flags: unix.NLM_F_REQUEST,
		Seq:   1,
		Pid:   uint32(unix.Getpid()),
	}

	if err := binary.Write(&buffer, binary.LittleEndian, nlmsg); err != nil {
		t.Fatalf("failed to write message header: %v", err)
	}

	// Call sendMessage and verify it does not return an error
	if err := sendMessage(fd, &buffer); err != nil {
		t.Fatalf("sendMessage() failed: %v", err)
	}
}
