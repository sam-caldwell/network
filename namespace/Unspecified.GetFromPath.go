//go:build !linux

package namespace

func GetFromPath(path string) (Handle, error) {
	return closedHandle, errors.New(ErrNotImplemented)
}
