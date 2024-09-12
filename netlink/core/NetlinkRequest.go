package core

import (
	"golang.org/x/sys/unix"
)

// NetlinkRequest - represents a request message that can be sent via netlink.
type NetlinkRequest struct {

	// unix.NlMsghdr - represents a request message that can be sent via netlink.
	unix.NlMsghdr

	// Data - The payload data to include in the message.
	Data []NetlinkRequestData

	// RawData - The sequence number for the request.
	RawData []byte

	// Sockets - Additional flags for the request.
	Sockets map[IpProtocol]*SocketHandle
}
