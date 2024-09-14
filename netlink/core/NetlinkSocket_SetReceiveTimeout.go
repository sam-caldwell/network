package core

import (
	"golang.org/x/sys/unix"
)

// SetReceiveTimeout - set receive timeout on the socket
//
// Set a read timeout of SOCKET_READ_TIMEOUT, this will allow the Read to periodically unblock and avoid that a routine
// remains stuck on a recvmsg on a closed fd
func (s *NetlinkSocket) SetReceiveTimeout(timeout *unix.Timeval) error {

	return unix.SetsockoptTimeval(int(s.fd), unix.SOL_SOCKET, unix.SO_RCVTIMEO, timeout)

}
