//go:build linux

package core

// Serialize - Serialize VfTxRate to []byte
func (msg *VfTxRate) Serialize() ([]byte, error) {
	buf := make([]byte, SizeOfVfTxRate)

	// Serialize the Vf field
	NativeEndian.PutUint32(buf[0:4], msg.Vf)

	// Serialize the Rate field
	NativeEndian.PutUint32(buf[4:8], msg.Rate)

	return buf, nil
}
