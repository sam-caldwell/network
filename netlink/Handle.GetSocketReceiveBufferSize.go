package netlink

import (
	"golang.org/x/sys/unix"
)

// GetSocketReceiveBufferSize - For each socket in the netlink handle, get the receiver buffer size. The retrieved
// value should be the double to the one set for SetSocketReceiveBufferSize.
func (h *Handle) GetSocketReceiveBufferSize() ([]int, error) {

	results := make([]int, len(h.sockets))

	i := 0

	for _, sh := range h.sockets {

		fd := sh.Socket.GetFd()

		size, err := unix.GetsockoptInt(fd, unix.SOL_SOCKET, unix.SO_RCVBUF)

		if err != nil {
			return nil, err
		}

		results[i] = size

		i++

	}

	return results, nil

}
