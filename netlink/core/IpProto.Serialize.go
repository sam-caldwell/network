package core

// Serialize converts the IPProto enum value to a byte slice.
// This can be useful when serializing network protocols that require the protocol field
// to be transmitted or used in other operations as a single byte.
//
// For example, this could be used when encoding a protocol identifier for network
// packet manipulation.
//
// Returns a byte slice containing the serialized IPProto.
func (i *IPProto) Serialize() []byte {
	arr := make([]byte, 1)
	arr[0] = byte(*i)
	return arr
}
