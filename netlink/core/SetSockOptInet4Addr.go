package core

import (
	"golang.org/x/sys/unix"
)

// SetSockOptInet4Addr - wrapper for unix.SetsockoptInet4Addr
func SetSockOptInet4Addr(fd, level int, opt SocketOption, value [4]byte) (err error) {
	return unix.SetsockoptInet4Addr(fd, level, int(opt), value)
}
