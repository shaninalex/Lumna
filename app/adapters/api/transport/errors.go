package transport

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/core/errs"
)

// httpStatus is the ONE place where a domain error class becomes a status
// code. No handler picks a status by hand.
var httpStatus = map[errs.Kind]int{
	errs.KindValidation:      http.StatusBadRequest,
	errs.KindUnauthenticated: http.StatusUnauthorized,
	errs.KindForbidden:       http.StatusForbidden,
	errs.KindNotFound:        http.StatusNotFound,
	errs.KindConflict:        http.StatusConflict,
	errs.KindRateLimited:     http.StatusTooManyRequests,
	errs.KindInternal:        http.StatusInternalServerError,
}

const (
	codeInternal   = "INTERNAL_ERROR"
	codeValidation = "VALIDATION_ERROR"

	// An internal error's own text may name tables, hosts or queries, so the
	// client gets a fixed string instead. The original belongs in the log.
	messageInternal = "internal server error"
)

// Fail answers with the status and code that the error itself carries.
func Fail(c *gin.Context, err error) {
	kind := errs.KindOf(err)

	status, ok := httpStatus[kind]
	if !ok {
		status = http.StatusInternalServerError
	}

	if kind == errs.KindInternal {
		ReturnJSON(c, status, nil, NewApiError(messageInternal, codeInternal, nil))
		return
	}

	code := errs.CodeOf(err)
	if code == "" {
		code = codeInternal
	}

	var meta any
	if fields := errs.FieldsOf(err); len(fields) > 0 {
		meta = fields
	}

	ReturnJSON(c, status, nil, NewApiError(err.Error(), code, meta))
}

// Invalid answers 400 for input that never reached a module: malformed JSON
// or a failed struct validation. Its text comes from the binder, so it is
// safe to echo and useful to the client.
func Invalid(c *gin.Context, err error) {
	ReturnJSON(c, http.StatusBadRequest, nil, NewApiError(err.Error(), codeValidation, nil))
}
