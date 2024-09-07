package core

import (
	"bytes"
	"encoding/binary"
	"errors"
)

// DeserializeXfrmUserPolicyInfo safely deserializes a byte slice into an XfrmUserPolicyInfo structure.
// It uses DeserializeXfrmSelector, DeserializeXfrmLifetimeCfg, and DeserializeXfrmLifetimeCur for respective fields.
// Returns an error if the byte slice is too small or deserialization fails.
func DeserializeXfrmUserPolicyInfo(b []byte) (*XfrmUserPolicyInfo, error) {
	// Check if the byte slice is large enough to hold the mandatory fields of XfrmUserPolicyInfo.
	if len(b) < SizeOfXfrmUserPolicyInfo {
		return nil, errors.New("byte slice too small to deserialize XfrmUserPolicyInfo")
	}

	// Create a new XfrmUserPolicyInfo struct.
	info := XfrmUserPolicyInfo{}
	reader := bytes.NewReader(b)

	// Deserialize the Sel (XfrmSelector) using DeserializeXfrmSelector
	sel, err := DeserializeXfrmSelector(b)
	if err != nil {
		return nil, err
	}
	info.Sel = *sel
	offset := SizeOfXfrmSelector // Move the offset after the selector

	// Deserialize the Lft (XfrmLifetimeCfg) using DeserializeXfrmLifetimeCfg
	lft, err := DeserializeXfrmLifetimeCfg(b[offset:])
	if err != nil {
		return nil, err
	}
	info.Lft = *lft
	offset += SizeOfXfrmLifetimeCfg

	// Deserialize the Curlft (XfrmLifetimeCur) using DeserializeXfrmLifetimeCur
	curlft, err := DeserializeXfrmLifetimeCur(b[offset:])
	if err != nil {
		return nil, err
	}
	info.Curlft = *curlft
	offset += SizeOfXfrmLifetimeCur

	// Deserialize the remaining fields directly using binary.Read.

	// Deserialize Priority
	if err := binary.Read(reader, binary.BigEndian, &info.Priority); err != nil {
		return nil, err
	}

	// Deserialize Index
	if err := binary.Read(reader, binary.BigEndian, &info.Index); err != nil {
		return nil, err
	}

	// Deserialize Dir
	if err := binary.Read(reader, binary.BigEndian, &info.Dir); err != nil {
		return nil, err
	}

	// Deserialize Action
	if err := binary.Read(reader, binary.BigEndian, &info.Action); err != nil {
		return nil, err
	}

	// Deserialize Flags
	if err := binary.Read(reader, binary.BigEndian, &info.Flags); err != nil {
		return nil, err
	}

	// Deserialize Share
	if err := binary.Read(reader, binary.BigEndian, &info.Share); err != nil {
		return nil, err
	}

	// Deserialize the Pad (padding)
	if err := binary.Read(reader, binary.BigEndian, &info.Pad); err != nil {
		return nil, err
	}

	// Return the deserialized XfrmUserPolicyInfo.
	return &info, nil
}
