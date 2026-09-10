package contract

import "time"

type ProfileView struct {
	ID       int
	Email    string
	FullName string
	Active   bool
}

type SessionView struct {
	IdentityID   int
	AccessToken  string
	RefreshToken string
	ExpiresAt    time.Time
}

type ListProfilesView struct {
	Profiles []ProfileView
	Limit    int
	Offset   int
}
