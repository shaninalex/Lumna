package task

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/api/utils"
	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/services"
)

func (s *TaskController) handletransferTask(c *gin.Context) {
	var data models.TransferTaskBetweenColumns
	if err := c.BindJSON(&data); err != nil {
		utils.Error(c, 400, err)
		return
	}

	if err := s.taskService.Transfer(c.Request.Context(), data); err != nil {
		utils.Error(c, 400, err)
		return
	}

	result, err := s.queryTaskList(c.Request.Context(), services.ServiceTaskListQuery{
		BoardId: data.BoardId,
	})
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}
	utils.Success(c, result)
}
