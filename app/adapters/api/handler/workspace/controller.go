package workspace

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/adapters/httpx"
	"gitlab.com/shaninalex/lumna/app/core"
	"gitlab.com/shaninalex/lumna/app/modules/workspace/contract"
)

func Register(resolve core.Resolve, router *gin.RouterGroup) {
	router.GET("", handlerList(resolve))
	router.POST("", handlerCreate(resolve))
}

func handlerList(resolve core.Resolve) gin.HandlerFunc {
	return func(c *gin.Context) {
		result, err := contract.AskWorkspaceList(c.Request.Context(), resolve(), contract.WorkspaceList{})
		if err != nil {
			httpx.Fail(c, err)
			return
		}

		workspaces := make([]workspaceDTO, len(result))
		for i, w := range result {
			workspaces[i] = workspaceDTO{
				ID:         w.Id,
				Title:      w.Title,
				Active:     w.Active,
				OwnerEmail: w.OwnerEmail,
				CreatedAt:  w.CreatedAt,
			}
		}

		httpx.Success(c, workspaces)
	}
}

func handlerCreate(resolve core.Resolve) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data workspaceCreateDTO
		if err := c.ShouldBindJSON(&data); err != nil {
			httpx.Fail(c, err)
			return
		}

		if err := data.Validate(); err != nil {
			httpx.Fail(c, err)
			return
		}

		workspace, err := contract.ExecCreateWorkspace(c.Request.Context(), resolve(), contract.CreateWorkspace{
			Title:      data.Title,
			OwnerEmail: data.Email,
			Active:     true,
		})
		if err != nil {
			httpx.Fail(c, err)
			return
		}

		httpx.Success(c, workspaceDTO{
			ID:         workspace.Id,
			Title:      workspace.Title,
			Active:     workspace.Active,
			OwnerEmail: workspace.OwnerEmail,
			CreatedAt:  workspace.CreatedAt,
		})
	}
}
