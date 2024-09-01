package core

// DeserializeCnMsgOp - Deserialize []byte into CnMsgOp
func DeserializeCnMsgOp(b []byte) *CnMsgOp {
	var o CnMsgOp
	if err := o.Deserialize(b); err != nil {
		panic(err)
	}
	return &o
}
