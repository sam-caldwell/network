package core

// DeserializeIfaCacheInfo - deserialize the IfaCacheInfo from a byte slice
func (msg *IfaCacheInfo) DeserializeIfaCacheInfo(b []byte) {
	*msg = *DeserializeIfaCacheInfo(b)
}
