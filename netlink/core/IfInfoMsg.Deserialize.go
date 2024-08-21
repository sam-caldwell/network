package core

// Deserialize - Deserialize byte slice into IfInfoMsg struct
func (msg *IfInfoMsg) Deserialize(b []byte) (err error) {
	var result *IfInfoMsg
	if result, err = DeserializeIfInfoMsg(b); err != nil {
		return err
	}
	*msg = *result
	return nil
}
