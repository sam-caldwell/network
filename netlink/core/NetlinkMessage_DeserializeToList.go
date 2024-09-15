//go:build linux

package core

// NetlinkMessageHeaderLength - unix.NLMSG_HDRLEN -
const NetlinkMessageHeaderLength = 0x10

// DeserializeToList - deserialize []byte into a list of NetlinkMessage objects.
func DeserializeToList(b []byte) ([]NetlinkMessage, error) {

	var messages []NetlinkMessage

	for len(b) >= NetlinkMessageHeaderLength {

		header, dataBuffer, dataLength, err := DeserializeNetlinkMessage(b)

		if err != nil {
			return nil, err
		}

		m := NetlinkMessage{Header: *header, Data: dataBuffer[:int(header.Len)-NetlinkMessageHeaderLength]}

		messages = append(messages, m)

		b = b[dataLength:]

	}

	return messages, nil

}
