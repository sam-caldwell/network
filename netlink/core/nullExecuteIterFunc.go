package core

// nullExecuteIterFunc - a noop function we will use when draining the queue
func nullExecuteIterFunc(msg []byte) bool {
	return true
}
