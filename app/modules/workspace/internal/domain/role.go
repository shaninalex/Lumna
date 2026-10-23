package domain

const (
	RoleAdmin  = "admin"
	RoleMember = "member"
)

// ValidateRole reports whether role is one this module knows.
func ValidateRole(role string) error {
	switch role {
	case RoleAdmin, RoleMember:
		return nil
	}
	return ErrUnknownRole
}
