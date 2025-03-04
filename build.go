package errx

import (
	"errors"
	"fmt"
)

type Builder[T comparable] struct {
	err error
	msg string

	from string
	code T

	stack []Frame
}

func build[T comparable]() Builder[T] {
	return Builder[T]{}
}

func (eb Builder[T]) clone() Builder[T] {
	return Builder[T]{
		from: eb.from,
		code: eb.code,
	}
}

func (eb Builder[T]) New(msg string) error {
	if msg != "" {
		ex := eb.clone()
		ex.msg = msg
		ex.stack = stack()
		return (*Error[T])(&ex)
	}
	return nil
}

func (eb Builder[T]) Errorf(format string, args ...interface{}) error {
	return eb.New(fmt.Sprintf(format, args...))
}

func Wrap(err error) error {
	return Code(Unspecified).Wrap(err)
}

func (eb Builder[T]) Wrap(err error) error {
	if err != nil {
		ex := eb.clone()
		ex.err = err
		ex.stack = stack()
		return (*Error[T])(&ex)
	}
	return nil
}

func (eb Builder[T]) Join(e ...error) error {
	return eb.clone().Wrap(errors.Join(e...))
}

func Recover(fallback error) error {
	if r := recover(); r != nil {
		err, ok := r.(error)
		if !ok {
			err = Errorf("%v", r)
		}
		return err
	}
	return fallback
}

func From[T comparable](from string) Builder[T] {
	return build[T]().From(from)
}

func (eb Builder[T]) From(from string) Builder[T] {
	nb := eb.clone()
	nb.from = from
	return nb
}

func Code[T comparable](code T) Builder[T] {
	return build[T]().Code(code)
}

func (eb Builder[T]) Code(code T) Builder[T] {
	nb := eb.clone()
	nb.code = code
	return nb
}
