package core

// NetlinkRouteAttr - adaptation of syscall.NetlinkRouteAttr since Golang seems to be moving to unix.*
// but this particular struct has not been migrated yet.
type NetlinkRouteAttr struct {
	Attr  RtAttr
	Value []byte
}
