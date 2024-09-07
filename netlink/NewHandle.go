package netlink

import "github.com/sam-caldwell/network/namespace"

// NewHandle returns a netlink handle on the current network namespace.
// Caller may specify the netlink families the handle should support.
// If no families are specified, all the families the netlink package
// supports will be automatically added.
func NewHandle(nlFamilies ...int) (*Handle, error) {
	return newHandle(namespace.None(), namespace.None(), nlFamilies...)
}
