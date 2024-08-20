//go:build !linux

package namespace

func New() (Handle, error) {
	return closedHandle, errors.New(ErrNotImplemented)
}
