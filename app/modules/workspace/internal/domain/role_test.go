package domain_test

import (
	"errors"
	"testing"

	"gitlab.com/shaninalex/lumna/app/modules/workspace/internal/domain"
)

func TestValidateRole(t *testing.T) {
	for _, role := range []string{domain.RoleAdmin, domain.RoleMember} {
		if err := domain.ValidateRole(role); err != nil {
			t.Errorf("ValidateRole(%q) = %v, want nil", role, err)
		}
	}
	for _, role := range []string{"", "owner", "Admin"} {
		if err := domain.ValidateRole(role); !errors.Is(err, domain.ErrUnknownRole) {
			t.Errorf("ValidateRole(%q) = %v, want %v", role, err, domain.ErrUnknownRole)
		}
	}
}
