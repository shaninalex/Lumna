package project

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/adapters/httpx"
	"gitlab.com/shaninalex/lumna/app/core"
	"gitlab.com/shaninalex/lumna/app/core/actor"
	"gitlab.com/shaninalex/lumna/app/modules/workspace/contract"
)

func Register(resolve core.Resolve, router *gin.RouterGroup) {
	router.GET("", handleList(resolve))
	router.POST("", handleCreate(resolve))
}

func handleList(resolve core.Resolve) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Query("workspace_id"))
		if err != nil {
			httpx.Fail(c, err)
			return
		}

		result, err := contract.AskProjectList(c.Request.Context(), resolve(), contract.ProjectList{
			WorkspaceId: id,
		})
		if err != nil {
			httpx.Fail(c, err)
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

		httpx.Success(c, projects)
	}
}

func handleCreate(resolve core.Resolve) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data projectCreateDTO
		if err := c.ShouldBindJSON(&data); err != nil {
			httpx.Fail(c, err)
			return
		}
		if err := data.Validate(); err != nil {
			httpx.Fail(c, err)
			return
		}

		a, _ := actor.From(c.Request.Context())
		project, err := contract.ExecCreateProject(c.Request.Context(), resolve(), contract.CreateProject{
			Title:       data.Title,
			OwnerId:     a.IdentityID,
			WorkspaceId: data.WorkspaceId,
		})
		if err != nil {
			httpx.Fail(c, err)
			return
		}

		httpx.Success(c, projectDTO{
			Id:          project.Id,
			Title:       project.Title,
			WorkspaceId: project.WorkspaceId,
			OwnerId:     project.OwnerId,
			Meta:        project.Meta,
			CreatedAt:   project.CreatedAt,
			UpdatedAt:   project.UpdatedAt,
		})
	}
}
