package core

import "unsafe"

// Len returns the length of the Genlmsg structure in bytes.
func (msg *Genlmsg) Len() int {
	return int(unsafe.Sizeof(Genlmsg{}))
}
