//go:build !linux

package namespace

func DeleteNamespace(name string) error {
	return errors.New(ErrNotImplemented)
}
