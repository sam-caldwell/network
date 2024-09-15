//go:build linux

package core

const (
	ErrInputTooLarge = "input too large"
	ErrInputTooShort = "input too short"
	ErrNilInput      = "nil input"

	ErrInvalidAttributeLength = "invalid attribute length"
	ErrTruncatedAttribute     = "truncated attribute"
)

const (
	E2BIG       = uintptr(0x7)
	EACCES      = uintptr(0xd)
	EAGAIN      = uintptr(0xb)
	EBADF       = uintptr(0x9)
	EBUSY       = uintptr(0x10)
	ECHILD      = uintptr(0xa)
	EDOM        = uintptr(0x21)
	EEXIST      = uintptr(0x11)
	EFAULT      = uintptr(0xe)
	EFBIG       = uintptr(0x1b)
	EINTR       = uintptr(0x4)
	EINVAL      = uintptr(0x16)
	EIO         = uintptr(0x5)
	EISDIR      = uintptr(0x15)
	EMFILE      = uintptr(0x18)
	EMLINK      = uintptr(0x1f)
	ENFILE      = uintptr(0x17)
	ENODEV      = uintptr(0x13)
	ENOENT      = uintptr(0x2)
	ENOEXEC     = uintptr(0x8)
	ENOMEM      = uintptr(0xc)
	ENOSPC      = uintptr(0x1c)
	ENOTBLK     = uintptr(0xf)
	ENOTDIR     = uintptr(0x14)
	ENOTTY      = uintptr(0x19)
	ENXIO       = uintptr(0x6)
	EPERM       = uintptr(0x1)
	EPIPE       = uintptr(0x20)
	ERANGE      = uintptr(0x22)
	EROFS       = uintptr(0x1e)
	ESPIPE      = uintptr(0x1d)
	ESRCH       = uintptr(0x3)
	ETXTBSY     = uintptr(0x1a)
	EWOULDBLOCK = uintptr(0xb)
	EXDEV       = uintptr(0x12)
)
