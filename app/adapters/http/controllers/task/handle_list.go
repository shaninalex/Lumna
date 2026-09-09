package task

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/api/adapters"
	"gitlab.com/shaninalex/lumna/app/api/utils"
)

func (s *TaskController) handleListTask(c *gin.Context) {
	q, err := adapters.ParseQueryParams(c)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}

	if q == nil {
		utils.Error(c, http.StatusBadRequest, adapters.TaskListParseParamsError)
		return
	}

	result, err := s.queryTaskList(c.Request.Context(), *q)
	if err != nil {
		utils.Error(c, http.StatusBadRequest, err)
		return
	}
	utils.Success(c, result)
}
