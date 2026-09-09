package errs

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
	panic("not implemented")
}

func Conflict(code, msg string) *Error {
	panic("not implemented")
}

func Forbidden(code, msg string) *Error {
	panic("not implemented")
}

func Validation(code, msg string) *Error {
	panic("not implemented")
}

func KindOf(err error) Kind {
	panic("not implemented")
}

// через errors.As
