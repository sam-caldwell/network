//go:build !linux

package namespace

func GetFromThread(pid int, tid int) (Handle, error) {
	return closedHandle, errors.New(ErrNotImplemented)
}
