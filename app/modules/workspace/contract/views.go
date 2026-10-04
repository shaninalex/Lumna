package contract

import "time"

type WorkspaceView struct {
	Id         int
	Title      string
	OwnerEmail string
	Active     bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type ProjectView struct {
	Id          int
	Title       string
	WorkspaceId int
	OwnerId     int
	Meta        string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type InvitationView struct {
	ID          int
	WorkspaceID int
	Email       string
	Role        string
	InvitedBy   int
	ExpiresAt   time.Time
	AcceptedAt  time.Time
	RevokedAt   time.Time
	CreatedAt   time.Time
}

type MemberView struct {
	ID          int
	Email       string
	FullName    string
	Image       string
	Role        string
	DateJoined  time.Time
	WorkspaceId int
}
