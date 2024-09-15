//go:build linux

package core

import (
	"errors"
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
			actual Errors
			expect Errors
		}{
			{actual: E2BIG, expect: Errors(0x7)},
			{actual: EACCES, expect: Errors(0xd)},
			{actual: EAGAIN, expect: Errors(0xb)},
			{actual: EBADF, expect: Errors(0x9)},
			{actual: EBUSY, expect: Errors(0x10)},
			{actual: ECHILD, expect: Errors(0xa)},
			{actual: EDOM, expect: Errors(0x21)},
			{actual: EEXIST, expect: Errors(0x11)},
			{actual: EFAULT, expect: Errors(0xe)},
			{actual: EFBIG, expect: Errors(0x1b)},
			{actual: EINTR, expect: Errors(0x4)},
			{actual: EINVAL, expect: Errors(0x16)},
			{actual: EIO, expect: Errors(0x5)},
			{actual: EISDIR, expect: Errors(0x15)},
			{actual: EMFILE, expect: Errors(0x18)},
			{actual: EMLINK, expect: Errors(0x1f)},
			{actual: ENFILE, expect: Errors(0x17)},
			{actual: ENODEV, expect: Errors(0x13)},
			{actual: ENOENT, expect: Errors(0x2)},
			{actual: ENOEXEC, expect: Errors(0x8)},
			{actual: ENOMEM, expect: Errors(0xc)},
			{actual: ENOSPC, expect: Errors(0x1c)},
			{actual: ENOTBLK, expect: Errors(0xf)},
			{actual: ENOTDIR, expect: Errors(0x14)},
			{actual: ENOTTY, expect: Errors(0x19)},
			{actual: ENXIO, expect: Errors(0x6)},
			{actual: EPERM, expect: Errors(0x1)},
			{actual: EPIPE, expect: Errors(0x20)},
			{actual: ERANGE, expect: Errors(0x22)},
			{actual: EROFS, expect: Errors(0x1e)},
			{actual: ESPIPE, expect: Errors(0x1d)},
			{actual: ESRCH, expect: Errors(0x3)},
			{actual: ETXTBSY, expect: Errors(0x1a)},
			{actual: EWOULDBLOCK, expect: Errors(0xb)},
			{actual: EXDEV, expect: Errors(0x12)},
		}
		for _, tt := range testData {
			if !errors.Is(tt.expect, tt.actual) {
				t.Fatalf("error mismatch. actual:%v, expected:%v", tt.actual, tt.expect)
			}
		}
	})
}
