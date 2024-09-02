//go:build linux

package core

// Close closes the netlink socket
func (sh *SocketHandle) Close() error {
	if sh.Socket != nil {
		return sh.Socket.Close()
	}
	return nil
}
