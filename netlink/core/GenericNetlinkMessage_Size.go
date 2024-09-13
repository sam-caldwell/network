package core

import "unsafe"

// GenericNetlinkMessageSize - size of GenericNetlinkMessage struct
const GenericNetlinkMessageSize = int(unsafe.Sizeof(GenericNetlinkMessage{}))
