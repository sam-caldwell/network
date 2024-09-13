package core

import (
	"testing"
)

func TestConnectorMessage_Size(t *testing.T) {
	const expectedSizeInBytes = 20
	if ConnectorMessageSize != expectedSizeInBytes {
		t.Fatal("Wrong size")
	}
}
