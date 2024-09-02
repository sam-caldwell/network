//go:build linux

package core

const (
	// ReceiveBufferSize - RECEIVE_BUFFER_SIZE - Arbitrary set value (greater than default 4k) to allow receiving
	// from kernel more verbose messages e.g. for statistics, tc rules or filters, or other more memory requiring data.
	ReceiveBufferSize = 65536
)
