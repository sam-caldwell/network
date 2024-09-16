package core

import (
	"golang.org/x/sys/unix"
)

// SetsockoptInet4Addr - wrapper for unix.SetsockoptInet4Addr
func SetsockoptInet4Addr(fd, level int, opt SocketOption, value [4]byte) (err error) {
	return unix.SetsockoptInet4Addr(fd, level, int(opt), value)
}
