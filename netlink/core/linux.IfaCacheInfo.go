package core

import (
	"golang.org/x/sys/unix"
)

// IfaCacheInfo -
// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_addr.h
//
//	struct ifa_cacheinfo {
//		__u32	ifa_prefered;
//		__u32	ifa_valid;
//		__u32	cstamp; /* created timestamp, hundredths of seconds */
//		__u32	tstamp; /* updated timestamp, hundredths of seconds */
//	};
type IfaCacheInfo struct {
	unix.IfaCacheinfo
}
