package core

import (
	"errors"
)

// Deserialize deserializes the byte slice `data` into a CnMsgOp structure.
//
// The `CnMsgOp` structure contains fields for message identification, sequence numbers, flags,
// and operations that can be deserialized from a byte slice. The function expects the byte slice
// to be at least the size of `CnMsgOp`. If the data slice is too short, it returns an error.
//
// The deserialization assumes that the data is encoded in the system's native byte order, which is
// handled using the `NativeEndian` utility.
//
// Parameters:
// - data: A byte slice containing the serialized `CnMsgOp` data.
//
// Returns:
// - An error if the data slice is too short to contain the required fields; otherwise nil.
//
// Example:
// ```go
// var msg CnMsgOp
// err := msg.Deserialize(data)
//
//	if err != nil {
//	    fmt.Println("Error deserializing:", err)
//	}
//
// ```
func (msg *CnMsgOp) Deserialize(data []byte) error {
	// Ensure that the byte slice is large enough to hold a CnMsgOp.
	if len(data) < SizeOfCnMsgOp {
		return errors.New("data too short to deserialize CnMsgOp")
	}

	// Deserialize ID (Idx and Val)
	msg.ID.Idx = NativeEndian.Uint32(data[0:4]) // Extract the first 4 bytes for ID.Idx
	msg.ID.Val = NativeEndian.Uint32(data[4:8]) // Extract the next 4 bytes for ID.Val

	// Deserialize sequence and acknowledgment numbers
	msg.Seq = NativeEndian.Uint32(data[8:12])  // Extract 4 bytes for the sequence number
	msg.Ack = NativeEndian.Uint32(data[12:16]) // Extract 4 bytes for the acknowledgment number

	// Deserialize Length and Flags
	msg.Length = NativeEndian.Uint16(data[16:18]) // Extract 2 bytes for the message length
	msg.Flags = NativeEndian.Uint16(data[18:20])  // Extract 2 bytes for the message flags

	// Deserialize the operation code
	msg.Op = NativeEndian.Uint32(data[20:24]) // Extract 4 bytes for the operation code

	// Return nil if deserialization is successful.
	return nil
}
