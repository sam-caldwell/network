package core

import (
	"golang.org/x/sys/unix"
	"testing"
)

func TestNewNetlinkSocket(t *testing.T) {
	// Sample file descriptor and SockaddrNetlink
	fd := 10 // Sample file descriptor (could be any integer)
	sa := &unix.SockaddrNetlink{
		Family: unix.AF_NETLINK,
		Pid:    uint32(1234),
		Groups: 0,
	}

	// Call the function under test
	nlSocket := NewNetlinkSocket(fd, sa)

	// Verify that the returned NetlinkSocket is not nil
	if nlSocket == nil {
		t.Fatal("Expected NewNetlinkSocket to return a non-nil NetlinkSocket")
	}

	// Verify that the file descriptor is set correctly
	if nlSocket.fd != int32(fd) {
		t.Errorf("Expected fd to be %d, got %d", fd, nlSocket.fd)
	}

	// Verify that the SockaddrNetlink (lsa) is set correctly
	if nlSocket.lsa.Family != sa.Family {
		t.Errorf("Expected lsa.Family to be %d, got %d", sa.Family, nlSocket.lsa.Family)
	}
	if nlSocket.lsa.Pid != sa.Pid {
		t.Errorf("Expected lsa.Pid to be %d, got %d", sa.Pid, nlSocket.lsa.Pid)
	}
	if nlSocket.lsa.Groups != sa.Groups {
		t.Errorf("Expected lsa.Groups to be %d, got %d", sa.Groups, nlSocket.lsa.Groups)
	}
}
