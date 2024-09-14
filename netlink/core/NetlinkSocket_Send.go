package core

import (
	"fmt"
	"golang.org/x/sys/unix"
	"sync/atomic"
)

// Send - send a given NetlinkRequest via the current socket file descriptor to netlink
func (s *NetlinkSocket) Send(request *NetlinkRequest) (err error) {

	if request == nil {
		return fmt.Errorf("nil request")
	}

	socketFd := int(atomic.LoadInt32(&s.fd))

	if socketFd < 0 {

		return fmt.Errorf("send called on a closed socket")

	}

	var requestBytes []byte
	if requestBytes, err = request.Serialize(); err != nil {
		return err
	}

	if err = unix.Sendto(socketFd, requestBytes, 0, &s.lsa); err != nil {
		return err
	}

	return err

}
