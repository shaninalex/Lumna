package setup

import (
	"context"
	"fmt"
	"net/http"
	"net/mail"
	"strings"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/services/auth"
	"gitlab.com/shaninalex/lumna/app/web/setup/templates"
	"gorm.io/gorm"
)

type SetupData struct {
	Email           string `form:"email"`
	FirstName       string `form:"first_name"`
	LastName        string `form:"last_name"`
	Password        string `form:"password"`
	PasswordConfirm string `form:"password_confirm"`
}

func (d *SetupData) validate() map[string]string {
	errors := map[string]string{}

	if strings.TrimSpace(d.Email) == "" {
		errors["email"] = "Email is required"
	} else if _, err := mail.ParseAddress(d.Email); err != nil {
		errors["email"] = "Invalid email address"
	}

	if strings.TrimSpace(d.FirstName) == "" {
		errors["first_name"] = "First name is required"
	}

	if strings.TrimSpace(d.LastName) == "" {
		errors["last_name"] = "Last name is required"
	}

	if d.Password == "" {
		errors["password"] = "Password is required"
	}

	if d.PasswordConfirm == "" {
		errors["password_confirm"] = "Password confirmation is required"
	} else if d.Password != d.PasswordConfirm {
		errors["password_confirm"] = "Passwords do not match"
	}

	return errors
}

func handleSetupSubmit(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data SetupData
		if err := c.ShouldBind(&data); err != nil {
			templ.Handler(templates.SetupView(templates.SetupViewData{
				Errors: map[string]string{"email": err.Error()},
			})).ServeHTTP(c.Writer, c.Request)
			return
		}

		if errors := data.validate(); len(errors) > 0 {
			templ.Handler(templates.SetupView(templates.SetupViewData{
				Email:     data.Email,
				FirstName: data.FirstName,
				LastName:  data.LastName,
				Errors:    errors,
			})).ServeHTTP(c.Writer, c.Request)
			return
		}

		if err := createAdminUser(c, db, data); err != nil {
			templ.Handler(templates.SetupView(templates.SetupViewData{
				Errors: map[string]string{"general": err.Error()},
			})).ServeHTTP(c.Writer, c.Request)
			return
		}

		c.Redirect(http.StatusFound, "/")
	}
}

func createAdminUser(ctx context.Context, db *gorm.DB, data SetupData) error {
	identity := &models.Identity{
		FullName: fmt.Sprintf("%s %s", data.FirstName, data.LastName),
		Email:    data.Email,
		Active:   true,
	}

	if err := db.WithContext(ctx).Create(identity).Error; err != nil {
		return err
	}
	pwdHash, err := auth.CreatePasswordHash(data.Password)
	if err != nil {
		return err
	}

	credential := &models.Credential{
		IdentityID:   identity.ID,
		Provider:     "local",
		Email:        &identity.Email,
		PasswordHash: &pwdHash,
	}
	if err := db.WithContext(ctx).Create(credential).Error; err != nil {
		return err
	}

	return nil
}
