package netlink

import (
	"github.com/sam-caldwell/network/namespace"
	"golang.org/x/sys/unix"
)

// AddressSubscribeOptions - contains a set of options to use with AddressSubscribeWithOptions.
type AddressSubscribeOptions struct {
	Namespace              *namespace.Handle
	ErrorCallback          func(error)
	ListExisting           bool
	ReceiveBufferSize      int
	ReceiveBufferForceSize bool
	ReceiveTimeout         *unix.Timeval
}
