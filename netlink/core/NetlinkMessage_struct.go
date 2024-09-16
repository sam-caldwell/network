package core

// NetlinkMessage - Full netlink message (with header).
//
// This had to be created in this project because it does not exist in the unix package yet,
// and we are trying to avoid use of syscall which is being deprecated.
type NetlinkMessage struct {
	Header NetlinkMessageHeader
	Data   []byte
}
