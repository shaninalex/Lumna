package task

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/adapters/httpx"
	"gitlab.com/shaninalex/lumna/app/core"
	"gitlab.com/shaninalex/lumna/app/core/errs"
	"gitlab.com/shaninalex/lumna/app/modules/tracker/contract"
)

func Register(resolve core.Resolve, router *gin.RouterGroup) {
	router.GET("", handleList(resolve))
	router.POST("", handleCreate(resolve))
	router.POST("move", handleBoardAction(resolve))
	router.PATCH(":task_id", handleTaskUpdate(resolve))
	router.PATCH(":task_id/assign", handleTaskAssignment(resolve))
	router.DELETE(":task_id", handleTaskDelete(resolve))
}

func handleList(resolve core.Resolve) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Query("board_id"))
		if err != nil {
			httpx.Fail(c, err)
			return
		}
		results, err := contract.AskWorkItemList(c.Request.Context(), resolve(), contract.WorkItemList{
			ScopeId: id,
		})
		if err != nil {
			httpx.Fail(c, err)
			return
		}

		tasks := make([]taskDTO, len(results))
		for i, task := range results {
			tasks[i] = toTaskDTO(task)
		}

		httpx.Success(c, tasks)
	}
}

func handleCreate(resolve core.Resolve) gin.HandlerFunc {
	return func(c *gin.Context) {
		data := httpx.BindPayload(c, taskCreateDTO{})
		result, err := contract.ExecWorkItemCreate(c.Request.Context(), resolve(), contract.WorkItemCreate{
			Title:       data.Title,
			Description: data.Body,
			ProjectId:   data.ProjectId,
			Position:    data.Position,
			StageId:     data.ColumnId,
			ScopeId:     data.BoardId,
			DueTo:       data.DueTo,
		})
		if err != nil {
			httpx.Fail(c, err)
			return
		}
		httpx.Success(c, toTaskDTO(result))
	}
}

func handleBoardAction(resolve core.Resolve) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data boardActionPayload
		if err := c.ShouldBindJSON(&data); err != nil {
			httpx.Fail(c, err)
			return
		}

		switch data.Action {
		case boardActionMoveTask:
			var action boardActionMoveTaskDTO
			if err := json.Unmarshal(data.Data, &action); err != nil {
				httpx.Fail(c, err)
				return
			}
			actionMoveWorkItem(c, resolve, action)
			return

		case boardActionMoveColumn:
			var action boardActionMoveColumnDTO
			if err := json.Unmarshal(data.Data, &action); err != nil {
				httpx.Fail(c, err)
				return
			}
			actionMoveColumn(c, resolve, action)
			return

		case boardActionChangeStage:
			var action boardActionTransferTaskDTO
			if err := json.Unmarshal(data.Data, &action); err != nil {
				httpx.Fail(c, err)
				return
			}
			actionTransferWorkItem(c, resolve, action)
			return

		default:
			httpx.Fail(c, errs.Validation("unknown_board_action", "unknown board action"))
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
		httpx.Fail(c, err)
		return
	}
	httpx.Success(c, toTaskDTO(result))
}

func actionTransferWorkItem(c *gin.Context, resolve core.Resolve, action boardActionTransferTaskDTO) {
	result, err := contract.ExecWorkItemTransfer(c.Request.Context(), resolve(), contract.WorkItemTransfer{
		WorkItemId: action.TaskId,
		Rank:       action.Position,
		ScopeId:    action.BoardId,
		StageId:    action.ColumnId,
	})
	if err != nil {
		httpx.Fail(c, err)
		return
	}
	httpx.Success(c, toTaskDTO(result))
}

func actionMoveColumn(c *gin.Context, resolve core.Resolve, action boardActionMoveColumnDTO) {
	result, err := contract.ExecStageMove(c.Request.Context(), resolve(), contract.StageMove{
		StageId:  action.ColumnId,
		Position: action.Position,
		ScopeId:  action.BoardId,
	})
	if err != nil {
		httpx.Fail(c, err)
		return
	}
	httpx.Success(c, toColumnDTO(result))
}

func handleTaskUpdate(resolve core.Resolve) gin.HandlerFunc {
	return func(c *gin.Context) {
		data := httpx.BindPayload(c, taskUpdateDTO{})
		result, err := contract.ExecWorkItemUpdate(c.Request.Context(), resolve(), contract.WorkItemUpdate{
			WorkItemId:  data.TaskId,
			Title:       data.Title,
			Description: data.Body,
		})
		if err != nil {
			httpx.Fail(c, err)
			return
		}
		httpx.Success(c, toTaskDTO(result))
	}
}

func handleTaskAssignment(resolve core.Resolve) gin.HandlerFunc {
	return func(c *gin.Context) {
		data := httpx.BindPayload(c, taskAssignmentDTO{})
		_, err := contract.ExecWorkItemAssign(c.Request.Context(), resolve(), contract.WorkItemAssign{
			WorkItemId: data.TaskId,
			IdentityId: data.IdentityId,
		})
		if err != nil {
			httpx.Fail(c, err)
			return
		}
		httpx.Success(c, data)
	}
}

func handleTaskDelete(resolve core.Resolve) gin.HandlerFunc {
	return func(c *gin.Context) {
		taskId, err := strconv.Atoi(c.Param("task_id"))
		if err != nil {
			httpx.Fail(c, err)
			return
		}
		_, err = contract.ExecWorkItemDelete(c.Request.Context(), resolve(), contract.WorkItemDelete{
			WorkItemId: taskId,
		})
		if err != nil {
			httpx.Fail(c, err)
			return
		}
		httpx.Success(c, true, fmt.Sprintf("Task [%d] deleted.", taskId))
	}
}
