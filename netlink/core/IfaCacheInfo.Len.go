package core

// Len - return the length (in bytes) of the IfaCacheInfo struct.
func (msg *IfaCacheInfo) Len() int {
	return SizeOfIfaCacheinfo
}
