package handlers

import (
	"context"
	"errors"
	"log/slog"

	"gitlab.com/shaninalex/lumna/app/core/actor"
	"gitlab.com/shaninalex/lumna/app/modules/workspace/contract"
	"gitlab.com/shaninalex/lumna/app/modules/workspace/internal/domain"
	"gitlab.com/shaninalex/lumna/app/platform/clock"
)

type CreateWorkspace struct {
	workspaceRepo         domain.WorkspaceRepo
	identityWorkspaceRepo domain.IdentityWorkspaceRepo
	clock                 clock.Clock
	log                   *slog.Logger
}

func NewCreateWorkspace(workspaceRepo domain.WorkspaceRepo, identityWorkspaceRepo domain.IdentityWorkspaceRepo, log *slog.Logger, clk clock.Clock) *CreateWorkspace {
	return &CreateWorkspace{
		workspaceRepo:         workspaceRepo,
		identityWorkspaceRepo: identityWorkspaceRepo,
		clock:                 clk,
		log:                   log,
	}
}

func (s *CreateWorkspace) Handle(ctx context.Context, cmd contract.CreateWorkspace) (contract.WorkspaceView, error) {
	workspace, err := s.workspaceRepo.Save(ctx, domain.NewWorkspace(cmd.Title, cmd.OwnerEmail, cmd.Active, s.clock.Now()))
	if err != nil {
		return contract.WorkspaceView{}, err
	}
	if a, ok := actor.From(ctx); ok && a.IdentityID != 0 {
		_, err = s.identityWorkspaceRepo.Save(ctx, domain.IdentityWorkspace{
			IdentityID:  a.IdentityID,
			WorkspaceID: workspace.ID,
			CreatedAt:   s.clock.Now(),
			Role:        "admin", // TODO: global enum
		})
		if err != nil {
			s.log.Error(errors.Join(domain.ErrUnableToAddUserToWorkspace, err).Error())
		}
	}
	return toWorkspaceView(workspace), nil
}
