package core

import (
	"golang.org/x/sys/unix"
)

// SetExtAck - request error messages to be reported on the socket
func (s *NetlinkSocket) SetExtAck(enable bool) error {

	var n int

	if enable {
		n = 1
	}

	return unix.SetsockoptInt(int(s.fd), unix.SOL_NETLINK, unix.NETLINK_EXT_ACK, n)

}
