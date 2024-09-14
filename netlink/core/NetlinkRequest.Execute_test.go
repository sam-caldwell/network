package core

import (
	"testing"
)

func TestNetlinkRequestExecuteWithMock(t *testing.T) {

	t.Run("Test NetlinkRequest.Execute() method happy path", func(t *testing.T) {
		t.Skip("disabled")

		//t.Run("build container", func(t *testing.T) {
		//	if err := test.BuildTestContainer(testDockerImage); err != nil {
		//		t.Fatalf("container build error: %v", err)
		//	}
		//})
		//
		//t.Run("run test container", func(t *testing.T) {
		//	if err := test.RunContainer(testDockerImage, "TestNetlinkRequestExecuteFuncHappy"); err != nil {
		//		t.Fatalf("test container failed: %v", err)
		//	}
		//})

	})

	t.Run("Test NetlinkRequest.Execute() method error scenario", func(t *testing.T) {
		t.Skip("disabled")

		//t.Run("build container", func(t *testing.T) {
		//	if err := test.BuildTestContainer(testDockerImage); err != nil {
		//		t.Fatalf("container build error: %v", err)
		//	}
		//})
		//
		//t.Run("run test container", func(t *testing.T) {
		//	err := test.RunContainer(testDockerImage, "TestNetlinkRequestExecuteFuncErrorScenario")
		//	if err != nil {
		//		t.Fatalf("test container failed: %v", err)
		//	}
		//})

	})
}
