//go:build linux

package core

import (
	"fmt"
)

// DeserializeVfVlanList - Given []byte deserialize to a list of IfLaVfVlanInfoStruct objects by reference.
func DeserializeVfVlanList(b []byte) ([]*IfLaVfVlanInfoStruct, error) {

	var (
		err            error
		vfVlanInfoList []*IfLaVfVlanInfoStruct
		attrs          []NetlinkRouteAttr
	)

	if attrs, err = ParseRouteAttr(b); err != nil {
		return nil, err
	}

	for _, element := range attrs {
		if element.Attr.Type == IfLaVfVlanInfo {
			vfVlanInfoList = append(vfVlanInfoList, DeserializeVfVlanInfo(element.Value))
		}
	}

	if len(vfVlanInfoList) == 0 {
		return nil, fmt.Errorf("VF vlan list is defined but no vf vlan info elements found")
	}

	return vfVlanInfoList, nil
}
