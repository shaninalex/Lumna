package workspace

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/adapters/httpx"
	"gitlab.com/shaninalex/lumna/app/core"
	"gitlab.com/shaninalex/lumna/app/modules/workspace/contract"
)

func Register(resolve core.Resolve, router *gin.RouterGroup) {
	router.GET("", handleWorkspaceList(resolve))
	router.POST("", handleCreateWorkspace(resolve))
	router.GET(":workspace_id/members", handleWorkspaceMembers(resolve))
	router.POST(":workspace_id/invitations", handleCreateInvitation(resolve))
}

func handleWorkspaceList(resolve core.Resolve) gin.HandlerFunc {
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

func handleCreateWorkspace(resolve core.Resolve) gin.HandlerFunc {
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

func handleCreateInvitation(resolve core.Resolve) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data invitationDTO
		if err := c.ShouldBindJSON(&data); err != nil {
			httpx.Fail(c, err)
			return
		}
		if err := data.Validate(); err != nil {
			httpx.Fail(c, err)
			return
		}

		invitation, err := contract.ExecCreateInvitation(c.Request.Context(), resolve(), contract.CreateInvitation{})
		if err != nil {
			httpx.Fail(c, err)
			return
		}
		httpx.Success(c, toInvitationDTO(invitation))
	}
}

func handleWorkspaceMembers(resolve core.Resolve) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("workspace_id"))
		if err != nil {
			httpx.Fail(c, err)
			return
		}
		memberDomains, err := contract.AskMembersList(c.Request.Context(), resolve(), contract.MembersList{
			WorkspaceId: id,
		})
		if err != nil {
			httpx.Fail(c, err)
			return
		}
		members := make([]memberDTO, len(memberDomains))
		for i, m := range memberDomains {
			members[i] = toMemberDTO(m)
		}
		httpx.Success(c, members)
	}
}
