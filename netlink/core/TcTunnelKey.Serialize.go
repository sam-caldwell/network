package core

// Serialize - Safely converts the TcTunnelKey structure into a byte slice.
// This method uses the binary package to serialize the fields into a byte slice
// while ensuring proper endianness, providing a safe alternative to unsafe pointer casting.
func (msg *TcTunnelKey) Serialize() ([]byte, error) {
	buf := make([]byte, SizeofTcTunnelKey)

	// Serialize each field of the TcTunnelKey struct
	NativeEndian.PutUint32(buf[0:], msg.Index)
	NativeEndian.PutUint32(buf[4:], msg.Capab)
	NativeEndian.PutUint32(buf[8:], uint32(msg.Action))
	NativeEndian.PutUint32(buf[12:], uint32(msg.Refcnt))
	NativeEndian.PutUint32(buf[16:], uint32(msg.Bindcnt))

	return buf, nil
}
