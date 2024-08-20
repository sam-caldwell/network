//go:build !linux

package namespace

// SetNamespace - set the namespace
func SetNamespace(namespace Handle, namespaceType int) error {
	return errors.New(ErrNotImplemented)
}
