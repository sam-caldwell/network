//go:build linux

package core

// VfTxRate - port of ifla_vf_tx_rate - https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
//
//	struct ifla_vf_tx_rate {
//	  __u32 vf;
//	  __u32 rate; /* Max TX bandwidth in Mbps, 0 disables throttling */
//	};
type VfTxRate struct {
	Vf   uint32
	Rate uint32
}
