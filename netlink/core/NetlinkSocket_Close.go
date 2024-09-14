package core

import (
	"golang.org/x/sys/unix"
	"sync/atomic"
)

// Close - close the NetlinkSocket by safely setting its file descriptor (fd) to -1 and then closing the original fd.
//
// Thread Safety:
//
//	The method uses an atomic operation (atomic.SwapInt32) to ensure that the file descriptor is safely set to -1
//	before closing the original file descriptor. This prevents race conditions where multiple goroutines might attempt
//	to close the socket simultaneously. The atomic swap ensures that only the first call to Close will actually close
//	the file descriptor, while subsequent calls will safely do nothing, as the fd will already be set to -1.
//
// Returns:
//
//	An error, if any, encountered while closing the socket.
func (s *NetlinkSocket) Close() error {
	fd := int(atomic.SwapInt32(&s.fd, -1))
	return unix.Close(fd)
}
