package netlink

import (
	"golang.org/x/sys/unix"
)

// SetSocketReceiveBufferSize - for each socket in the netlink handle, set receive buffer size.
// The maximum value is capped by /proc/sys/net/core/rmem_max.
func (h *Handle) SetSocketReceiveBufferSize(size int, force bool) error {
	opt := unix.SO_RCVBUF
	if force {
		opt = unix.SO_RCVBUFFORCE
	}
	for _, sh := range h.sockets {
		fd := sh.Socket.GetFd()
		err := unix.SetsockoptInt(fd, unix.SOL_SOCKET, opt, size)
		if err != nil {
			return err
		}
	}
	return nil
}
