package core

// GenericNetlinkControlAttributeMcastGrp - Enumeration for Generic Netlink Multicast Group Attributes
//
// This enumeration defines the attributes related to multicast groups in the Generic Netlink (Genl) framework.
// These attributes are used to manage and identify multicast groups associated with a Netlink family.
// See https://docs.kernel.org/userspace-api/netlink/intro.html
type GenericNetlinkControlAttributeMcastGrp uint8

const (
	// GenericNetlinkControlAttributeMcastGrpUnspec - Unspecified attribute.
	// This is a placeholder for an unspecified or unknown multicast group attribute.
	// See https://docs.kernel.org/userspace-api/netlink/intro.html
	GenericNetlinkControlAttributeMcastGrpUnspec GenericNetlinkControlAttributeMcastGrp = 0

	// GenericNetlinkControlAttributeMcastGrpName - Multicast group name attribute.
	// This attribute specifies the name of the multicast group, used to identify and reference the group.
	// See https://docs.kernel.org/userspace-api/netlink/intro.html
	GenericNetlinkControlAttributeMcastGrpName GenericNetlinkControlAttributeMcastGrp = 1

	// GenericNetlinkControlAttributeMcastGrpId - Multicast group ID attribute.
	// This attribute specifies the identifier (ID) of the multicast group, used to uniquely identify the group
	// within the Netlink family.
	// See https://docs.kernel.org/userspace-api/netlink/intro.html
	GenericNetlinkControlAttributeMcastGrpId GenericNetlinkControlAttributeMcastGrp = 2
)
