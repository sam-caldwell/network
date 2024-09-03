package core

import (
	"errors"
)

// Error - Return the error for a given IpSetError
func (e IpSetErrorEnum) Error() error {
	return errors.New(e.ErrorString())
}
