package core

// NlaPolicy represents the Netlink attribute policy used for validation.
type NlaPolicy struct {

	// Type - of the attribute (e.g., NLA_U8, NLA_U16)
	Type NlaType

	// Length - of the attribute (if applicable)
	Len int

	// StrictStartType - Strict start type for certain attributes
	StrictStartType IflaGeneveEnum
}
