//go:build linux

package namespace

import (
	"golang.org/x/sys/unix"
)

// NsTypes - Define an enumerated set of values for setns().
//
// See https://man7.org/linux/man-pages/man2/setns.2.html
type NsTypes int

const (
	// JoinAnyNamespace - Allow any type of namespace to be joined.
	JoinAnyNamespace NsTypes = 0

	//CloneNewcgroup - (since Linux 4.6)  - fd must refer to a cgroup namespace.
	CloneNewcgroup NsTypes = unix.CLONE_NEWCGROUP

	//CloneNewipc -  (since Linux 3.0)     - fd must refer to an IPC namespace.
	CloneNewipc NsTypes = unix.CLONE_NEWIPC

	//CloneNewnet - (since Linux 3.0)     - fd must refer to a network namespace.
	CloneNewnet NsTypes = unix.CLONE_NEWNET

	//CloneNewns -  (since Linux 3.8)      - fd must refer to a mount namespace.
	CloneNewns NsTypes = unix.CLONE_NEWNS

	//CloneNewpid - (since Linux 3.8)     - fd must refer to a descendant PID namespace.
	CloneNewpid NsTypes = unix.CLONE_NEWPID

	//CloneNewtime - (since Linux 5.8)    - fd must refer to a time namespace.
	CloneNewtime NsTypes = unix.CLONE_NEWTIME

	//CloneNewuser - (since Linux 3.8)    - fd must refer to a user namespace.
	CloneNewuser NsTypes = unix.CLONE_NEWUSER

	//CloneNewuts - (since Linux 3.0)     - fd must refer to a UTS namespace.
	CloneNewuts NsTypes = unix.CLONE_NEWUTS
)
