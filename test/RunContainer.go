package test

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

// RunContainer - Run a given test program inside a docker container
func RunContainer(imageName, testName string) error {
	defer func() { _ = DeleteContainer(testName) }()
	const command = "docker"
	args := []string{
		"run",
		"--name", testName,
		//"--cap-drop=ALL",
		"--cap-add=NET_ADMIN",
		"--cap-add=NET_RAW",
		"--cap-add=NET_BIND_SERVICE",
		"--cap-add=SYS_ADMIN",
		"--cap-add=DAC_OVERRIDE",
		"--cap-add=SYS_CHROOT",
		"--cap-add=MKNOD",
		"--cap-add=FOWNER",
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
