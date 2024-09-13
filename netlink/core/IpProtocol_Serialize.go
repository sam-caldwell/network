package core

// Serialize converts the IPProto enum value to a byte slice.
// This can be useful when serializing network protocols that require the protocol field
// to be transmitted or used in other operations as a single byte.
//
// For example, this could be used when encoding a protocol identifier for network
// packet manipulation.
//
// Returns a byte slice containing the serialized IPProto.
func (p *IpProtocol) Serialize() ([]byte, error) {
	result := make([]byte, 1)
	result[0] = byte(*p)
	return result, nil
}
