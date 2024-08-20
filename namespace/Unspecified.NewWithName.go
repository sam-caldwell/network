//go:build !linux

package namespace

func NewWithName(name string) (Handle, error) {
	return closedHandle, errors.New(ErrNotImplemented)
}
