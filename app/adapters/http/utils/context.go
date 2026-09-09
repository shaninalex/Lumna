package utils

import (
	"context"
	"errors"

	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/pkg"
)

var (
	ErrorUserIDNotInContext   error = errors.New("user id not found in request context")
	ErrorInvalidUserID        error = errors.New("invalid user id")
	ErrorIdentityNotInContext error = errors.New("identity not found in context")
)

func GetUserID(c *gin.Context) (int, error) {
	userIDAny, ok := c.Get(pkg.ContextUserID)
	if !ok {
		return 0, ErrorUserIDNotInContext
	}

	userID, ok := userIDAny.(int)
	if !ok {
		return 0, ErrorInvalidUserID
	}
	return userID, nil
}

func GetIdentity(c context.Context) (*models.Identity, error) {
	identity, ok := c.Value(pkg.ContextIdentity).(*models.Identity)
	if !ok {
		return nil, ErrorIdentityNotInContext
	}

	return identity, nil
}
