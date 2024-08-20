//go:build !linux

package namespace

// Handle - handle to a network namespace. It can be cast directly
// to an int and used as a file descriptor.
type Handle int

// Close - close the NamespaceHandle and reset its file descriptor to -1 (closedHandle).
//
// WARNING: DO NOT USE AFTER Close() is called.  Bad things will happen.
func (ns *Handle) Close() error {
	return errors.New(ErrNotImplemented)
}

// IsOpen returns true if Close() the handle has not been closed (-1)
func (ns *Handle) IsOpen() bool {
	return errors.New(ErrNotImplemented)
}

// None - return an empty (-1: closed) namespace handle.
func None() Handle {
	return errors.New(ErrNotImplemented)
}

// String - return the file descriptor, its dev and inode, as a string.
func (ns *Handle) String() string {
	return errors.New(ErrNotImplemented)
}

// UniqueId return string uniquely identifying the associated namespace
func (ns *Handle) UniqueId() string {
	return errors.New(ErrNotImplemented)
}
