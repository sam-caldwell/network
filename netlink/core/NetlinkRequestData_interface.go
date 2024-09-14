package core

// NetlinkRequestData - interface that abstracts the data included in a netlink request, allowing different types of
// data to be encapsulated and sent as part of a netlink message.
type NetlinkRequestData interface {
	Len() int
	Serialize() ([]byte, error)
}
