//go:build linux

package core

import "sync/atomic"

// GetFd - Return the NetlinkSocket file descriptor
func (s *NetlinkSocket) GetFd() int {

	return int(atomic.LoadInt32(&s.fd))

}
