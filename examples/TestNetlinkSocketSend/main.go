package main

import (
	"github.com/sam-caldwell/network/netlink/core"
	"golang.org/x/sys/unix"
	"log"
)

func main() {
	// Create a netlink socket
	fd, err := unix.Socket(unix.AF_NETLINK, unix.SOCK_RAW, unix.NETLINK_GENERIC)
	if err != nil {
		log.Fatalf("Failed to create netlink socket: %v", err)
	}
	defer func() { _ = unix.Close(fd) }()

	// Bind the socket to our PID
	sa := &unix.SockaddrNetlink{
		Family: unix.AF_NETLINK,
		Pid:    uint32(unix.Getpid()),
		Groups: 0,
	}
	if err := unix.Bind(fd, sa); err != nil {
		log.Fatalf("Failed to bind netlink socket: %v", err)
	}

	// Create a NetlinkSocket instance
	nlSocket := core.NewNetlinkSocket(fd, sa)

	// Prepare a NetlinkRequest to send
	request := &core.NetlinkRequest{
		NetlinkMessageHeader: core.NetlinkMessageHeader{
			Len:   uint32(core.NetlinkMessageHeaderSize),
			Type:  unix.NLMSG_NOOP,
			Flags: 0,
			Seq:   1,
			Pid:   uint32(unix.Getpid()),
		},
		Data: []core.NetlinkRequestData{},
	}

	// Call the Send method
	err = nlSocket.Send(request)
	if err != nil {
		log.Fatalf("Send() returned error: %v", err)
	}

	// Since we're sending a NLMSG_NOOP message, we don't expect a response.
	// However, we can call Receive to check if any error messages are returned.

	// Set socket read timeout to avoid blocking indefinitely
	timeout := unix.Timeval{
		Sec:  1,
		Usec: 0,
	}
	if err := unix.SetsockoptTimeval(fd, unix.SOL_SOCKET, unix.SO_RCVTIMEO, &timeout); err != nil {
		log.Fatalf("Failed to set socket timeout: %v", err)
	}

	// Attempt to receive any messages (should be 1)
	messages, _, err := nlSocket.Receive()
	if err == nil && len(messages) > 1 {
		log.Fatalf("Expected 1 messages, but received %d messages", len(messages))
	}
}
