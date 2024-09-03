package core

// GenlCtrlAttrMcastGrp - Enumeration for Generic Netlink Multicast Group Attributes
//
// This enumeration defines the attributes related to multicast groups in the Generic Netlink (Genl) framework.
// These attributes are used to manage and identify multicast groups associated with a Netlink family.
// See https://docs.kernel.org/userspace-api/netlink/intro.html
type GenlCtrlAttrMcastGrp uint8

const (
	// GenlCtrlAttrMcastGrpUnspec - Unspecified attribute.
	// This is a placeholder for an unspecified or unknown multicast group attribute.
	// See https://docs.kernel.org/userspace-api/netlink/intro.html
	GenlCtrlAttrMcastGrpUnspec GenlCtrlAttrMcastGrp = 0

	// GenlCtrlAttrMcastGrpName - Multicast group name attribute.
	// This attribute specifies the name of the multicast group, used to identify and reference the group.
	// See https://docs.kernel.org/userspace-api/netlink/intro.html
	GenlCtrlAttrMcastGrpName GenlCtrlAttrMcastGrp = 1

	// GenlCtrlAttrMcastGrpId - Multicast group ID attribute.
	// This attribute specifies the identifier (ID) of the multicast group, used to uniquely identify the group
	// within the Netlink family.
	// See https://docs.kernel.org/userspace-api/netlink/intro.html
	GenlCtrlAttrMcastGrpId GenlCtrlAttrMcastGrp = 2
)
