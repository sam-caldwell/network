//go:build linux

package core

// VfGUID - see https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
//
//	struct ifla_vf_guid {
//	  __u32 vf;
//	  __u32 rsvd;
//	  __u64 guid;
//	};
type VfGUID struct {
	Vf   uint32
	Rsvd uint32
	GUID uint64
}
