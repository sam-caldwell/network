package core

// Uint32Attribute  - represents an attribute consisting of a type identifier and a 32-bit unsigned integer value.
type Uint32Attribute struct {
	// Type - specifies the attribute's type identifier, typically used to categorize or define the kind of attribute.
	Type uint16

	// Value - holds the 32-bit unsigned integer value associated with the attribute.
	Value uint32
}
