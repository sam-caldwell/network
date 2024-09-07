package core

import (
	"bytes"
	"encoding/binary"
	"errors"
)

// DeserializeXfrmUsersaInfo safely deserializes a byte slice into an XfrmUsersaInfo structure.
// It uses appropriate deserialization functions for fields like XfrmSelector, XfrmId, XfrmAddress, etc.
func DeserializeXfrmUsersaInfo(b []byte) (*XfrmUsersaInfo, error) {
	if len(b) < SizeofXfrmUsersaInfo {
		return nil, errors.New("byte slice too small to deserialize XfrmUsersaInfo")
	}

	// Create a reader for the byte slice
	reader := bytes.NewReader(b)

	// Create a new XfrmUsersaInfo struct
	msg := &XfrmUsersaInfo{}

	// Deserialize XfrmSelector (Sel)
	sel, err := DeserializeXfrmSelector(b)
	if err != nil {
		return nil, err
	}
	msg.Sel = *sel
	offset := SizeOfXfrmSelector

	// Deserialize XfrmId (Id)
	id, err := DeserializeXfrmId(b[offset:])
	if err != nil {
		return nil, err
	}
	msg.Id = *id
	offset += SizeofXfrmId

	// Deserialize XfrmAddress (Saddr)
	saddr, err := DeserializeXfrmAddress(b[offset:])
	if err != nil {
		return nil, err
	}
	msg.Saddr = *saddr
	offset += SizeofXfrmAddress

	// Deserialize XfrmLifetimeCfg (Lft)
	lft, err := DeserializeXfrmLifetimeCfg(b[offset:])
	if err != nil {
		return nil, err
	}
	msg.Lft = *lft
	offset += SizeofXfrmLifetimeCfg

	// Deserialize XfrmLifetimeCur (Curlft)
	curlft, err := DeserializeXfrmLifetimeCur(b[offset:])
	if err != nil {
		return nil, err
	}
	msg.Curlft = *curlft
	offset += SizeofXfrmLifetimeCur

	// Deserialize XfrmStats (Stats)
	stats, err := DeserializeXfrmStats(b[offset:])
	if err != nil {
		return nil, err
	}
	msg.Stats = *stats
	offset += SizeofXfrmStats

	// Deserialize Seq (uint32)
	if err := binary.Read(reader, binary.BigEndian, &msg.Seq); err != nil {
		return nil, err
	}

	// Deserialize Reqid (uint32)
	if err := binary.Read(reader, binary.BigEndian, &msg.Reqid); err != nil {
		return nil, err
	}

	// Deserialize Family (uint16)
	if err := binary.Read(reader, binary.BigEndian, &msg.Family); err != nil {
		return nil, err
	}

	// Deserialize Mode (uint8)
	if err := binary.Read(reader, binary.BigEndian, &msg.Mode); err != nil {
		return nil, err
	}

	// Deserialize ReplayWindow (uint8)
	if err := binary.Read(reader, binary.BigEndian, &msg.ReplayWindow); err != nil {
		return nil, err
	}

	// Deserialize Flags (uint8)
	if err := binary.Read(reader, binary.BigEndian, &msg.Flags); err != nil {
		return nil, err
	}

	// Deserialize Pad (7 bytes padding)
	if err := binary.Read(reader, binary.BigEndian, &msg.Pad); err != nil {
		return nil, err
	}

	// Return the deserialized XfrmUsersaInfo structure
	return msg, nil
}
