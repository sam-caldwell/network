package core

import "encoding/binary"

// NewCnMsg - Create a new connector message.
func NewCnMsg(idx, val, op uint32) *ConnectorMessageOperation {
	var cm ConnectorMessageOperation

	cm.ID.Idx = idx
	cm.ID.Val = val

	cm.Ack = 0
	cm.Seq = 1
	cm.Length = uint16(binary.Size(op))
	cm.Operation = op

	return &cm
}
