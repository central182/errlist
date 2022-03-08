// Package errlist provides an easier way to wrap errors in Go:
// provide errlist.New with a list of errors and they will appear as if they were wrapping one another.
package errlist

import (
	"errors"
	"strings"
)

type errorList struct {
	errs []error
	text string
}

// New converts a list of errors into a chain of errors,
// where each one in the chain can be viewed as wrapping the error to its right.
//
// Calling errors.Is on the resulting error evaluates to true if
// either the target is one of the list items
// or the target is being wrapped in one of the list items.
func New(errs ...error) error {
	if len(errs) == 0 {
		return nil
	}

	texts := make([]string, len(errs))
	for i, err := range errs {
		texts[i] = err.Error()
	}

	return errorList{
		errs: errs,
		text: strings.Join(texts, ": "),
	}
}

func (e errorList) Error() string {
	return e.text
}

func (e errorList) Is(target error) bool {
	for _, err := range e.errs {
		if errors.Is(err, target) {
			return true
		}
	}
	return false
}
