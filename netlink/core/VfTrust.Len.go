//go:build linux

package core

// Len - return the size of the VfTrust struct
func (msg *VfTrust) Len() int {
	return SizeofVfTrust
}
