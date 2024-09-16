package core

import (
	"errors"
)

// Deserialize deserializes the byte slice `data` into a ConnectorMessageOperation structure.
//
// The `ConnectorMessageOperation` structure contains fields for message identification, sequence numbers, flags,
// and operations that can be deserialized from a byte slice. The function expects the byte slice
// to be at least the size of `ConnectorMessageOperation`. If the data slice is too short, it returns an error.
//
// The deserialization assumes that the data is encoded in the system's native byte order, which is
// handled using the `NativeEndian` utility.
//
// Parameters:
// - data: A byte slice containing the serialized `ConnectorMessageOperation` data.
//
// Returns:
// - An error if the data slice is too short to contain the required fields; otherwise nil.
//
// Example:
// ```go
// var msg ConnectorMessageOperation
// err := msg.Deserialize(data)
//
//	if err != nil {
//	    fmt.Println("Error deserializing:", err)
//	}
//
// ```
func (msg *ConnectorMessageOperation) Deserialize(data []byte) error {
	// Ensure that the byte slice is large enough to hold a ConnectorMessageOperation.
	if len(data) < ConnectorMessageOperationSize {
		return errors.New(ErrInputTooShort)
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
	msg.Operation = NativeEndian.Uint32(data[20:24]) // Extract 4 bytes for the operation code

	// Return nil if deserialization is successful.
	return nil
}

// DeserializeCnMsgOp - Deserialize []byte into ConnectorMessageOperation
func DeserializeCnMsgOp(b []byte) (*ConnectorMessageOperation, error) {
	if err := checkInputSize(b, ConnectorMessageOperationSize, ConnectorMessageOperationSize); err != nil {
		return nil, err
	}
	var o ConnectorMessageOperation
	if err := o.Deserialize(b); err != nil {
		return nil, err
	}
	return &o, nil
}
