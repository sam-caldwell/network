//go:build !linux

package namespace

func GetFromName(name string) (Handle, error) {
	return closedHandle, errors.New(ErrNotImplemented)
}
