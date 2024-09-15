package core

// Error - return string for Errors
func (err Errors) Error() string {
	switch err {
	case E2BIG:
		return "E2BIG"
	case EACCES:
		return "EACCES"
	case EAGAIN:
		return "EAGAIN"
	case EBADF:
		return "EBADF"
	case EBUSY:
		return "EBUSY"
	case ECHILD:
		return "ECHILD"
	case EDOM:
		return "EDOM"
	case EEXIST:
		return "EEXIST"
	case EFAULT:
		return "EFAULT"
	case EFBIG:
		return "EFBIG"
	case EINTR:
		return "EINTR"
	case EINVAL:
		return "EINVAL"
	case EIO:
		return "EIO"
	case EISDIR:
		return "EISDIR"
	case EMFILE:
		return "EMFILE"
	case EMLINK:
		return "EMLINK"
	case ENFILE:
		return "ENFILE"
	case ENODEV:
		return "ENODEV"
	case ENOENT:
		return "ENOENT"
	case ENOEXEC:
		return "ENOEXEC"
	case ENOMEM:
		return "ENOMEM"
	case ENOSPC:
		return "ENOSPC"
	case ENOTBLK:
		return "ENOTBLK"
	case ENOTDIR:
		return "ENOTDIR"
	case ENOTTY:
		return "ENOTTY"
	case ENXIO:
		return "ENXIO"
	case EPERM:
		return "EPERM"
	case EPIPE:
		return "EPIPE"
	case ERANGE:
		return "ERANGE"
	case EROFS:
		return "EROFS"
	case ESPIPE:
		return "ESPIPE"
	case ESRCH:
		return "ESRCH"
	case ETXTBSY:
		return "ETXTBSY"
	//case EWOULDBLOCK:
	//	return "EWOULDBLOCK"
	case EXDEV:
		return "EXDEV"
	default:
		return "invalid error"
	}
}
