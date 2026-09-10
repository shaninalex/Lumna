package setup

import (
	"fmt"
	"net/http"
	"net/mail"
	"strings"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/adapters/webui/setup/templates"
	"gitlab.com/shaninalex/lumna/app/core"
	"gitlab.com/shaninalex/lumna/app/core/bus"
	"gitlab.com/shaninalex/lumna/app/modules/identity/contract"
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

func handleSetupSubmit(app core.Resolve) gin.HandlerFunc {
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

		_, err := contract.ExecRegister(c.Request.Context(), app(), contract.Register{
			Email:    data.Email,
			Password: bus.Secret(data.Password),
			FullName: fmt.Sprintf("%s %s", data.FirstName, data.LastName),
		})
		if err != nil {
			templ.Handler(templates.SetupView(templates.SetupViewData{
				Errors: map[string]string{"general": err.Error()},
			})).ServeHTTP(c.Writer, c.Request)
			return
		}

		c.Redirect(http.StatusFound, "/")
	}
}
