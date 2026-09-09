package utils_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/shaninalex/lumna/app/pkg/utils"
)

func Test_ProjectKey(t *testing.T) {
	s := utils.ProjectKey("project title")
	assert.Equal(t, "PT", s)

	s = utils.ProjectKey("")
	assert.Equal(t, "", s)

	s = utils.ProjectKey("a")
	assert.Equal(t, "A", s)

	s = utils.ProjectKey("ab")
	assert.Equal(t, "AB", s)

	s = utils.ProjectKey("abc")
	assert.Equal(t, "AB", s)

	s = utils.ProjectKey("abcd")
	assert.Equal(t, "AB", s)

	s = utils.ProjectKey("project")
	assert.Equal(t, "PR", s)

	s = utils.ProjectKey("123")
	assert.Equal(t, "12", s)

	s = utils.ProjectKey("Project 123")
	assert.Equal(t, "P1", s)
}
