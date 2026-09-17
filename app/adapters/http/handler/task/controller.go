package task

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/adapters/http/transport"
	"gitlab.com/shaninalex/lumna/app/core"
	"gitlab.com/shaninalex/lumna/app/modules/tracker/contract"
)

func Register(resolve core.Resolve, router *gin.RouterGroup) {
	router.GET("", handleList(resolve))
	router.POST("", handleCreate(resolve))
}

func handleCreate(resolve core.Resolve) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data taskCreateDTO
		if err := c.ShouldBindJSON(&data); err != nil {
			transport.Fail(c, err)
			return
		}
		result, err := contract.ExecWorkItemCreate(c.Request.Context(), resolve(), contract.WorkItemCreate{
			Title:       data.Title,
			Description: &data.Body,
			ProjectId:   data.ProjectId,
			Position:    data.Position,
			StageId:     &data.ColumnId,
			ScopeId:     &data.BoardId,
		})
		if err != nil {
			transport.Fail(c, err)
			return
		}
		transport.Success(c, toDTO(result))
	}
}

func handleList(resolve core.Resolve) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Query("board_id"))
		if err != nil {
			transport.Fail(c, err)
			return
		}
		result, err := contract.AskWorkItemList(c.Request.Context(), resolve(), contract.WorkItemList{
			ScopeId: id,
		})
		if err != nil {
			transport.Fail(c, err)
			return
		}

		tasks := make([]taskDTO, len(result))
		for i, item := range result {
			tasks[i] = toDTO(item)
		}
		transport.Success(c, tasks)
	}
}
