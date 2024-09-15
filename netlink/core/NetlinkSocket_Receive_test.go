//go:build linux

package core

import (
	"encoding/binary"
	"golang.org/x/sys/unix"
	"runtime"
	"testing"
)

func TestNetlinkSocket_Receive(t *testing.T) {
	t.Skip("ToDo: Move test to docker")
	// Ensure the test runs only on Linux
	if runtime.GOOS != "linux" {
		t.Skip("Netlink sockets are only available on Linux")
	}

	// Create a netlink socket
	fd, err := unix.Socket(unix.AF_NETLINK, unix.SOCK_RAW, unix.NETLINK_GENERIC)
	if err != nil {
		t.Fatalf("Failed to create netlink socket: %v", err)
	}
	defer func() { _ = unix.Close(fd) }()

	// Bind the socket to our PID
	sa := &unix.SockaddrNetlink{
		Family: unix.AF_NETLINK,
		Pid:    uint32(unix.Getpid()),
	}
	if err := unix.Bind(fd, sa); err != nil {
		t.Fatalf("Failed to bind netlink socket: %v", err)
	}

	// Prepare a NetlinkMessage to send
	msg := NetlinkMessage{
		Header: NetlinkMessageHeader{
			Len:   uint32(NetlinkMessageHeaderSize),
			Type:  unix.NLMSG_NOOP,
			Flags: 0,
			Seq:   1,
			Pid:   uint32(unix.Getpid()),
		},
		Data: []byte{},
	}

	// Serialize the message
	msgBytes := make([]byte, msg.Header.Len)
	binary.LittleEndian.PutUint32(msgBytes[0:4], msg.Header.Len)
	binary.LittleEndian.PutUint16(msgBytes[4:6], msg.Header.Type)
	binary.LittleEndian.PutUint16(msgBytes[6:8], msg.Header.Flags)
	binary.LittleEndian.PutUint32(msgBytes[8:12], msg.Header.Seq)
	binary.LittleEndian.PutUint32(msgBytes[12:16], msg.Header.Pid)
	// No data payload

	// Send the message to ourselves
	dest := &unix.SockaddrNetlink{
		Family: unix.AF_NETLINK,
		Pid:    uint32(unix.Getpid()),
	}
	if err := unix.Sendto(fd, msgBytes, 0, dest); err != nil {
		t.Fatalf("Failed to send netlink message: %v", err)
	}

	// Create a NetlinkSocket instance
	nlSocket := &NetlinkSocket{
		fd:  int32(fd),
		lsa: *sa,
	}

	// Call Receive
	messages, fromAddr, err := nlSocket.Receive()
	if err != nil {
		t.Fatalf("Receive() returned error: %v", err)
	}

	// Verify the received message
	if len(messages) != 1 {
		t.Fatalf("Expected 1 message, got %d", len(messages))
	}

	receivedMsg := messages[0]

	// Check the header
	if receivedMsg.Header.Type != unix.NLMSG_NOOP {
		t.Errorf("Expected message type %d, got %d", unix.NLMSG_NOOP, receivedMsg.Header.Type)
	}
	if receivedMsg.Header.Seq != 1 {
		t.Errorf("Expected sequence number 1, got %d", receivedMsg.Header.Seq)
	}
	if receivedMsg.Header.Pid != uint32(unix.Getpid()) {
		t.Errorf("Expected PID %d, got %d", unix.Getpid(), receivedMsg.Header.Pid)
	}

	// Check that the from address is correct
	if fromAddr == nil {
		t.Fatal("fromAddr is nil")
	}
	if fromAddr.Pid != uint32(unix.Getpid()) {
		t.Errorf("Expected fromAddr.Pid %d, got %d", unix.Getpid(), fromAddr.Pid)
	}
	if fromAddr.Family != unix.AF_NETLINK {
		t.Errorf("Expected fromAddr.Family %d, got %d", unix.AF_NETLINK, fromAddr.Family)
	}
}
