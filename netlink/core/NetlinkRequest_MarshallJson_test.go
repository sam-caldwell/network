package core

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestNetlinkRequest_MarshalJSON(t *testing.T) {
	// Sample NetlinkRequest for testing
	req := &NetlinkRequest{
		NetlinkMessageHeader: NetlinkMessageHeader{
			Len:   32,
			Type:  1,
			Flags: 2,
			Seq:   12345,
			Pid:   67890,
		},
		Data:    []NetlinkRequestData{}, // Assuming NetlinkRequestData is defined elsewhere
		RawData: []byte{0xDE, 0xAD, 0xBE, 0xEF},
		Sockets: map[IpProtocol]*SocketHandle{}, // Assuming IpProtocol and SocketHandle are defined elsewhere
	}

	// Expected JSON result
	expectedJSON := `{
		"len":32,
		"type":1,
		"flags":2,
		"seq":12345,
		"pid":67890,
		"data":[],
		"rawData":"3q2+7w==",
		"sockets":{}
	}`

	// Marshal to JSON
	jsonData, err := req.MarshalJSON()
	if err != nil {
		t.Fatalf("MarshalJSON() returned an error: %v", err)
	}

	// Unmarshal the expected JSON string to map[string]interface{} for comparison
	var expectedMap map[string]interface{}
	if err := json.Unmarshal([]byte(expectedJSON), &expectedMap); err != nil {
		t.Fatalf("Failed to unmarshal expected JSON: %v", err)
	}

	// Unmarshal the actual JSON data to map[string]interface{} for comparison
	var actualMap map[string]interface{}
	if err := json.Unmarshal(jsonData, &actualMap); err != nil {
		t.Fatalf("Failed to unmarshal actual JSON: %v", err)
	}

	// Compare the expected and actual JSON data
	if !reflect.DeepEqual(expectedMap, actualMap) {
		t.Errorf("MarshalJSON() output mismatch. Expected:\n%s\nGot:\n%s", expectedJSON, string(jsonData))
	}
}
