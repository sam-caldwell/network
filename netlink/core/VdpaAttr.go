package core

type VdpaAttr uint8

const (
	// VDPAAttrUnspec - VDPA_ATTR_UNSPEC
	// VDPAAttrUnspec represents an unspecified attribute.
	// Reference: https://github.com/torvalds/linux/blob/master/include/uapi/linux/vdpa.h
	VDPAAttrUnspec VdpaAttr = iota

	// VDPAAttrMgmtDevBusName - VDPA_ATTR_MGMTDEV_BUS_NAME
	// VDPAAttrMgmtDevBusName represents the management device bus name attribute.
	// Reference: https://github.com/torvalds/linux/blob/master/include/uapi/linux/vdpa.h
	VDPAAttrMgmtDevBusName

	// VDPAAttrMgmtDevDevName - VDPA_ATTR_MGMTDEV_DEV_NAME
	// VDPAAttrMgmtDevDevName represents the management device name attribute.
	// Reference: https://github.com/torvalds/linux/blob/master/include/uapi/linux/vdpa.h
	VDPAAttrMgmtDevDevName

	// VDPAAttrMgmtDevSupportedClasses - VDPA_ATTR_MGMTDEV_SUPPORTED_CLASSES
	// VDPAAttrMgmtDevSupportedClasses represents the supported classes of the management device.
	// Reference: https://github.com/torvalds/linux/blob/master/include/uapi/linux/vdpa.h
	VDPAAttrMgmtDevSupportedClasses

	// VDPAAttrDevName - VDPA_ATTR_DEV_NAME
	// VDPAAttrDevName represents the device name attribute.
	// Reference: https://github.com/torvalds/linux/blob/master/include/uapi/linux/vdpa.h
	VDPAAttrDevName

	// VDPAAttrDevID - VDPA_ATTR_DEV_ID
	// VDPAAttrDevID represents the device ID attribute.
	// Reference: https://github.com/torvalds/linux/blob/master/include/uapi/linux/vdpa.h
	VDPAAttrDevID

	// VDPAAttrDevVendorID - VDPA_ATTR_DEV_VENDOR_ID
	// VDPAAttrDevVendorID represents the device vendor ID attribute.
	// Reference: https://github.com/torvalds/linux/blob/master/include/uapi/linux/vdpa.h
	VDPAAttrDevVendorID

	// VDPAAttrDevMaxVqs - VDPA_ATTR_DEV_MAX_VQS
	// VDPAAttrDevMaxVqs represents the maximum number of VQs supported by the device.
	// Reference: https://github.com/torvalds/linux/blob/master/include/uapi/linux/vdpa.h
	VDPAAttrDevMaxVqs

	// VDPAAttrDevMaxVqSize - VDPA_ATTR_DEV_MAX_VQ_SIZE
	// VDPAAttrDevMaxVqSize represents the maximum VQ size supported by the device.
	// Reference: https://github.com/torvalds/linux/blob/master/include/uapi/linux/vdpa.h
	VDPAAttrDevMaxVqSize

	// VDPAAttrDevMinVqSize - VDPA_ATTR_DEV_MIN_VQ_SIZE
	// VDPAAttrDevMinVqSize represents the minimum VQ size supported by the device.
	// Reference: https://github.com/torvalds/linux/blob/master/include/uapi/linux/vdpa.h
	VDPAAttrDevMinVqSize

	// VDPAAttrDevNetCfgMacaddr - VDPA_ATTR_DEV_NET_CFG_MACADDR
	// VDPAAttrDevNetCfgMacaddr represents the MAC address configuration for the network device.
	// Reference: https://github.com/torvalds/linux/blob/master/include/uapi/linux/vdpa.h
	VDPAAttrDevNetCfgMacaddr

	// VDPAAttrDevNetStatus - VDPA_ATTR_DEV_NET_STATUS
	// VDPAAttrDevNetStatus represents the status of the network device.
	// Reference: https://github.com/torvalds/linux/blob/master/include/uapi/linux/vdpa.h
	VDPAAttrDevNetStatus

	// VDPAAttrDevNetCfgMaxVqp - VDPA_ATTR_DEV_NET_CFG_MAX_VQP
	// VDPAAttrDevNetCfgMaxVqp represents the maximum number of VQs supported per network configuration.
	// Reference: https://github.com/torvalds/linux/blob/master/include/uapi/linux/vdpa.h
	VDPAAttrDevNetCfgMaxVqp

	// VDPAAttrDevNetCfgMtu - VDPA_ATTR_DEV_NET_CFG_MTU
	// VDPAAttrDevNetCfgMtu represents the MTU configuration for the network device.
	// Reference: https://github.com/torvalds/linux/blob/master/include/uapi/linux/vdpa.h
	VDPAAttrDevNetCfgMtu

	// VDPAAttrDevNegotiatedFeatures - VDPA_ATTR_DEV_NEGOTIATED_FEATURES
	// VDPAAttrDevNegotiatedFeatures represents the negotiated features supported by the device.
	// Reference: https://github.com/torvalds/linux/blob/master/include/uapi/linux/vdpa.h
	VDPAAttrDevNegotiatedFeatures

	// VDPAAttrDevMgmtDevMaxVqs - VDPA_ATTR_DEV_MGMTDEV_MAX_VQS
	// VDPAAttrDevMgmtDevMaxVqs represents the maximum number of VQs supported by the management device.
	// Reference: https://github.com/torvalds/linux/blob/master/include/uapi/linux/vdpa.h
	VDPAAttrDevMgmtDevMaxVqs

	// VDPAAttrDevSupportedFeatures - VDPA_ATTR_DEV_SUPPORTED_FEATURES
	// VDPAAttrDevSupportedFeatures represents the features supported by the device.
	// Reference: https://github.com/torvalds/linux/blob/master/include/uapi/linux/vdpa.h
	VDPAAttrDevSupportedFeatures

	// VDPAAttrDevQueueIndex - VDPA_ATTR_DEV_QUEUE_INDEX
	// VDPAAttrDevQueueIndex represents the index of the device queue.
	// Reference: https://github.com/torvalds/linux/blob/master/include/uapi/linux/vdpa.h
	VDPAAttrDevQueueIndex

	// VDPAAttrDevVendorAttrName - VDPA_ATTR_DEV_VENDOR_ATTR_NAME
	// VDPAAttrDevVendorAttrName represents the vendor-specific attribute name.
	// Reference: https://github.com/torvalds/linux/blob/master/include/uapi/linux/vdpa.h
	VDPAAttrDevVendorAttrName

	// VDPAAttrDevVendorAttrValue - VDPA_ATTR_DEV_VENDOR_ATTR_VALUE
	// VDPAAttrDevVendorAttrValue represents the vendor-specific attribute value.
	// Reference: https://github.com/torvalds/linux/blob/master/include/uapi/linux/vdpa.h
	VDPAAttrDevVendorAttrValue

	// VDPAAttrDevFeatures - VDPA_ATTR_DEV_FEATURES
	// VDPAAttrDevFeatures represents the features attribute of the device.
	// Reference: https://github.com/torvalds/linux/blob/master/include/uapi/linux/vdpa.h
	VDPAAttrDevFeatures
)
