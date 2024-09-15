package core

import "unsafe"

const (
	BpfProgPolicySize = unsafe.Sizeof(BpfProgPolicy(0))
)
