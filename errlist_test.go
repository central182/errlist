package errlist_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/central182/errlist"
)

var (
	errFoo = errors.New("foo")
	errBar = errors.New("bar")
	errBaz = errors.New("baz")
)

type errorWrapper struct {
	err error
}

func (e errorWrapper) Error() string {
	return fmt.Sprintf("in wrapper: %s", e.err)
}

func (e errorWrapper) Unwrap() error {
	return e.err
}

func TestErrorList(t *testing.T) {
	t.Run("Multiple errors can be chained together", func(t *testing.T) {
		err := errlist.New(errFoo, fmt.Errorf("wrapped: %w", errBar), errorWrapper{err: errBaz})
		assert.Error(t, err)
		assert.EqualError(t, err, "foo: wrapped: bar: in wrapper: baz")
		assert.True(t, errors.Is(err, errFoo))
		assert.True(t, errors.Is(err, errBar))
		assert.True(t, errors.Is(err, errBaz))
	})

	t.Run("A single error can be put into the list as well", func(t *testing.T) {
		err := errlist.New(errFoo)
		assert.Error(t, err)
		assert.EqualError(t, err, "foo")
		assert.True(t, errors.Is(err, errFoo))
		assert.False(t, errors.Is(err, errBar))
		assert.False(t, errors.Is(err, errBaz))
	})

	t.Run("Nil is returned if no error is provided", func(t *testing.T) {
		err := errlist.New()
		assert.NoError(t, err)
		assert.False(t, errors.Is(err, errFoo))
		assert.False(t, errors.Is(err, errBar))
		assert.False(t, errors.Is(err, errBaz))
	})
}
