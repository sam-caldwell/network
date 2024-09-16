package core

// IteratorFunctionPtr - Function pattern used by NetlinkRequest.Execute() and NetlinkRequest.ExecuteIter()
type IteratorFunctionPtr func(msg []byte) bool
