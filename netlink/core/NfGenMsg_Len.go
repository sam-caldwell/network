//go:build linux

package core

// Len - return the size of the Nfgenmsg struct
func (msg *Nfgenmsg) Len() int {
	return NfGenMsgSize
}
