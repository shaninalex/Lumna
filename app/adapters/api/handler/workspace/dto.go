package workspace

import (
	"time"

	"github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"gitlab.com/shaninalex/lumna/app/modules/workspace/contract"
)

type workspaceDTO struct {
	ID         int       `json:"id"`
	Title      string    `json:"title"`
	Active     bool      `json:"active"`
	OwnerEmail string    `json:"owner_email"`
	CreatedAt  time.Time `json:"created_at"`
}

type workspaceCreateDTO struct {
	Title string `json:"title" binding:"required"`
	Email string `json:"email" binding:"required"`
}

func (a workspaceCreateDTO) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Title, validation.Required, validation.Length(3, 50)),
		validation.Field(&a.Email, validation.Required, is.Email),
	)
}

type invitationDTO struct {
	ID          int        `json:"id"`
	WorkspaceID int        `json:"workspace_id"`
	Email       string     `json:"email"`
	Role        string     `json:"role"`
	TokenHash   string     `json:"token_hash"`
	InvitedBy   *int       `json:"invited_by,omitempty"`
	ExpiresAt   time.Time  `json:"expires_at"`
	AcceptedAt  *time.Time `json:"accepted_at,omitzero"`
	RevokedAt   *time.Time `json:"revoked_at,omitzero"`
	CreatedAt   time.Time  `json:"created_at"`
}

func (a invitationDTO) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.WorkspaceID, validation.Required),
		validation.Field(&a.Email, validation.Required, is.Email),
		validation.Field(&a.Role, validation.Required),
		validation.Field(&a.ExpiresAt, validation.Min(time.Now())),
	)
}

func toInvitationView(inv invitationDTO) contract.InvitationView {
	return contract.InvitationView{
		ID:          inv.ID,
		WorkspaceID: inv.WorkspaceID,
		Email:       inv.Email,
		Role:        inv.Role,
		TokenHash:   inv.TokenHash,
		InvitedBy:   inv.InvitedBy,
		ExpiresAt:   inv.ExpiresAt,
		AcceptedAt:  inv.AcceptedAt,
		RevokedAt:   inv.RevokedAt,
		CreatedAt:   inv.CreatedAt,
	}
}

func toInvitationDTO(inv contract.InvitationView) invitationDTO {
	return invitationDTO{
		ID:          inv.ID,
		WorkspaceID: inv.WorkspaceID,
		Email:       inv.Email,
		Role:        inv.Role,
		TokenHash:   inv.TokenHash,
		InvitedBy:   inv.InvitedBy,
		ExpiresAt:   inv.ExpiresAt,
		AcceptedAt:  inv.AcceptedAt,
		RevokedAt:   inv.RevokedAt,
		CreatedAt:   inv.CreatedAt,
	}
}

type memberDTO struct {
	ID          int       `json:"id"`
	Email       string    `json:"email"`
	FullName    string    `json:"full_name"`
	Image       *string   `json:"image,omitempty"`
	Role        string    `json:"role"`
	DateJoined  time.Time `json:"date_joined"`
	WorkspaceId int       `json:"workspace_id"`
}

func toMemberDTO(m contract.MemberView) memberDTO {
	return memberDTO{
		ID:          m.ID,
		Email:       m.Email,
		FullName:    m.FullName,
		Image:       m.Image,
		Role:        m.Role,
		DateJoined:  m.DateJoined,
		WorkspaceId: m.WorkspaceId,
	}
}
