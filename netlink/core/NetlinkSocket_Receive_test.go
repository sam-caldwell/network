package core

import (
	"github.com/sam-caldwell/network/test"
	"testing"
)

func TestNetlinkSocket_Receive(t *testing.T) {

	const testDockerImage = "network-test:latest"

	t.Run("Test NetlinkSocket.Receive() method happy path", func(t *testing.T) {

		t.Run("build container", func(t *testing.T) {
			if err := test.BuildTestContainer(testDockerImage); err != nil {
				t.Fatalf("container build error: %v", err)
			}
		})

		t.Run("run test container", func(t *testing.T) {
			if err := test.RunContainer(testDockerImage, "TestNetlinkSocketReceive"); err != nil {
				t.Fatalf("test container failed: %v", err)
			}
		})

	})

}
