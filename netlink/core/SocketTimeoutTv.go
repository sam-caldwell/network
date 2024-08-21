//go:build linux

package core

import (
	"golang.org/x/sys/unix"
)

// SocketTimeoutTv - Default netlink socket timeout, 60s
var SocketTimeoutTv = unix.Timeval{
	Sec:  60,
	Usec: 0,
}
