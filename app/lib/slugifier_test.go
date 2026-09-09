package utils_test

import (
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/shaninalex/lumna/app/pkg/utils"
)

func Test_Slugify(t *testing.T) {
	s := utils.Slugify("This Is Example Text")
	assert.Equal(t, "this-is-example-text", s)
}

func Test_SlugifySpecialCharacters(t *testing.T) {
	s := utils.Slugify("This Is Example ^ )@$)&@)(*&$ Text")
	assert.Equal(t, "this-is-example---text", s)
}

func Test_CreateEntitySlug(t *testing.T) {
	s := utils.CreateEntitySlug("New task")
	assert.NotEmpty(t, s)

	parts := strings.Split(s, "-")
	assert.Len(t, parts, 3)

	// digit part
	_, err := strconv.Atoi(parts[0])
	assert.Nil(t, err)
	assert.Equal(
		t,
		"new-task",
		strings.Join([]string{parts[1], parts[2]}, "-"),
	)
}
