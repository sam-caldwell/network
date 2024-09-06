package core

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// Serialize converts the XfrmSelector struct into a byte slice safely
// using the encoding/binary package. This ensures correct byte order and portability.
func (msg *XfrmSelector) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	// Serialize Daddr (Destination IP address)
	if err := binary.Write(buf, binary.BigEndian, msg.Daddr); err != nil {
		return nil, fmt.Errorf("failed to serialize Daddr: %w", err)
	}

	// Serialize Saddr (Source IP address)
	if err := binary.Write(buf, binary.BigEndian, msg.Saddr); err != nil {
		return nil, fmt.Errorf("failed to serialize Saddr: %w", err)
	}

	// Serialize Dport (Destination port)
	if err := binary.Write(buf, binary.BigEndian, msg.Dport); err != nil {
		return nil, fmt.Errorf("failed to serialize Dport: %w", err)
	}

	// Serialize DportMask (Destination port mask)
	if err := binary.Write(buf, binary.BigEndian, msg.DportMask); err != nil {
		return nil, fmt.Errorf("failed to serialize DportMask: %w", err)
	}

	// Serialize Sport (Source port)
	if err := binary.Write(buf, binary.BigEndian, msg.Sport); err != nil {
		return nil, fmt.Errorf("failed to serialize Sport: %w", err)
	}

	// Serialize SportMask (Source port mask)
	if err := binary.Write(buf, binary.BigEndian, msg.SportMask); err != nil {
		return nil, fmt.Errorf("failed to serialize SportMask: %w", err)
	}

	// Serialize Family (Address family)
	if err := binary.Write(buf, binary.BigEndian, msg.Family); err != nil {
		return nil, fmt.Errorf("failed to serialize Family: %w", err)
	}

	// Serialize PrefixlenD (Destination prefix length)
	if err := binary.Write(buf, binary.BigEndian, msg.PrefixlenD); err != nil {
		return nil, fmt.Errorf("failed to serialize PrefixlenD: %w", err)
	}

	// Serialize PrefixlenS (Source prefix length)
	if err := binary.Write(buf, binary.BigEndian, msg.PrefixlenS); err != nil {
		return nil, fmt.Errorf("failed to serialize PrefixlenS: %w", err)
	}

	// Serialize Proto (Protocol identifier)
	if err := binary.Write(buf, binary.BigEndian, msg.Proto); err != nil {
		return nil, fmt.Errorf("failed to serialize Proto: %w", err)
	}

	// Serialize Pad (Padding for 8-byte alignment)
	if err := binary.Write(buf, binary.BigEndian, msg.Pad); err != nil {
		return nil, fmt.Errorf("failed to serialize Pad: %w", err)
	}

	// Serialize Ifindex (Interface index)
	if err := binary.Write(buf, binary.BigEndian, msg.Ifindex); err != nil {
		return nil, fmt.Errorf("failed to serialize Ifindex: %w", err)
	}

	// Serialize User (User ID)
	if err := binary.Write(buf, binary.BigEndian, msg.User); err != nil {
		return nil, fmt.Errorf("failed to serialize User: %w", err)
	}

	// Return the serialized byte slice
	return buf.Bytes(), nil
}
