//go:build linux

package core

import "github.com/sam-caldwell/network/namespace"

// SubscribeAt - works like Subscribe() plus lets the caller choose the network namespace in which the socket would
// be opened (newNamespace). Then control goes back to currentNamespace if open, otherwise to the namespace at the
// time this function was called.
func SubscribeAt(newNamespace, currentNamespace namespace.Handle,
	protocol int, groups ...uint) (*NetlinkSocket, error) {

	c, err := executeInNetNamespace(newNamespace, currentNamespace)

	if err != nil {
		return nil, err
	}

	defer c()

	return Subscribe(protocol, groups...)

}
