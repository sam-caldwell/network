package netlink

import "github.com/sam-caldwell/network/netlink/core"

// Handle - a handle for the netlink requests on a specific network namespace. Requests on the same netlink family
// share the same netlink socket, which gets released when the handle is Closed.
type Handle struct {
	sockets      map[int]*core.SocketHandle
	lookupByDump bool
}
