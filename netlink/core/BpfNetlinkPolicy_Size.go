package core

import "unsafe"

const (
	BpfNetlinkPolicySize = unsafe.Sizeof(BpfNetlinkPolicy(0))
)
