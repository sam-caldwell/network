package core

import (
	"golang.org/x/sys/unix"
)

// Signal -
type Signal int

const (
	SIGABRT = Signal(unix.SIGABRT) // 0x6
	SIGALRM = Signal(unix.SIGALRM) // 0xe
	SIGFPE  = Signal(unix.SIGFPE)  // 0x8
	SIGHUP  = Signal(unix.SIGHUP)  // 0x1
	SIGILL  = Signal(unix.SIGILL)  // 0x4
	SIGINT  = Signal(unix.SIGINT)  // 0x2
	SIGIOT  = Signal(unix.SIGIOT)  // 0x6
	SIGKILL = Signal(unix.SIGKILL) // 0x9
	SIGPIPE = Signal(unix.SIGPIPE) // 0xd
	SIGQUIT = Signal(unix.SIGQUIT) // 0x3
	SIGSEGV = Signal(unix.SIGSEGV) // 0xb
	SIGTERM = Signal(unix.SIGTERM) // 0xf
	SIGTRAP = Signal(unix.SIGTRAP) // 0x5
)
