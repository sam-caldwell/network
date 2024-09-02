//go:build linux

package core

import (
	"golang.org/x/sys/unix"
)

// SetSendTimeout - set send timeout on the socket
//
// Set send timeout of SOCKET_SEND_TIMEOUT, this will allow the Send to periodically unblock and avoid some routine
// from hanging on send on a closed fd.
func (s *NetlinkSocket) SetSendTimeout(timeout *unix.Timeval) error {

	return unix.SetsockoptTimeval(int(s.fd), unix.SOL_SOCKET, unix.SO_SNDTIMEO, timeout)

}
