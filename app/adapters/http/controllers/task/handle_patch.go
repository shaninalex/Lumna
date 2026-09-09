package task

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/api/utils"
	"gitlab.com/shaninalex/lumna/app/models"
)

func (s *TaskController) handlePatchTask(c *gin.Context) {
	taskId, err := strconv.Atoi(c.Param("task_id"))
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	task, err := s.taskService.GetTask(c.Request.Context(), taskId)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	payload := models.Task{}
	if err = c.ShouldBindJSON(&payload); err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	task.Title = payload.Title
	task.Body = payload.Body
	task.Completed = payload.Completed

	if err = s.taskService.UpdateTask(c.Request.Context(), taskId, task); err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	utils.Success(c, task)
}
