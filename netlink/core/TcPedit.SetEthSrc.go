package core

import "net"

// SetEthSrc sets the Ethernet source MAC address for the pedit action.
// This function modifies packet headers by creating pedit keys and extended pedit keys
// for the Ethernet source address (MAC). It splits the MAC address into two parts:
// the first 2 bytes as uint16 and the remaining 4 bytes as uint32.
//
// Parameters:
//
//	mac - The Ethernet source MAC address to set (net.HardwareAddr).
func (p *TcPedit) SetEthSrc(mac net.HardwareAddr) {
	// Convert the first 2 bytes of the MAC address to uint16.
	u16 := NativeEndian.Uint16(mac)

	// Convert the last 4 bytes of the MAC address to uint32.
	u32 := NativeEndian.Uint32(mac[2:])

	// Create a new TcPeditKey and TcPeditKeyEx for the first 2 bytes.
	tKey := TcPeditKey{}
	tKeyEx := TcPeditKeyEx{}

	// Set the value for the first 2 bytes of the MAC address, shifted left by 16 bits.
	tKey.Val = uint32(u16) << 16

	// Apply a mask to only modify the first 2 bytes (0x0000ffff).
	tKey.Mask = 0x0000ffff

	// Set the offset to 4 (starting at byte 5 in the Ethernet header).
	tKey.Off = 4

	// Set the extended key type to Ethernet and command to "Set".
	tKeyEx.HeaderType = TcaPeditKeyExHdrTypeEth
	tKeyEx.Cmd = PeditCmdSet

	// Append the new key and extended key to the TcPedit struct.
	p.Keys = append(p.Keys, tKey)
	p.KeysEx = append(p.KeysEx, tKeyEx)

	// Increment the number of keys in the TcPedit selection.
	p.Sel.NKeys++

	// Create a new TcPeditKey and TcPeditKeyEx for the last 4 bytes.
	tKey = TcPeditKey{}
	tKeyEx = TcPeditKeyEx{}

	// Set the value for the last 4 bytes of the MAC address.
	tKey.Val = u32

	// Apply a zero mask (no specific masking).
	tKey.Mask = 0

	// Set the offset to 8 (starting at byte 9 in the Ethernet header).
	tKey.Off = 8

	// Set the extended key type to Ethernet and command to "Set".
	tKeyEx.HeaderType = TcaPeditKeyExHdrTypeEth
	tKeyEx.Cmd = PeditCmdSet

	// Append the new key and extended key to the TcPedit struct.
	p.Keys = append(p.Keys, tKey)
	p.KeysEx = append(p.KeysEx, tKeyEx)

	// Increment the number of keys in the TcPedit selection.
	p.Sel.NKeys++
}
