//go:build linux

package core

// Serialize - Serialize VfMac to []byte
func (msg *VfMac) Serialize() ([]byte, error) {
	buf := make([]byte, SizeOfVfMac)

	// Serialize the Vf field
	NativeEndian.PutUint32(buf[0:4], msg.Vf)

	// Serialize the Mac field
	copy(buf[4:], msg.Mac[:])

	return buf, nil
}
