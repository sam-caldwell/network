package core

import (
	"errors"
)

// DeserializeTcPolice - Safely converts a byte slice into a TcPolice struct.
// This method ensures that the length of the input byte slice matches the expected size of TcPolice.
//
// If the length of the input byte slice is less than the size of TcPolice, it returns an error.
func DeserializeTcPolice(b []byte) (*TcPolice, error) {
	if len(b) < SizeOfTcPolice {
		return nil, errors.New("DeserializeTcPolice: byte slice too short")
	}

	msg := &TcPolice{}

	// Copy the content of the byte slice into the TcPolice struct.
	// Manually copy each field to avoid the risks of using unsafe.Pointer.
	msg.Index = NativeEndian.Uint32(b[0:4])
	msg.Action = int32(NativeEndian.Uint32(b[4:8]))
	msg.Limit = NativeEndian.Uint32(b[8:12])
	msg.Burst = NativeEndian.Uint32(b[12:16])
	msg.Mtu = NativeEndian.Uint32(b[16:20])
	msg.Rate = *DeserializeTcRateSpec(b[20 : 20+SizeOfTcRateSpec])
	msg.PeakRate = *DeserializeTcRateSpec(b[20+SizeOfTcRateSpec : 20+2*SizeOfTcRateSpec])
	msg.Refcnt = int32(NativeEndian.Uint32(b[20+2*SizeOfTcRateSpec : 20+2*SizeOfTcRateSpec+4]))
	msg.Bindcnt = int32(NativeEndian.Uint32(b[20+2*SizeOfTcRateSpec+4 : 20+2*SizeOfTcRateSpec+8]))
	msg.Capab = NativeEndian.Uint32(b[20+2*SizeOfTcRateSpec+8 : 20+2*SizeOfTcRateSpec+12])

	return msg, nil
}
