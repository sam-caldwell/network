package core

import (
	"golang.org/x/sys/unix"
)

// MockNetlinkSocket for testing purposes
type MockNetlinkSocket struct {
	fd int
}

// SetExtAck - mock implementation of the SetExtAck method for testing
func (s *MockNetlinkSocket) SetExtAck(enable bool) error {
	var n int
	if enable {
		n = 1
	}
	return unix.SetsockoptInt(s.fd, unix.SOL_NETLINK, unix.NETLINK_EXT_ACK, n)
}
