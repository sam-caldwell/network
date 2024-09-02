package core

// VfRssQueryEn - represent the ifla_vf_rss_query_en attribute from the Linux kernel.
//
// It contains the Virtual Function (VF) index and the RSS query enable setting.
//
// See https://github.com/torvalds/linux/blob/master/include/uapi/linux/if_link.h
type VfRssQueryEn struct {
	Vf      uint32 // VF index
	Setting uint32 // RSS query enable setting (1 for enabled, 0 for disabled)
}
