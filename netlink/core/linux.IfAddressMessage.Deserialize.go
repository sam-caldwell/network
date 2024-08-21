package core

// Deserialize - deserialize the interface address message
func (msg *IfAddressMessage) Deserialize(b []byte) {
	*msg = *DeserializeIfAddressMessage(b)
}
