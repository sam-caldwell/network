package core

import (
	"errors"
	"fmt"
	"testing"
)

func TestCheckInputSize(t *testing.T) {
	type TestData struct {
		s       any
		err     error
		minSize int
		maxSize int
	}
	testData := []TestData{
		{"", nil, 0, 2},
		{"", errors.New(ErrInputTooShort), 1, 2},
		{"test", errors.New(ErrInputTooLarge), 1, 2},
		{[]byte("test"), errors.New(ErrInputTooLarge), 1, 2},
		{[]byte(""), nil, 0, 2},
		{[]byte(""), nil, disableSizeCheck, 2},
		{[]byte(""), errors.New(ErrInputTooShort), 1, 2},
		{[]byte(nil), errors.New(ErrNilInput), 10, 12},
		{[]byte(nil), errors.New(ErrNilInput), disableSizeCheck, 12},
		{[]byte(nil), errors.New(ErrNilInput), 12, disableSizeCheck},
		{[]byte(nil), errors.New(ErrNilInput), disableSizeCheck, disableSizeCheck},
		{[]byte("test"), nil, disableSizeCheck, disableSizeCheck},
	}

	for lineNo, test := range testData {

		testName := fmt.Sprintf("test #%d %s (%d/%d)", lineNo, test.s, test.minSize, test.maxSize)

		t.Run(testName, func(t *testing.T) {

			var err error

			switch test.s.(type) {
			case string:
				s := test.s.(string)
				err = checkInputSize(s, test.minSize, test.maxSize)
			case []byte:
				s := test.s.([]byte)
				err = checkInputSize(s, test.minSize, test.maxSize)
			default:
				t.Fatal("test design does not support this type")
			}

			if test.s == nil {
				if err == nil {
					t.Fatalf("expected error not found. err: %s", test.err.Error())
				} else {
					if err.Error() != test.err.Error() {
						t.Fatalf("expected error (on nil): %v, got: %v", test.err, err)
					}
				}
			}

			if test.minSize > 0 && err != nil {
				if err.Error() != test.err.Error() {
					if err.Error() != test.err.Error() {
						t.Fatalf("Expected error %v, got %v", test.err.Error(), err)
					}
				}
			}
			if test.minSize <= 0 && err != nil {
				if err.Error() != test.err.Error() {
					t.Fatalf("Expected error %v, got %v", test.err.Error(), err)
				}
			}
		})
	}
}
