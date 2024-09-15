package core

import "bytes"

// Equal - Compare two Attribute structs and return a boolean true if equal
func (attr *Attribute) Equal(other *Attribute) bool {

	return attr.Type == other.Type && bytes.Equal(attr.Value, other.Value)

}
