package core

import "testing"

func TestIpProtocol_Serialize(t *testing.T) {
	tests := []struct {
		name     string
		protocol IpProtocol
		expected []byte
	}{
		{"HOPOPT", IpProtoHOPOPT, []byte{byte(IpProtoHOPOPT)}},
		{"ICMP", IpProtoIcmp, []byte{byte(IpProtoIcmp)}},
		{"IGMP", IpProtoIGMP, []byte{byte(IpProtoIGMP)}},
		{"TCP", IpProtoTCP, []byte{byte(IpProtoTCP)}},
		{"UDP", IpProtoUDP, []byte{byte(IpProtoUDP)}},
		{"Unknown Protocol", IpProtocol(255), []byte{255}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := test.protocol.Serialize()
			if err != nil {
				t.Errorf("Expected no error, but got %v", err)
			}
			if len(result) != 1 || result[0] != test.expected[0] {
				t.Errorf("Expected %v but got %v for %s", test.expected, result, test.name)
			}
		})
	}
}
