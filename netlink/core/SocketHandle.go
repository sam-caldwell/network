//go:build linux

package core

// SocketHandle contains the netlink socket and the associated sequence counter for a specific netlink family
type SocketHandle struct {
	Seq    uint32
	Socket *NetlinkSocket
}
