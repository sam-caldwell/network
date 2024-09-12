//go:build linux

package namespace

import (
	"golang.org/x/sys/unix"
)

// Types - Define an enumerated set of values for setns().
//
// See https://man7.org/linux/man-pages/man2/setns.2.html
type Types int

const (
	// JoinAnyNamespace - Allow any type of namespace to be joined.
	JoinAnyNamespace Types = 0

	//CloneNewcgroup - (since Linux 4.6)  - fd must refer to a cgroup namespace.
	CloneNewcgroup Types = unix.CLONE_NEWCGROUP

	//CloneNewipc -  (since Linux 3.0)     - fd must refer to an IPC namespace.
	CloneNewipc Types = unix.CLONE_NEWIPC

	//CloneNewnet - (since Linux 3.0)     - fd must refer to a network namespace.
	CloneNewnet Types = unix.CLONE_NEWNET

	//CloneNewns -  (since Linux 3.8)      - fd must refer to a mount namespace.
	CloneNewns Types = unix.CLONE_NEWNS

	//CloneNewpid - (since Linux 3.8)     - fd must refer to a descendant PID namespace.
	CloneNewpid Types = unix.CLONE_NEWPID

	//CloneNewtime - (since Linux 5.8)    - fd must refer to a time namespace.
	CloneNewtime Types = unix.CLONE_NEWTIME

	//CloneNewuser - (since Linux 3.8)    - fd must refer to a user namespace.
	CloneNewuser Types = unix.CLONE_NEWUSER

	//CloneNewuts - (since Linux 3.0)     - fd must refer to a UTS namespace.
	CloneNewuts Types = unix.CLONE_NEWUTS
)
