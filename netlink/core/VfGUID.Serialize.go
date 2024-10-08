//go:build linux

package core

// Serialize - Serialize VfGUID to []byte
func (msg *VfGUID) Serialize() ([]byte, error) {
	buf := make([]byte, SizeOfVfGUID)

	// Serialize the Vf field
	NativeEndian.PutUint32(buf[0:4], msg.Vf)

	// Serialize the Rsvd field
	NativeEndian.PutUint32(buf[4:8], msg.Rsvd)

	// Serialize the GUID field
	NativeEndian.PutUint64(buf[8:16], msg.GUID)

	return buf, nil
}
