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
	KindPlatform
)

type Error struct {
	Kind    Kind
	Code    string
	Message string
	Data    any
	cause   error
}

func (e *Error) Error() string { return e.Message }
func (e *Error) Unwrap() error { return e.cause }

// Error parsing helper functions

// KindOf reports the classification of err, unwrapping as needed.
func KindOf(err error) Kind {
	if e, ok := errors.AsType[*Error](err); ok {
		return e.Kind
	}
	return KindInternal
}

// CodeOf returns the machine-readable code of a classified error, or "".
func CodeOf(err error) string {
	if e, ok := errors.AsType[*Error](err); ok {
		return e.Code
	}
	return ""
}

// Constructors

func Internal(code, msg string) *Error {
	return &Error{
		Kind:    KindInternal,
		Code:    code,
		Message: msg,
		Data:    nil,
		cause:   errors.New(msg),
	}
}

func Validation(code, msg string, data ...any) *Error {
	return &Error{
		Kind:    KindValidation,
		Code:    code,
		Message: msg,
		Data:    data,
		cause:   errors.New(msg),
	}
}

func NotFound(code, msg string) *Error {
	return &Error{
		Kind:    KindNotFound,
		Code:    code,
		Message: msg,
		Data:    nil,
		cause:   errors.New(msg),
	}
}

func Conflict(code, msg string) *Error {
	return &Error{
		Kind:    KindConflict,
		Code:    code,
		Message: msg,
		Data:    nil,
		cause:   errors.New(msg),
	}
}

func Forbidden(code, msg string) *Error {
	return &Error{
		Kind:    KindForbidden,
		Code:    code,
		Message: msg,
		Data:    nil,
		cause:   errors.New(msg),
	}
}

func Unauthenticated(code, msg string, data ...any) *Error {
	return &Error{
		Kind:    KindUnauthenticated,
		Code:    code,
		Message: msg,
		Data:    data,
		cause:   errors.New(msg),
	}
}

func RateLimited(code, msg string) *Error {
	return &Error{
		Kind:    KindRateLimited,
		Code:    code,
		Message: msg,
		Data:    nil,
		cause:   errors.New(msg),
	}
}

func Platform(code, msg string) *Error {
	return &Error{
		Kind:    KindPlatform,
		Code:    code,
		Message: msg,
		Data:    nil,
		cause:   errors.New(msg),
	}
}
