package httpx

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type Validatable interface {
	Verify() error
}

func BindPayload[T Validatable](c *gin.Context, payload T) (T, error) {
	var zero T
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
