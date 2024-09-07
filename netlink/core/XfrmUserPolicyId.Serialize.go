package core

import "unsafe"

func (msg *XfrmUserPolicyId) Serialize() []byte {
	return (*(*[SizeofXfrmUserPolicyId]byte)(unsafe.Pointer(msg)))[:]
}
