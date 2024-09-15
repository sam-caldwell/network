//go:build linux

package core

import (
	"testing"
)

func TestError_constants(t *testing.T) {
	t.Run("error strings", func(t *testing.T) {
		testData := []struct {
			actual string
			expect string
		}{
			{actual: ErrNilInput, expect: "nil input"},
			{actual: ErrInputTooShort, expect: "input too short"},
			{actual: ErrInputTooLarge, expect: "input too large"},
			{actual: ErrInvalidAttributeLength, expect: "invalid attribute length"},
			{actual: ErrTruncatedAttribute, expect: "truncated attribute"},
		}
		for _, tt := range testData {
			if tt.actual != tt.expect {
				t.Fatalf("error mismatch. actual:%v, expected:%v", tt.actual, tt.expect)
			}
		}
	})
	t.Run("syscall error numbers", func(t *testing.T) {
		testData := []struct {
			actual uintptr
			expect uintptr
		}{
			{actual: E2BIG, expect: uintptr(0x7)},
			{actual: EACCES, expect: uintptr(0xd)},
			{actual: EAGAIN, expect: uintptr(0xb)},
			{actual: EBADF, expect: uintptr(0x9)},
			{actual: EBUSY, expect: uintptr(0x10)},
			{actual: ECHILD, expect: uintptr(0xa)},
			{actual: EDOM, expect: uintptr(0x21)},
			{actual: EEXIST, expect: uintptr(0x11)},
			{actual: EFAULT, expect: uintptr(0xe)},
			{actual: EFBIG, expect: uintptr(0x1b)},
			{actual: EINTR, expect: uintptr(0x4)},
			{actual: EINVAL, expect: uintptr(0x16)},
			{actual: EIO, expect: uintptr(0x5)},
			{actual: EISDIR, expect: uintptr(0x15)},
			{actual: EMFILE, expect: uintptr(0x18)},
			{actual: EMLINK, expect: uintptr(0x1f)},
			{actual: ENFILE, expect: uintptr(0x17)},
			{actual: ENODEV, expect: uintptr(0x13)},
			{actual: ENOENT, expect: uintptr(0x2)},
			{actual: ENOEXEC, expect: uintptr(0x8)},
			{actual: ENOMEM, expect: uintptr(0xc)},
			{actual: ENOSPC, expect: uintptr(0x1c)},
			{actual: ENOTBLK, expect: uintptr(0xf)},
			{actual: ENOTDIR, expect: uintptr(0x14)},
			{actual: ENOTTY, expect: uintptr(0x19)},
			{actual: ENXIO, expect: uintptr(0x6)},
			{actual: EPERM, expect: uintptr(0x1)},
			{actual: EPIPE, expect: uintptr(0x20)},
			{actual: ERANGE, expect: uintptr(0x22)},
			{actual: EROFS, expect: uintptr(0x1e)},
			{actual: ESPIPE, expect: uintptr(0x1d)},
			{actual: ESRCH, expect: uintptr(0x3)},
			{actual: ETXTBSY, expect: uintptr(0x1a)},
			{actual: EWOULDBLOCK, expect: uintptr(0xb)},
			{actual: EXDEV, expect: uintptr(0x12)},
		}
		for _, tt := range testData {
			if tt.actual != tt.expect {
				t.Fatalf("error mismatch. actual:%v, expected:%v", tt.actual, tt.expect)
			}
		}
	})
}
