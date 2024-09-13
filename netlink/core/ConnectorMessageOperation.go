package core

import (
	"errors"
	"unsafe"
)

// ConnectorMessageOperation - represents connector message operation, combining message identification,
// sequence, acknowledgment, and operational flags with an additional operation-specific field.
//
// Note that this EXTENDS cn_msg found in the Linux source tree and adds an Operation field.
//
// https://github.com/torvalds/linux/blob/master/include/linux/connector.h,
// https://github.com/vishvananda/netlink/blob/main/nl/nl_linux.go
type ConnectorMessageOperation struct {
	ConnectorMessage
	// here we differ from the C header
	Operation uint32
}

// SizeOfConnectorMessageOperation - size of ConnectorMessageOperation struct
const SizeOfConnectorMessageOperation = int(unsafe.Sizeof(ConnectorMessageOperation{}))

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
	if len(data) < SizeOfConnectorMessageOperation {
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

// Len - return the size of the ConnectorMessageOperation struct
func (msg *ConnectorMessageOperation) Len() int {
	return SizeOfConnectorMessageOperation
}

// Serialize converts the ConnectorMessageOperation struct into a byte slice representation.
//
// This function serializes the fields of the ConnectorMessageOperation structure into a byte slice (`[]byte`),
// which can be transmitted over a network or stored in a binary format. The byte order
// is determined by the `NativeEndian`, which should reflect the system's native byte order
// (e.g., little-endian or big-endian). This method ensures that each field in the struct
// is properly converted and placed into the byte slice.
//
// Return Value:
//   - A byte slice (`[]byte`) of length `SizeOfConnectorMessageOperation` containing the serialized representation
//     of the `ConnectorMessageOperation` struct.
//
// Example:
// ```go
// var msg ConnectorMessageOperation
// msg.ID.Idx = 1
// msg.ID.Val = 2
// msg.Seq = 3
// msg.Ack = 4
// msg.Length = 5
// msg.Flags = 6
// msg.Operation = 7
// serializedData, err := msg.Serialize()
//
//	if err != nil {
//	    log.Fatal(err)
//	}
//
// ```
//
// Reference:
//   - The `ConnectorMessageOperation` structure and the `NativeEndian` are commonly used in netlink communication and are derived from the
//     Linux kernel netlink interface. For more information, refer to the Linux source code:
//   - https://github.com/torvalds/linux/blob/master/include/uapi/linux/netlink.h,
//     https://github.com/vishvananda/netlink/blob/main/nl/nl_linux.go
func (msg *ConnectorMessageOperation) Serialize() ([]byte, error) {
	buf := make([]byte, SizeOfConnectorMessageOperation)

	// Serialize CbID (Idx and Val)
	NativeEndian.PutUint32(buf[0:], msg.ID.Idx)
	NativeEndian.PutUint32(buf[4:], msg.ID.Val)

	// Serialize Seq, Ack, Length, Flags, and Operation
	NativeEndian.PutUint32(buf[8:], msg.Seq)
	NativeEndian.PutUint32(buf[12:], msg.Ack)
	NativeEndian.PutUint16(buf[16:], msg.Length)
	NativeEndian.PutUint16(buf[18:], msg.Flags)
	NativeEndian.PutUint32(buf[20:], msg.Operation)

	return buf, nil
}

// DeserializeCnMsgOp - Deserialize []byte into ConnectorMessageOperation
func DeserializeCnMsgOp(b []byte) (*ConnectorMessageOperation, error) {
	if err := checkInputSize(b, SizeOfConnectorMessageOperation, SizeOfConnectorMessageOperation); err != nil {
		return nil, err
	}
	var o ConnectorMessageOperation
	if err := o.Deserialize(b); err != nil {
		return nil, err
	}
	return &o, nil
}
