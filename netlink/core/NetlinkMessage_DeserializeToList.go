//go:build linux

package core

import (
	"errors"
)

// DeserializeToList - deserialize []byte into a list of NetlinkMessage objects.
func DeserializeToList(b []byte) ([]NetlinkMessage, error) {

	var (
		messages []NetlinkMessage
	)

	if len(b) < NetlinkMessageHeaderSize {
		return []NetlinkMessage{}, errors.New(ErrInputTooShort)
	}

	for len(b) >= NetlinkMessageHeaderSize {
		var (
			err        error
			header     *NlMsghdr
			dataBuffer []byte
			dataLength int
		)
		if header, dataBuffer, dataLength, err = DeserializeNetlinkMessage(b[dataLength:]); err != nil {
			return nil, err
		}
		m := NetlinkMessage{
			Header: *header,
			Data:   dataBuffer[:int(header.Len)-NetlinkMessageHeaderSize],
		}
		messages = append(messages, m)
		b = b[dataLength:]
	}
	return messages, nil
}
