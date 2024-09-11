package namespace

import (
	"golang.org/x/sys/unix"
)

// Set - set current network namespace to the namespace represented by NamespaceHandle.
func Set(ns Handle) error {
	return unix.Setns(int(ns), unix.CLONE_NEWNET)
}
