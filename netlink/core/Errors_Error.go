package core

import "errors"

// Error - return string for Errors
func (err Errors) Error() string {
	switch {
	case errors.Is(err, E2BIG):
		return "E2BIG"
	case errors.Is(err, EACCES):
		return "EACCES"
	case errors.Is(err, EAGAIN):
		return "EAGAIN"
	case errors.Is(err, EBADF):
		return "EBADF"
	case errors.Is(err, EBUSY):
		return "EBUSY"
	case errors.Is(err, ECHILD):
		return "ECHILD"
	case errors.Is(err, EDOM):
		return "EDOM"
	case errors.Is(err, EEXIST):
		return "EEXIST"
	case errors.Is(err, EFAULT):
		return "EFAULT"
	case errors.Is(err, EFBIG):
		return "EFBIG"
	case errors.Is(err, EINTR):
		return "EINTR"
	case errors.Is(err, EINVAL):
		return "EINVAL"
	case errors.Is(err, EIO):
		return "EIO"
	case errors.Is(err, EISDIR):
		return "EISDIR"
	case errors.Is(err, EMFILE):
		return "EMFILE"
	case errors.Is(err, EMLINK):
		return "EMLINK"
	case errors.Is(err, ENFILE):
		return "ENFILE"
	case errors.Is(err, ENODEV):
		return "ENODEV"
	case errors.Is(err, ENOENT):
		return "ENOENT"
	case errors.Is(err, ENOEXEC):
		return "ENOEXEC"
	case errors.Is(err, ENOMEM):
		return "ENOMEM"
	case errors.Is(err, ENOSPC):
		return "ENOSPC"
	case errors.Is(err, ENOTBLK):
		return "ENOTBLK"
	case errors.Is(err, ENOTDIR):
		return "ENOTDIR"
	case errors.Is(err, ENOTTY):
		return "ENOTTY"
	case errors.Is(err, ENXIO):
		return "ENXIO"
	case errors.Is(err, EPERM):
		return "EPERM"
	case errors.Is(err, EPIPE):
		return "EPIPE"
	case errors.Is(err, ERANGE):
		return "ERANGE"
	case errors.Is(err, EROFS):
		return "EROFS"
	case errors.Is(err, ESPIPE):
		return "ESPIPE"
	case errors.Is(err, ESRCH):
		return "ESRCH"
	case errors.Is(err, ETXTBSY):
		return "ETXTBSY"
		//case EWOULDBLOCK:
	//	return "EWOULDBLOCK"
	case errors.Is(err, EXDEV):
		return "EXDEV"
	default:
		return "invalid error"
	}
}
