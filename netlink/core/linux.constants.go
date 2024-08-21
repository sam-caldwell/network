package core

import "golang.org/x/sys/unix"

const (
	// SizeofIfAddrmsg     = 0x8 // bytes as derived from unix.SizeofIfAddrmsg
	SizeofIfAddrmsg = unix.SizeofIfAddrmsg
)
