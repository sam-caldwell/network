//go:build linux

package core

import (
	"errors"
)

// DeserializeBridgeVlanInfo - Safely deserialize the BridgeVlanInfo object
func DeserializeBridgeVlanInfo(b []byte) (*BridgeVlanInfo, error) {

	if len(b) < SizeofBridgeVlanInfo {
		return nil, errors.New("input byte slice is too short")
	}

	return &BridgeVlanInfo{
		Flags: NativeEndian.Uint16(b[0:2]),
		Vid:   NativeEndian.Uint16(b[2:4]),
	}, nil

}
