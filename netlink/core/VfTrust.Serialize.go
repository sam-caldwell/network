//go:build linux

package core

// Serialize - Serialize VfTrust to []byte
func (msg *VfTrust) Serialize() ([]byte, error) {
	buf := make([]byte, SizeofVfTrust)

	// Serialize the Vf field
	NativeEndian.PutUint32(buf[0:4], msg.Vf)

	// Serialize the Setting field
	NativeEndian.PutUint32(buf[4:8], msg.Setting)

	return buf, nil
}
