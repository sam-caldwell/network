package namespace

// IsOpen returns true if Close() the handle has not been closed (-1)
func (h *Handle) IsOpen() bool {
	return *h != closedHandle
}
