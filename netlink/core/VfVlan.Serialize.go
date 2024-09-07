package core

// Serialize - serialize VfVlan to []byte
func (msg *VfVlan) Serialize() ([]byte, error) {
	buf := make([]byte, SizeOfVfVlan)

	// Serialize each field into the byte slice
	NativeEndian.PutUint32(buf[0:4], msg.Vf)
	NativeEndian.PutUint32(buf[4:8], msg.Vlan)
	NativeEndian.PutUint32(buf[8:12], msg.Qos)

	return buf, nil
}
