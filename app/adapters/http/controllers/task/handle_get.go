package task

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/api/utils"
)

func (s *TaskController) handleGetTask(c *gin.Context) {
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

	utils.Success(c, task)
}
