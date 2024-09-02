//go:build linux

package core

import (
	"golang.org/x/sys/unix"
)

// SetReceiveBufferSize - set receive buffer size on the socket
func (s *NetlinkSocket) SetReceiveBufferSize(size int, force bool) error {

	opt := unix.SO_RCVBUF

	if force {
		opt = unix.SO_RCVBUFFORCE
	}

	return unix.SetsockoptInt(int(s.fd), unix.SOL_SOCKET, opt, size)

}
