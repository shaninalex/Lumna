package utils

import (
	"fmt"
	"math/rand/v2"
	"regexp"
	"strings"

	"github.com/google/uuid"
)

func CreateEntitySlug(s string) string {
	return fmt.Sprintf("%d-%s", randRange(100000, 999999), Slugify(s))
}

func Slugify(s string) string {
	res := strings.ToLower(strings.TrimSpace(s))
	res = strings.Replace(res, " ", "-", -1)
	regex, err := regexp.Compile("[^a-z0-9-]+")
	if err != nil {
		// fall back, always return string
		return uuid.New().String()
	}
	return regex.ReplaceAllString(res, "")
}

func randRange(min, max int) int {
	return rand.IntN(max-min) + min
}

type TitleContainer interface {
	GetTitle() string
}

func CreateSlug(s TitleContainer) string {
	return fmt.Sprintf("%s-%s", RandomCode(6), Slugify(s.GetTitle()))
}
