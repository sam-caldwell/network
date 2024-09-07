//go:build linux

package core

// Len - Return the size of the VfTxRate struct
func (msg *VfTxRate) Len() int {
	return SizeOfVfTxRate
}
