package core

import (
	"errors"
)

// checkInputSize - Check the input size against
func checkInputSize[T []byte | string](input T, minSize, maxSize int) error {
	if []byte(input) == nil {
		return errors.New(ErrNilInput)
	}
	if minSize >= 0 && len(input) < minSize {
		return errors.New(ErrInputTooShort)
	}
	if maxSize >= 0 && len(input) > maxSize {
		return errors.New(ErrInputTooLarge)
	}
	return nil
}
