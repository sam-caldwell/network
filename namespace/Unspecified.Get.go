//go:build !linux

package namespace

func Get() (Handle, error) {
	return closedHandle, errors.New(ErrNotImplemented)
}
