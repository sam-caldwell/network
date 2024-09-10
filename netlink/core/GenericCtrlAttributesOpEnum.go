package core

// GenericCtrlAttributesOpEnum - Enumeration for Generic Netlink Control Operation Attributes
//
// This enumeration defines various attributes used in the Generic Netlink (Genl) control operations.
// These attributes are utilized by the Generic Netlink controller to manage and query the operations of different
// Netlink families.
// See https://docs.kernel.org/networking/gtp.html
type GenericCtrlAttributesOpEnum uint8

const (
	// GenlCtrlAttrOpUnspec - Unspecified attribute.
	// This is a placeholder for an unspecified or unknown attribute in the Generic Netlink control operation.
	//
	// See https://docs.kernel.org/networking/gtp.html
	GenlCtrlAttrOpUnspec GenericCtrlAttributesOpEnum = 0

	// GenlCtrlAttrOpId - Operation identifier attribute.
	// This attribute specifies the identifier for a particular operation in the Generic Netlink control.
	//
	// See https://docs.kernel.org/networking/gtp.html
	GenlCtrlAttrOpId GenericCtrlAttributesOpEnum = 1

	// GenlCtrlAttrOpFlags - Operation flags attribute.
	// This attribute specifies the flags associated with a particular operation, defining its behavior or
	// capabilities.
	//
	// See https://docs.kernel.org/networking/gtp.html
	GenlCtrlAttrOpFlags GenericCtrlAttributesOpEnum = 2
)
