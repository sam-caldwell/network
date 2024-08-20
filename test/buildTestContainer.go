package test

import (
	"log"
	"os/exec"
)

// BuildTestContainer - Build the docker container used to test our solution.
func BuildTestContainer(imageName string) {
	var output []byte
	var err error
	commandStr := []string{
		"build", "--tag", imageName, ".",
	}
	output, err = exec.Command("docker", commandStr...).Output()
	log.Printf("Container build output:\n===\n%s\n===\n", string(output))
	if err != nil {
		log.Print(err)
	}
}
