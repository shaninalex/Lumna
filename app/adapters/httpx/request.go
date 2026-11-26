package httpx

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Validatable interface {
	Verify() error
}

// ErrUnsupportedMediaType is returned by BindPayload for a body that is not
// declared as JSON; Invalid turns it into 415.
var ErrUnsupportedMediaType = errors.New("Content-Type must be application/json")

// BindPayload decodes a JSON body into payload and verifies it.
//
// The body must be declared as application/json: gin would otherwise decode
// a text/plain body as JSON, and text/plain is what a cross-site HTML form
// can send (CSRF, issue #16). Endpoints taking other formats, such as file
// uploads, read the body themselves and never call this.
func BindPayload[T Validatable](c *gin.Context, payload T) (T, error) {
	var zero T
	if c.ContentType() != "application/json" {
		return zero, ErrUnsupportedMediaType
	}
	if err := c.ShouldBindJSON(&payload); err != nil {
		return zero, err
	}
	if err := payload.Verify(); err != nil {
		return zero, err
	}

	return payload, nil
}

func GetParamInt(c *gin.Context, name string) (int, error) {
	return strconv.Atoi(c.Param(name))
}
