package board

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/adapters/httpx"
	"gitlab.com/shaninalex/lumna/app/core"
	"gitlab.com/shaninalex/lumna/app/modules/tracker/contract"
)

func Register(resolve core.Resolve, router *gin.RouterGroup) {
	router.GET("", handleList(resolve))
	router.POST("", handleCreate(resolve))
}

func handleList(resolve core.Resolve) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Query("project_id"))
		if err != nil {
			httpx.Fail(c, err)
			return
		}

		result, err := contract.AskScopeList(c.Request.Context(), resolve(), contract.ScopeList{
			ProjectId: id,
		})
		if err != nil {
			httpx.Fail(c, err)
			return
		}

		boards := make([]boardDTO, len(result))
		for i, scope := range result {
			boards[i] = boardDTO{
				Id:        scope.Id,
				Title:     scope.Name,
				ProjectId: id,
				CreatedAt: time.Now(),
			}
		}

		httpx.Success(c, boards)
	}
}

func handleCreate(resolve core.Resolve) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data boardCreateDTO
		if err := c.ShouldBindJSON(&data); err != nil {
			httpx.Fail(c, err)
			return
		}
		if err := data.Validate(); err != nil {
			httpx.Fail(c, err)
			return
		}
		scope, err := contract.ExecCreateScope(c.Request.Context(), resolve(), contract.CreateScope{
			Name:        data.Title,
			Description: "",
			ProjectId:   data.ProjectId,
		})
		if err != nil {
			httpx.Fail(c, err)
			return
		}

		httpx.Success(c, boardDTO{
			Id:        scope.Id,
			Title:     scope.Name,
			ProjectId: scope.ProjectId,
			CreatedAt: scope.CreatedAt,
			UpdatedAt: scope.UpdatedAt,
		})
	}
}
