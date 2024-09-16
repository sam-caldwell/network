package core

// IfLaExtMask - New extended info filters for IFLA_EXT_MASK
// https://github.com/torvalds/linux/blob/master/include/uapi/linux/rtnetlink.h
type IfLaExtMask uint8

const (

	// RtextFilterVf - RTEXT_FILTER_VF
	RtextFilterVf IfLaExtMask = 1

	// RtextFilterBrVlan - RTEXT_FILTER_BRVLAN
	RtextFilterBrVlan IfLaExtMask = 2

	// RtextFilterBrVlanCompressed - RTEXT_FILTER_BRVLAN_COMPRESSED
	RtextFilterBrVlanCompressed IfLaExtMask = 4

	// RtextFilterSkipStats - RTEXT_FILTER_SKIP_STATS
	RtextFilterSkipStats IfLaExtMask = 8

	// RtextFilterMrp - RTEXT_FILTER_MRP
	RtextFilterMrp IfLaExtMask = 16

	// RtextFilterCfmConfig - RTEXT_FILTER_CFM_CONFIG
	RtextFilterCfmConfig IfLaExtMask = 32

	// RtextFilterCfmStatus - RTEXT_FILTER_CFM_STATUS
	RtextFilterCfmStatus IfLaExtMask = 64

	// RtextFilterMst - RTEXT_FILTER_MST
	RtextFilterMst IfLaExtMask = 128
)
