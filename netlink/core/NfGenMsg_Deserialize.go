package core

import "encoding/binary"

// DeserializeNfgenmsg - Deserialize NfGenMsg from []byte and return pointer to the object.
func DeserializeNfgenmsg(b []byte) (*NfGenMsg, error) {
	if err := checkInputSize(b, NfGenMsgSize, NfGenMsgSize); err != nil {
		return nil, err
	}
	nfgenmsg := &NfGenMsg{}

	// Deserialize the NfgenFamily field
	nfgenmsg.NfgenFamily = b[0]

	// Deserialize the Version field
	nfgenmsg.Version = b[1]

	// Deserialize the ResId field (big endian)
	nfgenmsg.ResId = binary.BigEndian.Uint16(b[2:4])

	return nfgenmsg, nil
}
