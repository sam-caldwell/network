package core

import "unsafe"

// ConnectorMessageOperationSize - size of ConnectorMessageOperation struct
const ConnectorMessageOperationSize = int(unsafe.Sizeof(ConnectorMessageOperation{}))
