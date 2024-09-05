package core

import (
	"bytes"
	"encoding/binary"
)

// Serialize - Converts the TcConnmark structure into a byte slice.
// This method uses binary encoding to safely serialize the structure fields.
// It avoids using unsafe.Pointer by manually converting each field to bytes.
//
// Returns a serialized byte slice.
func (msg *TcConnmark) Serialize() []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, msg.Index)   // TcGen: Index
	binary.Write(buf, binary.LittleEndian, msg.Capab)   // TcGen: Capab
	binary.Write(buf, binary.LittleEndian, msg.Action)  // TcGen: Action
	binary.Write(buf, binary.LittleEndian, msg.Refcnt)  // TcGen: Refcnt
	binary.Write(buf, binary.LittleEndian, msg.Bindcnt) // TcGen: Bindcnt
	binary.Write(buf, binary.LittleEndian, msg.Zone)    // TcConnmark: Zone

	return buf.Bytes()
}
