//go:build linux

package core

import (
	"fmt"
	"syscall"
)

// DeserializeVfVlanList - Given []byte deserialize to a list of VfVlanInfo objects by reference.
func DeserializeVfVlanList(b []byte) ([]*VfVlanInfo, error) {

	var (
		err            error
		vfVlanInfoList []*VfVlanInfo
		attrs          []syscall.NetlinkRouteAttr
	)

	if attrs, err = ParseRouteAttr(b); err != nil {
		return nil, err
	}

	for _, element := range attrs {
		if element.Attr.Type == IFLA_VF_VLAN_INFO {
			vfVlanInfoList = append(vfVlanInfoList, DeserializeVfVlanInfo(element.Value))
		}
	}

	if len(vfVlanInfoList) == 0 {
		return nil, fmt.Errorf("VF vlan list is defined but no vf vlan info elements were found")
	}

	return vfVlanInfoList, nil
}
