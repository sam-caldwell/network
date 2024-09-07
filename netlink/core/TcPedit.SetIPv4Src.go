package core

import "net"

// SetIPv4Src sets the IPv4 source address for the pedit action.
// This function modifies the packet's IPv4 source address by creating a pedit key
// and an extended pedit key. It handles the 4-byte IPv4 address and sets the appropriate
// offset, header type, and command.
//
// Parameters:
//
//	ip - The IPv4 source address to set (net.IP).
func (p *TcPedit) SetIPv4Src(ip net.IP) {
	// Convert the first 4 bytes of the IPv4 address to uint32.
	u32 := NativeEndian.Uint32(ip[:4])

	// Create a new TcPeditKey and TcPeditKeyEx for the IPv4 source address.
	tKey := TcPeditKey{}
	tKeyEx := TcPeditKeyEx{}

	// Set the value for the 4-byte IPv4 address.
	tKey.Val = u32

	// Set the offset to 12, which is the location of the source address in the IPv4 header.
	tKey.Off = 12

	// Set the extended key type to IPv4 and the command to "Set".
	tKeyEx.HeaderType = TcaPeditKeyExHdrTypeIP4
	tKeyEx.Cmd = PeditCmdSet

	// Append the new key and extended key to the TcPedit struct.
	p.Keys = append(p.Keys, tKey)
	p.KeysEx = append(p.KeysEx, tKeyEx)

	// Increment the number of keys in the TcPedit selection.
	p.Sel.NKeys++
}
