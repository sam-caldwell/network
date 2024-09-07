package core

import (
	"bytes"
	"encoding/binary"
)

// Serialize safely serializes the XfrmUserSaInfo structure into a byte slice.
// It returns the byte slice that represents the serialized structure, or an error if serialization fails.
func (msg *XfrmUserSaInfo) Serialize() ([]byte, error) {
	// Create a buffer to hold the serialized data.
	buf := new(bytes.Buffer)

	// Serialize the XfrmSelector (Sel)
	selBytes, err := msg.Sel.Serialize()
	if err != nil {
		return nil, err
	}
	if _, err := buf.Write(selBytes); err != nil {
		return nil, err
	}

	// Serialize the XfrmId (Id)
	idBytes, err := msg.Id.Serialize()
	if err != nil {
		return nil, err
	}
	if _, err := buf.Write(idBytes); err != nil {
		return nil, err
	}

	// Serialize the XfrmAddress (Saddr)
	data, err := msg.Saddr.Serialize()
	if err != nil {
		return nil, err
	}
	if _, err := buf.Write(data); err != nil {
		return nil, err
	}

	// Serialize the XfrmLifetimeCfg (Lft)
	data, err = msg.Lft.Serialize()
	if err != nil {
		return nil, err
	}
	lftBytes, err := msg.Lft.Serialize()
	if err != nil {
		return nil, err
	}
	if _, err := buf.Write(lftBytes); err != nil {
		return nil, err
	}

	// Serialize the XfrmLifetimeCur (Curlft)
	curlftBytes, err := msg.Curlft.Serialize()
	if err != nil {
		return nil, err
	}
	if _, err := buf.Write(curlftBytes); err != nil {
		return nil, err
	}

	// Serialize the XfrmStats (Stats)
	statsBytes, err := msg.Stats.Serialize()
	if err != nil {
		return nil, err
	}
	if _, err := buf.Write(statsBytes); err != nil {
		return nil, err
	}

	// Serialize the Seq (uint32)
	if err := binary.Write(buf, binary.BigEndian, msg.Seq); err != nil {
		return nil, err
	}

	// Serialize the Reqid (uint32)
	if err := binary.Write(buf, binary.BigEndian, msg.Reqid); err != nil {
		return nil, err
	}

	// Serialize the Family (uint16)
	if err := binary.Write(buf, binary.BigEndian, msg.Family); err != nil {
		return nil, err
	}

	// Serialize the Mode (uint8)
	if err := binary.Write(buf, binary.BigEndian, msg.Mode); err != nil {
		return nil, err
	}

	// Serialize the ReplayWindow (uint8)
	if err := binary.Write(buf, binary.BigEndian, msg.ReplayWindow); err != nil {
		return nil, err
	}

	// Serialize the Flags (uint8)
	if err := binary.Write(buf, binary.BigEndian, msg.Flags); err != nil {
		return nil, err
	}

	// Serialize the Pad (7 bytes padding)
	if err := binary.Write(buf, binary.BigEndian, msg.Pad[:]); err != nil {
		return nil, err
	}

	// Return the serialized byte slice.
	return buf.Bytes(), nil
}
