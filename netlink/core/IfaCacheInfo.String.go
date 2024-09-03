package core

import (
	"fmt"
)

// String provides a human-readable representation of IfaCacheInfo.
//
// From unix package...
//
//	type IfaCacheinfo struct {
//	   Prefered uint32
//	   Valid    uint32
//	   Cstamp   uint32
//	   Tstamp   uint32
//	}
//
// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_addr.h
func (msg *IfaCacheInfo) String() string {

	return fmt.Sprintf("Valid: %d, Prefered: %d, Created: %d, Updated: %d",
		msg.Valid, msg.Prefered, msg.Cstamp, msg.Tstamp)

}
