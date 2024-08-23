package test

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

// RunContainer - Run a given test program inside a docker container
func RunContainer(imageName, testName string) error {
	defer DeleteContainer(testName)
	const command = "docker"
	args := []string{
		"run",
		"--name", testName,
		"--cap-drop=ALL",
		"--cap-add=CAP_NET_ADMIN",
		"--cap-add=CAP_SYS_ADMIN",
		"--cap-add=CAP_DAC_OVERRIDE",
		"--cap-add=CAP_SYS_CHROOT",
		"--cap-add=CAP_MKNOD",
		"--cap-add=CAP_FOWNER",
		"--privileged",
		"-v", GetCurrentWorkingDirectory() + ":/opt",
		imageName,
		"/usr/local/go/bin/go", "run", fmt.Sprintf("examples/%s/main.go", testName),
	}
	log.Println(command, strings.Join(args, " "))
	output, err := exec.Command(command, args...).CombinedOutput()
	log.Println("container output:")
	log.Println(string(output))
	log.Println("-----------------")
	return err
}
