package main

import (
	"encoding/binary"
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
	}
	if err := unix.Bind(fd, sa); err != nil {
		log.Fatalf("Failed to bind netlink socket: %v", err)
	}

	// Prepare a NetlinkMessage to send
	msg := core.NetlinkMessage{
		Header: core.NetlinkMessageHeader{
			Len:   uint32(core.NetlinkMessageHeaderSize),
			Type:  unix.NLMSG_NOOP,
			Flags: 0,
			Seq:   1,
			Pid:   uint32(unix.Getpid()),
		},
		Data: []byte{},
	}

	// Serialize the message
	msgBytes := make([]byte, msg.Header.Len)
	binary.NativeEndian.PutUint32(msgBytes[0:4], msg.Header.Len)
	binary.NativeEndian.PutUint16(msgBytes[4:6], msg.Header.Type)
	binary.NativeEndian.PutUint16(msgBytes[6:8], msg.Header.Flags)
	binary.NativeEndian.PutUint32(msgBytes[8:12], msg.Header.Seq)
	binary.NativeEndian.PutUint32(msgBytes[12:16], msg.Header.Pid)
	// No data payload

	// Send the message to ourselves
	dest := &unix.SockaddrNetlink{
		Family: unix.AF_NETLINK,
		Pid:    uint32(unix.Getpid()),
	}
	if err := unix.Sendto(fd, msgBytes, 0, dest); err != nil {
		log.Fatalf("Failed to send netlink message: %v", err)
	}

	// Create a NetlinkSocket instance
	nlSocket := core.NewNetlinkSocket(fd, sa)

	// Call Receive
	messages, fromAddr, err := nlSocket.Receive()
	if err != nil {
		log.Fatalf("Receive() returned error: %v", err)
	}

	// Verify the received message
	if len(messages) != 1 {
		log.Fatalf("Expected 1 message, got %d", len(messages))
	}

	receivedMsg := messages[0]

	// Check the header
	if receivedMsg.Header.Type != unix.NLMSG_NOOP {
		log.Fatalf("Expected message type %d, got %d", unix.NLMSG_NOOP, receivedMsg.Header.Type)
	}
	if receivedMsg.Header.Seq != 1 {
		log.Fatalf("Expected sequence number 1, got %d", receivedMsg.Header.Seq)
	}
	if receivedMsg.Header.Pid != uint32(unix.Getpid()) {
		log.Fatalf("Expected PID %d, got %d", unix.Getpid(), receivedMsg.Header.Pid)
	}

	// Check that the from address is correct
	if fromAddr == nil {
		log.Fatal("fromAddr is nil")
	}
	if fromAddr.Pid != uint32(unix.Getpid()) {
		log.Fatalf("Expected fromAddr.Pid %d, got %d", unix.Getpid(), fromAddr.Pid)
	}
	if fromAddr.Family != unix.AF_NETLINK {
		log.Fatalf("Expected fromAddr.Family %d, got %d", unix.AF_NETLINK, fromAddr.Family)
	}
}
