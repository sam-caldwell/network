package core

import (
	"log"
)

// ParseAttributes - Parse the input byte slice into an Attribute struct and push to the channel.
func ParseAttributes(data []byte) <-chan Attribute {
	result := make(chan Attribute)

	go func() {
		i := 0
		for i+4 < len(data) {
			length := int(NativeEndian.Uint16(data[i : i+2]))
			attrType := NlaFlags(NativeEndian.Uint16(data[i+2 : i+4]))

			if length < 4 {
				log.Printf("attribute 0x%02x has invalid length of %d bytes", attrType, length)
				break
			}

			if len(data) < i+length {
				log.Printf("attribute 0x%02x of length %d is truncated, only %d bytes remaining", attrType, length, len(data)-i)
				break
			}

			result <- Attribute{
				Type:  attrType,
				Value: data[i+4 : i+length],
			}
			i += rtaAlignOf(length)
		}
		close(result)
	}()

	return result
}
