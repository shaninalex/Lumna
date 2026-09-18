package transport

import "github.com/gin-gonic/gin"

type Validatable interface {
	Verify() error
}

func BindPayload[T Validatable](c *gin.Context, payload T) *T {
	if err := c.ShouldBindJSON(&payload); err != nil {
		Fail(c, err)
		return nil
	}
	if err := payload.Verify(); err != nil {
		Fail(c, err)
		return nil
	}

	return &payload
}
