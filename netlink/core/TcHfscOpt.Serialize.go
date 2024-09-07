package core

// Serialize - Serializes the TcHfscOpt struct into a byte slice using binary encoding.
// This method ensures safe serialization by manually encoding each field into its binary format.
func (x *TcHfscOpt) Serialize() ([]byte, error) {
	buf := make([]byte, 2)
	NativeEndian.PutUint16(buf[0:], x.Defcls)
	return buf, nil
}
