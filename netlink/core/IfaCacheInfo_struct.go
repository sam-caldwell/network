package core

import (
	"golang.org/x/sys/unix"
)

// IfaCacheInfo - Wrapper around unix.IfaCacheinfo with additional functionality.
//
// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_addr.h
type IfaCacheInfo struct {
	unix.IfaCacheinfo
}
