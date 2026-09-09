package task

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/services"
)

type TaskController struct {
	taskService        services.TaskService
	entityEventService services.EntityEventService
}

func NewTaskController(taskService services.TaskService, entityEventService services.EntityEventService) *TaskController {
	return &TaskController{
		taskService:        taskService,
		entityEventService: entityEventService,
	}
}

func (s *TaskController) Register(router *gin.RouterGroup) {
	router.POST("tasks", s.handleCreateTask)
	router.GET("tasks", s.handleListTask)
	router.POST("tasks/move", s.handleMoveTask)
	router.POST("tasks/transfer", s.handletransferTask)
	router.GET("tasks/:task_id", s.handleGetTask)
	router.PATCH("tasks/:task_id", s.handlePatchTask)
}
