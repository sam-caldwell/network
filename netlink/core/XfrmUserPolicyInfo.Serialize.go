package core

import (
	"bytes"
	"encoding/binary"
)

// Serialize safely serializes the XfrmUserPolicyInfo structure into a byte slice.
// It returns the byte slice that represents the serialized structure, or an error if serialization fails.
func (msg *XfrmUserPolicyInfo) Serialize() ([]byte, error) {
	// Create a buffer to hold the serialized data.
	buf := new(bytes.Buffer)

	// Serialize the fields in the order they appear in the structure.

	// Serialize the Sel (XfrmSelector)
	if temp, err := msg.Sel.Serialize(); err != nil {
		return nil, err
	} else {
		if err := binary.Write(buf, binary.BigEndian, &temp); err != nil {
			return nil, err
		}
	}

	// Serialize the Lft (XfrmLifetimeCfg)
	if temp, err := msg.Lft.Serialize(); err != nil {
		return nil, err
	} else {
		if err := binary.Write(buf, binary.BigEndian, &temp); err != nil {
			return nil, err
		}
	}

	// Serialize the Curlft (XfrmLifetimeCur)
	if temp, err := msg.Curlft.Serialize(); err != nil {
		return nil, err
	} else {
		if err := binary.Write(buf, binary.BigEndian, &temp); err != nil {
			return nil, err
		}
	}

	// Serialize the Priority
	if err := binary.Write(buf, binary.BigEndian, msg.Priority); err != nil {
		return nil, err
	}

	// Serialize the Index
	if err := binary.Write(buf, binary.BigEndian, msg.Index); err != nil {
		return nil, err
	}

	// Serialize the Dir
	if err := binary.Write(buf, binary.BigEndian, msg.Dir); err != nil {
		return nil, err
	}

	// Serialize the Action
	if err := binary.Write(buf, binary.BigEndian, msg.Action); err != nil {
		return nil, err
	}

	// Serialize the Flags
	if err := binary.Write(buf, binary.BigEndian, msg.Flags); err != nil {
		return nil, err
	}

	// Serialize the Share
	if err := binary.Write(buf, binary.BigEndian, msg.Share); err != nil {
		return nil, err
	}

	// Serialize the Pad (padding)
	if err := binary.Write(buf, binary.BigEndian, &msg.Pad); err != nil {
		return nil, err
	}

	// Return the serialized byte slice.
	return buf.Bytes(), nil
}
