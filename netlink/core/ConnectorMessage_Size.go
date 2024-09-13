package core

import "unsafe"

// ConnectorMessageSize - size of ConnectorMessage struct
const ConnectorMessageSize = int(unsafe.Sizeof(ConnectorMessage{}))
