package task

import (
	"encoding/json"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/adapters/http/transport"
	"gitlab.com/shaninalex/lumna/app/core"
	"gitlab.com/shaninalex/lumna/app/core/errs"
	"gitlab.com/shaninalex/lumna/app/modules/tracker/contract"
)

func Register(resolve core.Resolve, router *gin.RouterGroup) {
	router.GET("", handleList(resolve))
	router.POST("", handleCreate(resolve))
	router.POST("move", handleBoardAction(resolve))
}

func handleList(resolve core.Resolve) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Query("board_id"))
		if err != nil {
			transport.Fail(c, err)
			return
		}

		results, err := contract.AskWorkItemList(c.Request.Context(), resolve(), contract.WorkItemList{
			ScopeId: id,
		})
		if err != nil {
			transport.Fail(c, err)
			return
		}

		tasks := make([]taskDTO, len(results))
		for i, task := range results {
			tasks[i] = toDTO(task)
		}

		transport.Success(c, tasks)
	}
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
			DueTo:       data.DueTo,
		})
		if err != nil {
			transport.Fail(c, err)
			return
		}
		transport.Success(c, toDTO(result))
	}
}

func handleBoardAction(resolve core.Resolve) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data boardActionPayload
		if err := c.ShouldBindJSON(&data); err != nil {
			transport.Fail(c, err)
			return
		}

		switch data.Action {
		case boardActionMoveTask:
			var action boardActionMoveTaskDTO
			if err := json.Unmarshal(data.Data, &action); err != nil {
				transport.Fail(c, err)
				return
			}
			actionMoveWorkItem(c, resolve, action)
			return

		case boardActionMoveColumn:
			var action boardActionMoveColumnDTO
			if err := json.Unmarshal(data.Data, &action); err != nil {
				transport.Fail(c, err)
				return
			}
			actionMoveColumn(c, resolve, action)
			return

		case boardActionChangeStage:
			var action boardActionTransferTaskDTO
			if err := json.Unmarshal(data.Data, &action); err != nil {
				transport.Fail(c, err)
				return
			}
			actionTransferWorkItem(c, resolve, action)
			return

		default:
			transport.Fail(c, errs.Validation("unknown_board_action", "unknown board action"))
			return
		}
	}
}

func actionMoveWorkItem(c *gin.Context, resolve core.Resolve, action boardActionMoveTaskDTO) {
	result, err := contract.ExecWorkItemMove(c.Request.Context(), resolve(), contract.WorkItemMove{
		WorkItemId: action.TaskId,
		Rank:       action.Position,
		ScopeId:    action.BoardId,
	})
	if err != nil {
		transport.Fail(c, err)
		return
	}
	transport.Success(c, toDTO(result))
}

func actionTransferWorkItem(c *gin.Context, resolve core.Resolve, action boardActionTransferTaskDTO) {
	result, err := contract.ExecWorkItemTransfer(c.Request.Context(), resolve(), contract.WorkItemTransfer{
		WorkItemId: action.TaskId,
		Rank:       action.Position,
		ScopeId:    action.BoardId,
		StageId:    action.ColumnId,
	})
	if err != nil {
		transport.Fail(c, err)
		return
	}
	transport.Success(c, toDTO(result))
}

func actionMoveColumn(c *gin.Context, resolve core.Resolve, action boardActionMoveColumnDTO) {
	result, err := contract.ExecStageMove(c.Request.Context(), resolve(), contract.StageMove{
		StageId:  action.ColumnId,
		Position: action.Position,
		ScopeId:  action.BoardId,
	})
	if err != nil {
		transport.Fail(c, err)
		return
	}
	transport.Success(c, toColumnDTO(result))
}
