package workspace

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/adapters/http/transport"
	"gitlab.com/shaninalex/lumna/app/core"
)

func Register(resolve core.Resolve, router *gin.RouterGroup) {
	router.GET("/", handlerList(resolve))
	router.POST("/", handlerCreate(resolve))
}

func handlerList(resolve core.Resolve) gin.HandlerFunc {
	return func(c *gin.Context) {
		workspace := []WorkspaceDTO{}
		transport.Success(c, workspace)
	}
}

func handlerCreate(resolve core.Resolve) gin.HandlerFunc {
	return func(c *gin.Context) {
		transport.Success(c, WorkspaceDTO{})
	}
}
