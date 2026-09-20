package column

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/adapters/api/transport"
	"gitlab.com/shaninalex/lumna/app/core"
	"gitlab.com/shaninalex/lumna/app/modules/tracker/contract"
)

func Register(resolve core.Resolve, router *gin.RouterGroup) {
	router.GET("", handleList(resolve))
	router.POST("", handleCreate(resolve))
	router.DELETE(":stage_id", handleDelete(resolve))
}

func handleList(resolve core.Resolve) gin.HandlerFunc {
	return func(c *gin.Context) {
		boardId, err := strconv.Atoi(c.Query("board_id"))
		if err != nil {
			transport.Fail(c, err)
			return
		}

		stages, err := contract.AskStageList(c.Request.Context(), resolve(), contract.StageList{ScopeId: boardId})
		if err != nil {
			transport.Fail(c, err)
			return
		}

		columns := make([]columnDto, len(stages))
		for i, stage := range stages {
			columns[i] = columnDto{
				Id:        stage.Id,
				Title:     stage.Name,
				Meta:      newDefaultColumnMeta(),
				BoardId:   stage.ScopeId,
				Position:  stage.Position,
				CreatedAt: stage.CreatedAt,
				UpdatedAt: &stage.UpdatedAt,
			}
		}

		transport.Success(c, columns)
	}
}

func handleCreate(resolve core.Resolve) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data columnCreateDto
		if err := c.ShouldBindJSON(&data); err != nil {
			transport.Fail(c, err)
			return
		}
		if err := data.Validate(); err != nil {
			transport.Fail(c, err)
			return
		}

		stage, err := contract.ExecCreateStage(c.Request.Context(), resolve(), contract.StageCreate{
			ScopeID:     data.BoardId,
			Name:        data.Title,
			Description: "",
			Category:    "backlog",
			Position:    data.Position,
		})
		if err != nil {
			transport.Fail(c, err)
			return
		}

		transport.Success(c, columnDto{
			Id:        stage.Id,
			Title:     stage.Name,
			Meta:      newDefaultColumnMeta(),
			BoardId:   stage.ScopeId,
			Position:  stage.Position,
			CreatedAt: stage.CreatedAt,
			UpdatedAt: &stage.UpdatedAt,
		})
	}
}

type stageDeleteQuery struct {
	WithTasks *bool `form:"with_tasks,omitempty"`
}

func handleDelete(resolve core.Resolve) gin.HandlerFunc {
	return func(c *gin.Context) {
		stageId, err := strconv.Atoi(c.Param("stage_id"))
		if err != nil {
			transport.Fail(c, err)
			return
		}

		var query stageDeleteQuery
		if err = c.ShouldBindQuery(&query); err != nil {
			transport.Fail(c, err)
			return
		}

		data := contract.StageDelete{StageId: stageId, WithTasks: false}
		if query.WithTasks != nil {
			if *query.WithTasks == true {
				data.WithTasks = true
			}
		}

		commandResponse, err := contract.ExecStageDelete(c.Request.Context(), resolve(), data)
		if err != nil {
			transport.Fail(c, err)
			return
		}

		transport.Success(c, columnDeleteResponseDto{
			Id:           commandResponse.StageId,
			DeletedTasks: commandResponse.DeletedTasks,
		}, "Stage deleted")
	}
}
