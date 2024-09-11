Network/Namespace
=================

## Description

This Go package provides a robust interface for managing and manipulating Linux network namespaces. It wraps
around common system calls such as `unshare()` and `setns()` to create, delete, and interact with network
namespaces using file descriptors. The package supports namespace operations by process ID (PID), thread ID
(TID), or a named network namespace. It also tracks namespaces in-memory to avoid redundant file descriptor
creation, ensuring efficient and thread-safe namespace management.

Linux network namespaces allow different processes or threads to have their own isolated network configurations
(interfaces, routing, etc.). In Linux, a thread can be part of a different network namespace than its parent process.

While a lot of functionality in this package closely maps to functionality in Linux C sources, because this package is
written in pure Golang, and because of how Golang handles memory and threads/processes, some functionality has to
differ a bit.  For example: This package has an in-memory means of tracking namespace handles and their associated
file handles to both avoid expensive system calls, thread locking to avoid unintentional namespace switching, etc.

## Key features include:

- Creating new network namespaces using the `unshare()` system call.
- Binding named namespaces for persistent access.
- Handling namespaces by process and thread ID.
- Deleting namespaces from both in-memory tracking and the filesystem.
- Safely closing namespace handles and preventing further usage after closure.

## Dependencies

This package requires the following dependencies:

- **Go Version**: Go 1.23 or later
- **Packages**:
    - `golang.org/x/sys/unix`: Provides access to low-level operating system primitives such as
      `unshare()`, `setns()`, `fstat()`, and `mount()`.
    - `sync`: For safe concurrency control.
    - `path/filepath`, `fmt`, `os`: For working with file paths, error formatting, and file operations.

## Installation

You can install these dependencies by running:

```bash
go get golang.org/x/sys/unix
```

## References

This package interacts with Linux kernel functionalities that handle network namespaces. Below are some key references:

### Linux Kernel Source for Network Namespaces:

- [Network Namespace Management](https://github.com/torvalds/linux/blob/master/net/core/net_namespace.c)
- [Process Namespace Management](https://github.com/torvalds/linux/blob/master/kernel/pid.c)

### Man Pages for Relevant System Calls:

- [unshare(2)](https://man7.org/linux/man-pages/man2/unshare.2.html):
  Describes how to create new namespaces.
- [setns(2)](https://man7.org/linux/man-pages/man2/setns.2.html):
  Describes how to switch to an existing namespace.
- [stat(2)](https://man7.org/linux/man-pages/man2/stat.2.html):
  Used for file descriptor comparison based on device and inode.

## Usage

This section explains how to use the key features of the package.

### Create a New Network Namespace

To create a new network namespace, call the `NewHandle()` function:

```go
nsHandle, err := NewHandle()
if err != nil {
log.Fatalf("Failed to create a new network namespace: %v", err)
}
defer nsHandle.Close() // Remember to close the handle when done
```

You can also create a named namespace, which will be bind-mounted for persistence:

```go
nsHandle, err := NewWithName("my_namespace")
if err != nil {
log.Fatalf("Failed to create named namespace: %v", err)
}
defer nsHandle.Close()
```

### Switch to a Namespace

To switch to an existing network namespace, use Set() with a valid Handle:

```go
if err := Set(nsHandle); err != nil {
log.Fatalf("Failed to switch to namespace: %v", err)
}
```

### Delete a Named Namespace

To delete a named namespace from both memory and the filesystem, call Delete():

```go
err := Delete("my_namespace")
if err != nil {
log.Fatalf("Failed to delete named namespace: %v", err)
}
```

### Get a Namespace by Process or Thread ID

To retrieve the network namespace of a running process or thread, use:

```go
pid := os.Getpid()
tid := unix.Gettid()
nsHandle, err := GetFromThread(pid, tid)
if err != nil {
log.Fatalf("Failed to get namespace: %v", err)
}
defer nsHandle.Close()
```

### Handle Comparison and Management

You can compare two namespace handles to check if they point to the same namespace:

```go
if nsHandle.Equal(otherNsHandle) {
fmt.Println("Handles refer to the same network namespace")
}
```

To check if a handle is still open:

```go
if nsHandle.IsOpen() {
fmt.Println("Namespace handle is open")
}
```

To get a unique identifier string for a namespace:

```go
fmt.Println("Namespace Unique ID:", nsHandle.UniqueId())
```

## License

See [LICENSE.txt](../LICENSE.txt) for more information.






