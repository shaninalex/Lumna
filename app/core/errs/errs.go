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

// KindOf reports the classification of err, unwrapping as needed.
func KindOf(err error) Kind {
	var e *Error
	if errors.As(err, &e) {
		return e.Kind
	}
	return KindInternal
}

// CodeOf returns the machine-readable code of a classified error, or "".
func CodeOf(err error) string {
	var e *Error
	if errors.As(err, &e) {
		return e.Code
	}
	return ""
}

// FieldsOf returns per-field validation details, or nil.
func FieldsOf(err error) map[string]string {
	var e *Error
	if errors.As(err, &e) {
		return e.Fields
	}
	return nil
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
