package core

import (
	"github.com/sam-caldwell/network/test"
	"testing"
)

func TestNetlinkRequestExecuteIter(t *testing.T) {

	t.Skip("disabled")

	const testDockerImage = "network-test:latest"

	t.Run("Test NetlinkRequest.ExecuteIter() method happy path", func(t *testing.T) {

		t.Run("build container", func(t *testing.T) {
			if err := test.BuildTestContainer(testDockerImage); err != nil {
				t.Fatalf("container build error: %v", err)
			}
		})

		t.Run("run test container", func(t *testing.T) {
			if err := test.RunContainer(testDockerImage, "TestNetlinkRequestExecuteIterFuncHappy"); err != nil {
				t.Fatalf("test container failed: %v", err)
			}
		})

	})

	t.Run("Test NetlinkRequest.ExecuteIter() method error scenario", func(t *testing.T) {

		t.Run("build container", func(t *testing.T) {
			if err := test.BuildTestContainer(testDockerImage); err != nil {
				t.Fatalf("container build error: %v", err)
			}
		})

		t.Run("run test container", func(t *testing.T) {
			err := test.RunContainer(testDockerImage, "TestNetlinkRequestExecuteIterFuncErrorScenario")
			if err != nil {
				t.Fatalf("test container failed: %v", err)
			}
		})

	})
}
