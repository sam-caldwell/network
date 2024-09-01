package core

// Len - return the length of the VfVlan struct
func (msg *VfVlan) Len() int {
	return SizeofVfVlan
}
