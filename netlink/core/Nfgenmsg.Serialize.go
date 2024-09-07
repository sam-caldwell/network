//go:build linux

package core

import (
	"encoding/binary"
)

// Serialize - Serialize Nfgenmsg to []byte
func (msg *Nfgenmsg) Serialize() ([]byte, error) {
	buf := make([]byte, SizeofNfgenmsg)

	// Serialize the NfgenFamily field
	buf[0] = msg.NfgenFamily

	// Serialize the Version field
	buf[1] = msg.Version

	// Serialize the ResId field (big endian)
	binary.BigEndian.PutUint16(buf[2:4], msg.ResId)

	return buf, nil
}
