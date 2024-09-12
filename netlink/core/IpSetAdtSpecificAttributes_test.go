package core

import (
	"testing"
)

func TestIpSetAdtSpecificAttributes(t *testing.T) {
	tests := []struct {
		name   string
		val    IpSetAdtSpecificAttributes
		expect IpSetAdtSpecificAttributes
	}{
		{"IpsetAttrEther", IpsetAttrEther, IpSetAdtSpecificAttributes(IpsetAttrCadtMax) + 1},
		{"IpsetAttrName", IpsetAttrName, IpsetAttrEther + 1},
		{"IpsetAttrNameref", IpsetAttrNameref, IpsetAttrEther + 2},
		{"IpsetAttrIp2", IpsetAttrIp2, IpsetAttrEther + 3},
		{"IpsetAttrCidr2", IpsetAttrCidr2, IpsetAttrEther + 4},
		{"IpsetAttrIp2To", IpsetAttrIp2To, IpsetAttrEther + 5},
		{"IpsetAttrIface", IpsetAttrIface, IpsetAttrEther + 6},
		{"IpsetAttrBytes", IpsetAttrBytes, IpsetAttrEther + 7},
		{"IpsetAttrPackets", IpsetAttrPackets, IpsetAttrEther + 8},
		{"IpsetAttrComment", IpsetAttrComment, IpsetAttrEther + 9},
		{"IpsetAttrSkbmark", IpsetAttrSkbmark, IpsetAttrEther + 10},
		{"IpsetAttrSkbprio", IpsetAttrSkbprio, IpsetAttrEther + 11},
		{"IpsetAttrSkbqueue", IpsetAttrSkbqueue, IpsetAttrEther + 12},
		{"IpsetAttrPad", IpsetAttrPad, IpsetAttrEther + 13},
		{"IpsetAttrIpaddrIpv4", IpsetAttrIpaddrIpv4, IpSetAdtSpecificAttributes(IpsetAttrUnspec + 1)},
		{"IpsetAttrIpaddrIpv6", IpsetAttrIpaddrIpv6, IpSetAdtSpecificAttributes(IpsetAttrUnspec + 2)},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.val != test.expect {
				t.Errorf("Expected %d but got %d for %s", test.expect, test.val, test.name)
			}
		})
	}
}
