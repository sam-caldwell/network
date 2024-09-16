package core

// Serialize converts the ConnectorMessageOperation struct into a byte slice representation.
//
// This function serializes the fields of the ConnectorMessageOperation structure into a byte slice (`[]byte`),
// which can be transmitted over a network or stored in a binary format. The byte order
// is determined by the `NativeEndian`, which should reflect the system's native byte order
// (e.g., little-endian or big-endian). This method ensures that each field in the struct
// is properly converted and placed into the byte slice.
//
// Return Value:
//   - A byte slice (`[]byte`) of length `ConnectorMessageOperationSize` containing the serialized representation
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
	buf := make([]byte, ConnectorMessageOperationSize)

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
