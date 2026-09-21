package storage

import (
	"database/sql"
	"time"

	"gitlab.com/shaninalex/lumna/app/modules/workspace/internal/domain"
)

type workspaceRecord struct {
	ID         int          `gorm:"primaryKey;autoIncrement"`
	Title      string       `gorm:"column:title;unique;not null"`
	Active     bool         `gorm:"column:active;not null;default:false"`
	OwnerEmail string       `gorm:"column:owner_email;not null"`
	CreatedAt  time.Time    `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt  sql.NullTime `gorm:"column:updated_at"`
}

func (workspaceRecord) TableName() string { return "workspaces" }

func toDomainWorkspace(w workspaceRecord) domain.Workspace {
	return domain.Workspace{
		ID:         w.ID,
		Title:      w.Title,
		OwnerEmail: w.OwnerEmail,
		Active:     w.Active,
		CreatedAt:  w.CreatedAt,
		UpdatedAt:  w.UpdatedAt.Time,
	}
}

type identityWorkspaceRecord struct {
	IdentityID  int       `gorm:"primaryKey"`
	WorkspaceID int       `gorm:"primaryKey"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime"`
}

func (identityWorkspaceRecord) TableName() string { return "identity_workspaces" }

type projectRecord struct {
	ID          int            `gorm:"primaryKey;autoIncrement"`
	Title       string         `gorm:"column:title;not null"`
	WorkspaceId int            `gorm:"column:workspace_id;not null"`
	OwnerId     int            `gorm:"column:owner_id;not null"`
	Meta        sql.NullString `gorm:"column:meta"`
	CreatedAt   time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   sql.NullTime   `gorm:"column:updated_at"`
}

func (projectRecord) TableName() string { return "projects" }

func toDomainProject(project projectRecord) domain.Project {
	return domain.Project{
		ID:          project.ID,
		Title:       project.Title,
		WorkspaceId: project.WorkspaceId,
		OwnerId:     project.OwnerId,
		Meta:        project.Meta.String,
		CreatedAt:   project.CreatedAt,
		UpdatedAt:   project.UpdatedAt.Time,
	}
}
