package inflexible

import "fmt"

func NewAppError(err error, code int) error {
	return &appError{err: err, code: code}
}

type appError struct {
	err  error
	code int
}

func (e *appError) Error() string {
	return e.err.Error()
}

func (e *appError) Code() int {
	return e.code
}

func (e *appError) Format(s fmt.State, v rune) {
	if inner, ok := e.err.(fmt.Formatter); ok {
		inner.Format(s, v)
		return
	}
	fmt.Fprintf(s, "%v", e.err)
}

type HasCode interface {
	Code() int
}
