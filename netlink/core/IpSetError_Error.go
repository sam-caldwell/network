package core

import (
	"errors"
)

// Error - Return the error for a given IpSetError
func (e IpSetError) Error() error {
	return errors.New(e.ErrorString())
}
