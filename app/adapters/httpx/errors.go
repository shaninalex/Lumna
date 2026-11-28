package httpx

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/core/errs"
)

// httpStatus is the ONE place where a domain error class becomes a status
// code. No handler picks a status by hand.
var httpStatus = map[errs.Kind]int{
	errs.KindInternal:        http.StatusInternalServerError,
	errs.KindValidation:      http.StatusBadRequest,
	errs.KindNotFound:        http.StatusNotFound,
	errs.KindConflict:        http.StatusConflict,
	errs.KindForbidden:       http.StatusForbidden,
	errs.KindUnauthenticated: http.StatusUnauthorized,
	errs.KindRateLimited:     http.StatusTooManyRequests,
	errs.KindPlatform:        http.StatusInternalServerError,
}

var (
	ErrUnknown = errors.New("unknown error")
)

// Fail answers with the status and code that the error itself carries.
func Fail(c *gin.Context, err error) {
	e, ok := errors.AsType[*errs.Error](err)
	if !ok {
		log.Println(err)
		// prevent leak internals to the userspace
		ReturnJSON(c, 400, nil, ErrUnknown)
		return
	}

	status, ok := httpStatus[e.Kind]
	if !ok {
		status = http.StatusBadRequest
	}

	ReturnJSON(c, status, nil, NewApiError(e.Message, e.Data))
}
