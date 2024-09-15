package core

import "unsafe"

const (
	// AttributeSize is the minimum length an attribute must have.
	AttributeSize = int(unsafe.Sizeof(Attribute{}))
)
