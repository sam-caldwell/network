package core

import "unsafe"

const (
	BpfNetlinkPolicyEnumSize = unsafe.Sizeof(BpfNetlinkPolicyEnum(0))
)
