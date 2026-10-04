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
	UpdatedAt  sql.NullTime `gorm:"column:updated_at;default:null"`
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
	Role        string    `gorm:"column:role;not null"`
}

func (identityWorkspaceRecord) TableName() string { return "identity_workspaces" }

type projectRecord struct {
	ID          int            `gorm:"primaryKey;autoIncrement"`
	Title       string         `gorm:"column:title;not null"`
	WorkspaceId int            `gorm:"column:workspace_id;not null"`
	OwnerId     int            `gorm:"column:owner_id;not null"`
	Meta        sql.NullString `gorm:"column:meta"`
	CreatedAt   time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   sql.NullTime   `gorm:"column:updated_at;default:null"`
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

type workspaceInvitationRecord struct {
	ID          int           `gorm:"primaryKey;autoIncrement"`
	WorkspaceID int           `gorm:"column:workspace_id"`
	Email       string        `gorm:"column:email"`
	Role        string        `gorm:"column:role"`
	TokenHash   string        `gorm:"column:token_hash"`
	InvitedBy   sql.NullInt32 `gorm:"column:invited_by"`
	ExpiresAt   time.Time     `gorm:"column:expires_at"`
	AcceptedAt  sql.NullTime  `gorm:"column:accepted_at"`
	RevokedAt   sql.NullTime  `gorm:"column:accepted_at"`
	CreatedAt   time.Time     `gorm:"column:created_at;autoCreateTime"`
}

func (workspaceInvitationRecord) TableName() string { return "workspace_invitations" }

func toInvitationDomain(r workspaceInvitationRecord) domain.Invitation {
	return domain.Invitation{
		ID:          r.ID,
		WorkspaceID: r.WorkspaceID,
		Email:       r.Email,
		Role:        r.Role,
		TokenHash:   r.TokenHash,
		InvitedBy:   int(r.InvitedBy.Int32),
		ExpiresAt:   r.ExpiresAt,
		AcceptedAt:  r.AcceptedAt.Time,
		RevokedAt:   r.RevokedAt.Time,
		CreatedAt:   r.CreatedAt,
	}
}

func toInvitationRecord(d domain.Invitation) workspaceInvitationRecord {
	return workspaceInvitationRecord{
		ID:          d.ID,
		WorkspaceID: d.WorkspaceID,
		Email:       d.Email,
		Role:        d.Role,
		TokenHash:   d.TokenHash,
		InvitedBy:   sql.NullInt32{Int32: int32(d.InvitedBy), Valid: d.InvitedBy > 0},
		ExpiresAt:   d.ExpiresAt,
		AcceptedAt:  sql.NullTime{Time: d.AcceptedAt, Valid: d.AcceptedAt.IsZero()},
		RevokedAt:   sql.NullTime{Time: d.RevokedAt, Valid: d.RevokedAt.IsZero()},
		CreatedAt:   d.CreatedAt,
	}
}
