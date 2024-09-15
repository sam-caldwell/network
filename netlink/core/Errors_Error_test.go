//go:build linux

package core

import "testing"

// TestErrors_Error tests the Error() method of the Errors type.
func TestErrors_Error(t *testing.T) {
	tests := []struct {
		name     string
		err      Errors
		expected string
	}{
		{"E2BIG", E2BIG, "E2BIG"},
		{"EACCES", EACCES, "EACCES"},
		{"EAGAIN", EAGAIN, "EAGAIN"},
		{"EBADF", EBADF, "EBADF"},
		{"EBUSY", EBUSY, "EBUSY"},
		{"ECHILD", ECHILD, "ECHILD"},
		{"EDOM", EDOM, "EDOM"},
		{"EEXIST", EEXIST, "EEXIST"},
		{"EFAULT", EFAULT, "EFAULT"},
		{"EFBIG", EFBIG, "EFBIG"},
		{"EINTR", EINTR, "EINTR"},
		{"EINVAL", EINVAL, "EINVAL"},
		{"EIO", EIO, "EIO"},
		{"EISDIR", EISDIR, "EISDIR"},
		{"EMFILE", EMFILE, "EMFILE"},
		{"EMLINK", EMLINK, "EMLINK"},
		{"ENFILE", ENFILE, "ENFILE"},
		{"ENODEV", ENODEV, "ENODEV"},
		{"ENOENT", ENOENT, "ENOENT"},
		{"ENOEXEC", ENOEXEC, "ENOEXEC"},
		{"ENOMEM", ENOMEM, "ENOMEM"},
		{"ENOSPC", ENOSPC, "ENOSPC"},
		{"ENOTBLK", ENOTBLK, "ENOTBLK"},
		{"ENOTDIR", ENOTDIR, "ENOTDIR"},
		{"ENOTTY", ENOTTY, "ENOTTY"},
		{"ENXIO", ENXIO, "ENXIO"},
		{"EPERM", EPERM, "EPERM"},
		{"EPIPE", EPIPE, "EPIPE"},
		{"ERANGE", ERANGE, "ERANGE"},
		{"EROFS", EROFS, "EROFS"},
		{"ESPIPE", ESPIPE, "ESPIPE"},
		{"ESRCH", ESRCH, "ESRCH"},
		{"ETXTBSY", ETXTBSY, "ETXTBSY"},
		//{"EWOULDBLOCK", EWOULDBLOCK, "EWOULDBLOCK"},
		{"EXDEV", EXDEV, "EXDEV"},
		// Test an undefined error code
		{"InvalidError", Errors(0xDEADBEEF), "invalid error"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errString := tt.err.Error()
			if errString != tt.expected {
				t.Errorf("Error() = %v; want %v", errString, tt.expected)
			}
		})
	}
}
