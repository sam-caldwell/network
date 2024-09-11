package namespace

// IsOpen returns true if Close() the handle has not been closed (-1)
func (ns *Handle) IsOpen() bool {
	return *ns != closedHandle
}
