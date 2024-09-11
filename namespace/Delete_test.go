package namespace

import (
	"github.com/sam-caldwell/network/test"
	"testing"
)

func TestDeleteFunction(t *testing.T) {

	t.Run("GetNewSetDelete test", func(t *testing.T) {

		const testDockerImage = "network-test:latest"

		t.Run("build container", func(t *testing.T) {
			if err := test.BuildTestContainer(testDockerImage); err != nil {
				t.Fatal(err)
			}
		})

		t.Run("run test container", func(t *testing.T) {
			if err := test.RunContainer(testDockerImage, "TestCreateAndDeleteNamespace"); err != nil {
				t.Fatal(err)
			}
		})

	})

}
