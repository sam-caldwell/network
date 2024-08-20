//go:build !linux

package namespace

// Set - set current network namespace to the namespace represented by NamespaceHandle.
func Set(ns Handle) error {
	return errors.New(ErrNotImplemented)
}
