package core

// IfLaExtMaskEnum - New extended info filters for IFLA_EXT_MASK
// https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
type IfLaExtMaskEnum uint8

const (

	// RtextFilterVf - RTEXT_FILTER_VF
	RtextFilterVf IfLaExtMaskEnum = 1

	// RtextFilterBrVlan - RTEXT_FILTER_BRVLAN
	RtextFilterBrVlan IfLaExtMaskEnum = 2

	// RtextFilterBrVlanCompressed - RTEXT_FILTER_BRVLAN_COMPRESSED
	RtextFilterBrVlanCompressed IfLaExtMaskEnum = 4

	// RtextFilterSkipStats - RTEXT_FILTER_SKIP_STATS
	RtextFilterSkipStats IfLaExtMaskEnum = 8

	// RtextFilterMrp - RTEXT_FILTER_MRP
	RtextFilterMrp IfLaExtMaskEnum = 16

	// RtextFilterCfmConfig - RTEXT_FILTER_CFM_CONFIG
	RtextFilterCfmConfig IfLaExtMaskEnum = 32

	// RtextFilterCfmStatus - RTEXT_FILTER_CFM_STATUS
	RtextFilterCfmStatus IfLaExtMaskEnum = 64

	// RtextFilterMst - RTEXT_FILTER_MST
	RtextFilterMst IfLaExtMaskEnum = 128
)
