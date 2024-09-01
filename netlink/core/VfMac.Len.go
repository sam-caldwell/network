//go:build linux

package core

func (msg *VfMac) Len() int {
	return SizeofVfMac
}
