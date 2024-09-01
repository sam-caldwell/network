package core

import (
	"encoding/binary"
	"errors"
)

// DeserializeNfgenmsg - Deserialize Nfgenmsg from []byte and return pointer to the object.
func DeserializeNfgenmsg(b []byte) (*Nfgenmsg, error) {
	if len(b) < SizeofNfgenmsg {
		return nil, errors.New("byte slice too short to deserialize Nfgenmsg")
	}

	nfgenmsg := &Nfgenmsg{}

	// Deserialize the NfgenFamily field
	nfgenmsg.NfgenFamily = b[0]

	// Deserialize the Version field
	nfgenmsg.Version = b[1]

	// Deserialize the ResId field (big endian)
	nfgenmsg.ResId = binary.BigEndian.Uint16(b[2:4])

	return nfgenmsg, nil
}
