package core

// Deserialize - deserialize the interface address message
func (msg *IfAddressMessage) Deserialize(b []byte) (err error) {
	var result *IfAddressMessage
	if result, err = DeserializeIfAddressMessage(b); err != nil {
		return err
	}
	*msg = *result
	return nil
}
