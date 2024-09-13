package core

import "unsafe"

// BridgeVlanInfoSize - size of BridgeVlanInfo
const BridgeVlanInfoSize = int(unsafe.Sizeof(BridgeVlanInfo{}))
