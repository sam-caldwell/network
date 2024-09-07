package netlink

// SupportsNetlinkFamily - return boolean whether the passed netlink family is supported by this Handle
func (h *Handle) SupportsNetlinkFamily(nlFamily int) bool {
	_, ok := h.sockets[nlFamily]
	return ok
}
