package core

import (
	"bytes"
	"encoding/binary"
)

// Serialize outputs a serialized []byte from the NetlinkRequest struct.
func (req *NetlinkRequest) Serialize() (out []byte, err error) {
	// Calculate the total length of the netlink message.
	length := SizeOfNlMsgHdr
	for _, data := range req.Data {
		s, err := data.Serialize()
		if err != nil {
			return nil, err
		}
		length += len(s)
	}
	length += len(req.RawData)

	// Update the message length in the header.
	req.Len = uint32(length)

	// Create a buffer to hold the serialized data.
	buf := bytes.NewBuffer(make([]byte, 0, length))

	// Serialize the netlink message header.
	if err = binary.Write(buf, NativeEndian, req.Len); err != nil {
		return nil, err
	}
	if err = binary.Write(buf, NativeEndian, req.Type); err != nil {
		return nil, err
	}
	if err = binary.Write(buf, NativeEndian, req.Flags); err != nil {
		return nil, err
	}
	if err = binary.Write(buf, NativeEndian, req.Seq); err != nil {
		return nil, err
	}
	if err = binary.Write(buf, NativeEndian, req.Pid); err != nil {
		return nil, err
	}

	// Serialize the payload data.
	for _, data := range req.Data {
		s, err := data.Serialize()
		if err != nil {
			return nil, err
		}
		buf.Write(s)
	}

	// Add the raw data, if any.
	if len(req.RawData) > 0 {
		buf.Write(req.RawData)
	}

	return buf.Bytes(), nil
}
