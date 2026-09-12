package contract

type ProfileView struct {
	ID       int
	Email    string
	FullName string
	Active   bool
}

type ListProfilesView struct {
	Profiles []ProfileView
	Limit    int
	Offset   int
}
