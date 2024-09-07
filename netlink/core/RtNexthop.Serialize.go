package core

import (
	"bytes"
	"encoding/binary"
)

func (msg *RtNexthop) Serialize() ([]byte, error) {
	// Create a buffer to store the serialized data
	buf := new(bytes.Buffer)

	// Serialize the RtNexthop struct fields manually
	if err := binary.Write(buf, NativeEndian, msg.Len); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, NativeEndian, msg.Flags); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, NativeEndian, msg.Hops); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, binary.LittleEndian, msg.Ifindex); err != nil {
		return nil, err
	}

	// Serialize children (netlink attributes)
	if len(msg.Children) > 0 {
		for _, child := range msg.Children {
			childBuf, err := child.Serialize()
			if err != nil {
				return nil, err
			}
			buf.Write(childBuf)
		}
	}

	return buf.Bytes(), nil
}
