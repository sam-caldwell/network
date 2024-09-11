package namespace

const (
	//bindMountPath - Bind mount path for named netns
	bindMountPath = "/run/netns"

	//closedHandle - the file handle used when a namespace file descriptor is closed.
	closedHandle = Handle(-1)

	//ErrNotImplemented - an error indicating some functionality is not implemented
	ErrNotImplemented = "not implemented"
)
