package core

import "golang.org/x/sys/unix"

const (
	// SizeofIfAddrmsg     = 0x8 // bytes as derived from unix.SizeofIfAddrmsg
	SizeofIfAddrmsg = unix.SizeofIfAddrmsg

	// SizeOfIfaCacheinfo = 0x16 // bytes as derived from unix.SizeofIfaCacheinfo
	SizeOfIfaCacheinfo = unix.SizeofIfaCacheinfo
)
