package core

import (
	"errors"
)

const (
	disableSizeCheck = -1
)

// checkInputSize - Check the input size against
func checkInputSize[T []byte | string](input T, minSize, maxSize int) error {
	if []byte(input) == nil {
		return errors.New(ErrNilInput)
	}
	if minSize != disableSizeCheck && len(input) < minSize {
		return errors.New(ErrInputTooShort)
	}
	if maxSize != disableSizeCheck && len(input) > maxSize {
		return errors.New(ErrInputTooLarge)
	}
	return nil
}
