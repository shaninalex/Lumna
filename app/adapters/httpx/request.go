package httpx

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/core/errs"
)

type Validatable interface {
	Verify() error
}

var (
	errUnsupportedMediaType = errs.Conflict("VLD000", "Content-Type must be application/json")
)

// BindPayload decodes a JSON body into payload and verifies it.
func BindPayload[T Validatable](c *gin.Context, payload T) (T, error) {
	var zero T
	if c.ContentType() != "application/json" {
		return zero, errUnsupportedMediaType
	}
	if err := c.ShouldBindJSON(&payload); err != nil {
		return zero, errs.Validation("VLD002", "invalid payload", err.Error())
	}
	if err := payload.Verify(); err != nil {
		return zero, err
	}

	return payload, nil
}

func GetParamInt(c *gin.Context, name string) (int, error) {
	v, err := strconv.Atoi(c.Param(name))
	if err != nil {
		return 0, errs.Validation("VLD001", "Validation error", map[string]any{
			name: fmt.Sprintf("%s is in wrong type", name),
		})
	}
	return v, nil
}
