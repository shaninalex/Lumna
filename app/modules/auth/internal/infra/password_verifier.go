package infra

import (
	"regexp"

	"gitlab.com/shaninalex/lumna/app/core/secure"
)

var (
	haveLowercase = regexp.MustCompile(`[a-z]`)
	haveUppercase = regexp.MustCompile(`[A-Z]`)
	haveDigit     = regexp.MustCompile(`[0-9]`)
	haveSpecial   = regexp.MustCompile(`[^a-zA-Z0-9]`)
)

type PasswordVerifier struct{}

func NewPasswordVerifier() *PasswordVerifier {
	return &PasswordVerifier{}
}

func (s *PasswordVerifier) Verify(pwd secure.Secret) int {
	password := pwd.Reveal()
	score := 0

	// At least 8 bytes.
	if len(password) >= 8 {
		score++
	}

	// Contains lowercase and uppercase letters.
	if haveLowercase.MatchString(password) &&
		haveUppercase.MatchString(password) {
		score++
	}

	// Contains a digit.
	if haveDigit.MatchString(password) {
		score++
	}

	// Contains a special character.
	if haveSpecial.MatchString(password) {
		score++
	}

	// At least 16 bytes.
	if len(password) >= 16 {
		score++
	}

	return score
}
