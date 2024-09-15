package core

import (
	"golang.org/x/sys/unix"
)

// GetIpsetFlags - return the int version of the IpSet Flags
func GetIpsetFlags(cmd IpSetCmdEnum) int {

	switch cmd {

	case IpsetCmdCreate:
		return NlmFRequest | unix.NLM_F_ACK | unix.NLM_F_CREATE

	case IpsetCmdDestroy,
		IpsetCmdFlush,
		IpsetCmdRename,
		IpsetCmdSwap,
		IpsetCmdTest:
		return NlmFRequest | unix.NLM_F_ACK

	case IpsetCmdList,
		IpsetCmdSave:
		return NlmFRequest | unix.NLM_F_ACK | unix.NLM_F_ROOT | unix.NLM_F_MATCH | unix.NLM_F_DUMP

	case IpsetCmdAdd,
		IpsetCmdDel:
		return NlmFRequest | unix.NLM_F_ACK

	case IpsetCmdHeader,
		IpsetCmdType,
		IpsetCmdProtocol:
		return NlmFRequest

	default:
		return 0

	}

}
