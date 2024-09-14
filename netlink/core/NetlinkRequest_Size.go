package core

import "unsafe"

// NetlinkRequestSize - Size of the NetlinkRequest struct
const NetlinkRequestSize = int(unsafe.Sizeof(NetlinkRequest{}))
