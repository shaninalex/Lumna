package project

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/adapters/http/transport"
	"gitlab.com/shaninalex/lumna/app/core"
	"gitlab.com/shaninalex/lumna/app/modules/workspace/contract"
)

func Register(resolve core.Resolve, router *gin.RouterGroup) {
	router.GET("", handleProjectList(resolve))
}

func handleProjectList(resolve core.Resolve) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Query("workspace_id"))
		if err != nil {
			transport.Fail(c, err)
			return
		}

		result, err := contract.AskProjectList(c.Request.Context(), resolve(), contract.ProjectList{
			WorkspaceId: id,
		})
		if err != nil {
			transport.Fail(c, err)
			return
		}

		projects := make([]projectDTO, len(result))
		for i, project := range result {
			projects[i] = projectDTO{
				Id:          project.Id,
				Title:       project.Title,
				WorkspaceId: project.WorkspaceId,
				OwnerId:     project.OwnerId,
				Meta:        project.Meta,
				CreatedAt:   project.CreatedAt,
				UpdatedAt:   project.UpdatedAt,
			}
		}

		transport.Success(c, projects)
	}
}
