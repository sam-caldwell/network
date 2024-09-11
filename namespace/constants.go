package namespace

const (
	// bindMountPath - Bind mount path for named netns
	bindMountPath = "/run/netns"

	// processNamespacePath - Path to the ns namespace file
	processNamespacePath = "/proc/%d/ns/net"

	// threadNamespacePath - Path to the ns namespace file
	threadNamespacePath = "/proc/%d/task/%d/ns/net"

	// closedHandle - the file handle used when a namespace file descriptor is closed.
	closedHandle = Handle(-1)

	// ErrNotImplemented - an error indicating some functionality is not implemented
	ErrNotImplemented = "not implemented"
)
