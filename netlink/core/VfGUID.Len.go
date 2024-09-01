//go:build linux

package core

// Len - return the size of the VfGUID struct
func (msg *VfGUID) Len() int {
	return SizeofVfGUID
}
