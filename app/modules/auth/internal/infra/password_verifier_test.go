package infra_test

import (
	"testing"

	"gitlab.com/shaninalex/lumna/app/core/secure"
	"gitlab.com/shaninalex/lumna/app/modules/auth/internal/infra"
)

func TestPasswordVerifier(t *testing.T) {
	v := infra.NewPasswordVerifier()

	testData := map[string]struct {
		pwd   secure.Secret
		score int
	}{
		"low":     {secure.Secret("password"), 1},
		"low2":    {secure.Secret("passwordpassword"), 2},
		"normal":  {secure.Secret("PassWord123"), 3},
		"normal2": {secure.Secret("PassWord123__"), 4},
		"high":    {secure.Secret("PassWord123__PassWord123__"), 5},
	}

	for name, tc := range testData {
		t.Run(name, func(t *testing.T) {
			if r := v.Verify(tc.pwd); r != tc.score {
				t.Errorf("Validate(%s) score = %d, wantErr %d", tc.pwd.Reveal(), r, tc.score)
			}
		})
	}

}
