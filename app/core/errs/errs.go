package errs

import "errors"

type Kind uint8

const (
	KindInternal Kind = iota
	KindValidation
	KindNotFound
	KindConflict
	KindForbidden
	KindUnauthenticated
	KindRateLimited
)

type Error struct {
	Kind    Kind
	Code    string
	Message string
	Fields  map[string]string
	cause   error
}

func (e *Error) Error() string { return e.Message }
func (e *Error) Unwrap() error { return e.cause }

func NotFound(code, msg string) *Error {
	return &Error{
		Kind:    KindNotFound,
		Code:    code,
		Message: msg,
		Fields:  nil,
		cause:   errors.New(msg),
	}
}

func Conflict(code, msg string) *Error {
	return &Error{
		Kind:    KindConflict,
		Code:    code,
		Message: msg,
		Fields:  nil,
		cause:   errors.New(msg),
	}
}

func Forbidden(code, msg string) *Error {
	return &Error{
		Kind:    KindForbidden,
		Code:    code,
		Message: msg,
		Fields:  nil,
		cause:   errors.New(msg),
	}
}

func Validation(code, msg string) *Error {
	return &Error{
		Kind:    KindValidation,
		Code:    code,
		Message: msg,
		Fields:  nil,
		cause:   errors.New(msg),
	}
}

func KindOf(err error) Kind {
	return KindInternal
}

func Unauthenticated(code, msg string) *Error {
	return &Error{
		Kind:    KindUnauthenticated,
		Code:    code,
		Message: msg,
		Fields:  nil,
		cause:   errors.New(msg),
	}
}
