package core

// DeserializeIfaCacheInfo - deserialize the IfaCacheInfo from a byte slice
func (msg *IfaCacheInfo) DeserializeIfaCacheInfo(b []byte) error {
	info, err := DeserializeIfaCacheInfo(b)
	if err == nil {
		*msg = *info
	}
	return err
}
