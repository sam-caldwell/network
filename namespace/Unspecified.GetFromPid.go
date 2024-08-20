//go:build !linux

package namespace

func GetFromPid(pid int) (Handle, error) {
	return closedHandle, errors.New(ErrNotImplemented)
}
