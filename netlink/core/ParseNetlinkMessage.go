//go:build linux

package core

import (
	"golang.org/x/sys/unix"
)

// ParseNetlinkMessage - deserialize []byte into a list of NetlinkMessage objects.
func ParseNetlinkMessage(b []byte) ([]NetlinkMessage, error) {

	var messages []NetlinkMessage

	for len(b) >= unix.NLMSG_HDRLEN {

		h, dataBuffer, dataLength, err := DeserializeNetlinkMessageHeaderAndData(b)

		if err != nil {
			return nil, err
		}

		m := NetlinkMessage{Header: *h, Data: dataBuffer[:int(h.Len)-unix.NLMSG_HDRLEN]}

		messages = append(messages, m)

		b = b[dataLength:]

	}

	return messages, nil

}
