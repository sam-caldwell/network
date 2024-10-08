//go:build linux

package core

import (
	"github.com/sam-caldwell/network/namespace"
)

// GetNetlinkSocketAt - open netlink socket in the newNamespace, and positions the thread back into the network
// namespace specified by currentNamespace, when done.
//
// If currentNamespace is closed, the function derives the current namespace and moves back into it when done.
// If newNamespace is close, the socket will be opened in the current network namespace.
func GetNetlinkSocketAt(newNamespace, currentNamespace namespace.Handle, protocol IpProtocol) (*NetlinkSocket, error) {

	c, err := ExecuteInNetNamespace(newNamespace, currentNamespace)

	if err != nil {
		return nil, err
	}

	defer c()

	return GetNetlinkSocket(protocol)

}
