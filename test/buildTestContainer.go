package test

import (
	"log"
	"os/exec"
)

// BuildTestContainer - Build the docker container used to test our solution.
func BuildTestContainer(imageName string) error {
	log.Printf("Current Directory: %s", GetCurrentWorkingDirectory())

	output, err := exec.Command("docker", "build", "-f", "test/Dockerfile", "--tag", imageName, ".").
		CombinedOutput()

	log.Printf("Container build output:\n"+
		"===\n"+
		"%s\n"+
		"===\n",
		output)

	return err
}
