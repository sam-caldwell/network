package core

import (
	"golang.org/x/sys/unix"
)

// SetsockoptInt - wrapper for unix.SetsockoptInt
func SetsockoptInt(fd int, level, opt SocketOption, value int) (err error) {
	return unix.SetsockoptInt(fd, int(level), int(opt), value)
}
