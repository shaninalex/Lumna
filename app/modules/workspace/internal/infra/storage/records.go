package storage

import (
	"time"

	"gitlab.com/shaninalex/lumna/app/modules/workspace/internal/domain"
)

type workspaceRecord struct {
	ID         int        `gorm:"primaryKey;autoIncrement"`
	Title      string     `gorm:"column:title;unique;not null"`
	Active     bool       `gorm:"column:active;not null;default:false"`
	OwnerEmail string     `gorm:"column:owner_email;not null"`
	CreatedAt  time.Time  `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt  *time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (workspaceRecord) TableName() string { return "workspaces" }

func toDomainWorkspace(w workspaceRecord) domain.Workspace {
	return domain.Workspace{
		ID:         w.ID,
		Title:      w.Title,
		OwnerEmail: w.OwnerEmail,
		Active:     w.Active,
		CreatedAt:  w.CreatedAt,
		UpdatedAt:  w.UpdatedAt,
	}
}

type identityWorkspaceRecord struct {
	IdentityID  int       `gorm:"primaryKey"`
	WorkspaceID int       `gorm:"primaryKey"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime"`
}

func (identityWorkspaceRecord) TableName() string { return "identity_workspaces" }
