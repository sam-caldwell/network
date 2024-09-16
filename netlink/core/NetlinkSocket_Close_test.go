//go:build linux

package core

import (
	"golang.org/x/sys/unix"
	"runtime"
	"sync/atomic"
	"testing"
)

func TestNetlinkSocket_Close(t *testing.T) {
	// Ensure the test runs only on Linux
	if runtime.GOOS != "linux" {
		t.Skip("Netlink sockets are only available on Linux")
	}

	// Create a netlink socket
	fd, err := unix.Socket(unix.AF_NETLINK, unix.SOCK_RAW, unix.NETLINK_GENERIC)
	if err != nil {
		t.Fatalf("Failed to create netlink socket: %v", err)
	}

	// Keep track of whether the fd has been closed
	fdClosed := false

	// Ensure the socket is closed at the end if not already closed
	defer func() {
		if !fdClosed {
			unix.Close(fd)
		}
	}()

	// Initialize NetlinkSocket
	sa := &unix.SockaddrNetlink{
		Family: unix.AF_NETLINK,
		Pid:    uint32(unix.Getpid()),
		Groups: 0,
	}

	nlSocket := NewNetlinkSocket(fd, sa)

	// Call Close
	err = nlSocket.Close()
	if err != nil {
		t.Errorf("Close() returned error: %v", err)
	} else {
		fdClosed = true
	}

	// Verify that fd is set to -1
	if atomic.LoadInt32(&nlSocket.fd) != -1 {
		t.Errorf("Expected fd to be -1 after close, got %d", nlSocket.fd)
	}

	// Attempt to call Close again and expect an error
	err = nlSocket.Close()
	if err == nil {
		t.Errorf("Expected error when closing an already closed socket, but got nil")
	}
}
