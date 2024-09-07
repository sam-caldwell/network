//go:build linux

package core

// Len - return length of VfMac struct
func (msg *VfMac) Len() int {
	return SizeOfVfMac
}
