package errlist_test

import (
	"errors"
	"fmt"

	"github.com/central182/errlist"
)

var (
	ErrFoo = errors.New("foo")
	ErrBar = errors.New("bar")
	ErrBaz = errors.New("baz")
)

type ErrorWrapper struct {
	Err error
}

func (e ErrorWrapper) Error() string {
	return fmt.Sprintf("in wrapper: %s", e.Err)
}

func (e ErrorWrapper) Unwrap() error {
	return e.Err
}

func Example() {
	err := errlist.New(ErrFoo, fmt.Errorf("wrapped: %w", ErrBar), ErrorWrapper{Err: ErrBaz})
	fmt.Println(err)
	fmt.Println(errors.Is(err, ErrFoo))
	fmt.Println(errors.Is(err, ErrBar))
	fmt.Println(errors.Is(err, ErrBaz))
}
