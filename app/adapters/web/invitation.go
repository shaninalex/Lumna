package web

import (
	"errors"
	"log"
	"net/http"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/adapters/httpx"
	"gitlab.com/shaninalex/lumna/app/adapters/web/templates"
	"gitlab.com/shaninalex/lumna/app/core"
	"gitlab.com/shaninalex/lumna/app/core/errs"
	"gitlab.com/shaninalex/lumna/app/core/secure"
	authContract "gitlab.com/shaninalex/lumna/app/modules/auth/contract"
	identityContract "gitlab.com/shaninalex/lumna/app/modules/identity/contract"
	"gitlab.com/shaninalex/lumna/app/modules/workspace/contract"
)

var ConfirmInvitationEmptyKey = errors.New("invalid invitation key")
var InvitationNotFoundError = errors.New("invitation not found")

// InvitationCookieAccessToken holds the invitation session token. It is
// scoped to /invitation, so it never reaches the API or the SPA.
var InvitationCookieAccessToken = "invitation_access_token"

func handleConfirmInvitation(app core.Resolve, cookieConfig httpx.CookieConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.Query("key")
		if key == "" {
			c.Data(http.StatusBadRequest, "text/html; charset=utf-8", []byte(ConfirmInvitationEmptyKey.Error()))
			return
		}

		// Validation
		invitation, err := contract.AskInvitationValidate(c.Request.Context(), app(), contract.InvitationValidate{
			Hash: key,
		})
		if err != nil {
			httpx.Fail(c, errors.Join(ConfirmInvitationEmptyKey, err))
			return
		}

		// check is user already exists by invitation.Email
		existed, err := identityContract.AskGetProfileByEmail(c.Request.Context(), app(), identityContract.GetProfileByEmail{Email: invitation.Email})
		if err != nil && errs.KindOf(err) != errs.KindNotFound {
			httpx.Fail(c, err)
			return
		}

		if err == nil && invitation.Email == existed.Email {
			_, err = contract.ExecAcceptInvitation(c.Request.Context(), app(), contract.AcceptInvitation{
				IdentityId:   existed.ID,
				InvitationId: invitation.ID,
			})
			if err != nil {
				httpx.Fail(c, err)
				return
			}
			c.Redirect(http.StatusFound, "/auth/login")
			return
		}

		// open an invitation session: the token carries the invitation
		// validated above, so later steps don't need the key again
		session, err := authContract.ExecInvitationSession(c.Request.Context(), app(), authContract.InvitationSession{
			Email:        invitation.Email,
			Role:         invitation.Role,
			InvitationID: invitation.ID,
			WorkspaceId:  invitation.WorkspaceID,
		})
		if err != nil {
			httpx.Fail(c, errors.Join(ConfirmInvitationEmptyKey, err))
			return
		}

		c.SetCookie(
			InvitationCookieAccessToken,
			session.AccessToken,
			int(session.AccessTokenDuration.Seconds()),
			"/invitation", "",
			cookieConfig.Secure,
			true,
		)

		// redirect to /invitation/password
		c.Redirect(http.StatusFound, "/invitation/password")
	}
}

// invitationClaimsKey is the gin context key the verified invitation is
// stored under by invitationSessionMiddleware.
const invitationClaimsKey = "invitation_claims"

// invitationSessionMiddleware lets a request through only with a valid
// invitation session cookie, and stores the invitation it carries for the
// handlers (see invitationFrom). A regular access token is not accepted here.
func invitationSessionMiddleware(verifier authContract.Verifier[authContract.InvitationClaims]) gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken, err := c.Cookie(InvitationCookieAccessToken)
		if err != nil {
			log.Println(err)
			c.Data(http.StatusForbidden, "text/html; charset=utf-8", []byte(ConfirmInvitationEmptyKey.Error()))
			c.Abort()
			return
		}

		claims, err := verifier.Verify(c.Request.Context(), accessToken)
		if err != nil {
			log.Println(err)
			c.Data(http.StatusForbidden, "text/html; charset=utf-8", []byte(ConfirmInvitationEmptyKey.Error()))
			c.Abort()
			return
		}

		c.Set(invitationClaimsKey, claims)
		c.Next()
	}
}

// invitationFrom returns the invitation of the current session. It is only
// meaningful behind invitationSessionMiddleware.
func invitationFrom(c *gin.Context) (authContract.InvitationClaims, bool) {
	v, ok := c.Get(invitationClaimsKey)
	if !ok {
		return authContract.InvitationClaims{}, false
	}
	claims, ok := v.(authContract.InvitationClaims)
	return claims, ok
}

func handleCreatePassword() gin.HandlerFunc {
	return func(c *gin.Context) {
		invitation, ok := invitationFrom(c)
		if !ok {
			c.Data(http.StatusNotFound, "text/html; charset=utf-8", []byte(InvitationNotFoundError.Error()))
			c.Abort()
			return
		}
		templ.Handler(templates.SetPasswordPage(invitation, nil)).ServeHTTP(c.Writer, c.Request)
	}
}

var (
	PasswordConfirmDoesNotMatch = errors.New("password confirm does not match")
	PasswordTooWeak             = errors.New("password is too weak")
	PasswordNameIsEmpty         = errors.New("password name is empty")
	FullNameIsEmpty             = errors.New("full name is empty")
	UnableToAddMember           = errors.New("unable to add member")
)

func handleAcceptInvitation(app core.Resolve, verifier authContract.PasswordVerifier, cookieConfig httpx.CookieConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		invitation, ok := invitationFrom(c)
		if !ok {
			c.Data(http.StatusNotFound, "text/html; charset=utf-8", []byte(InvitationNotFoundError.Error()))
			c.Abort()
			return
		}

		// validate payload (password/password_confirm)
		fullName, securePassword, err := parseInvitationProcessFormData(c, verifier)
		if err != nil {
			templ.Handler(templates.SetPasswordPage(invitation, err)).ServeHTTP(c.Writer, c.Request)
			c.Abort()
			return
		}

		// Create user (workspace validates the role)
		profile, err := identityContract.ExecRegister(c.Request.Context(), app(), identityContract.Register{
			Email:    invitation.Email,
			FullName: fullName,
			Password: securePassword,
		})
		if err != nil {
			c.Data(http.StatusNotFound, "text/html; charset=utf-8", []byte(err.Error()))
			c.Abort()
			return
		}

		// add new user to workspace, accept invitation and push event
		_, err = contract.ExecAcceptInvitation(c.Request.Context(), app(), contract.AcceptInvitation{
			IdentityId:   profile.ID,
			InvitationId: invitation.InvitationID,
		})
		if err != nil {
			templ.Handler(templates.SetPasswordPage(invitation, err)).ServeHTTP(c.Writer, c.Request)
			c.Abort()
			return
		}

		// remove temp session cookie
		c.SetCookie(InvitationCookieAccessToken, "", -1, "/invitation", "", cookieConfig.Secure, true)

		// redirect to /auth/login
		c.Redirect(http.StatusFound, "/auth/login")
	}
}

func parseInvitationProcessFormData(c *gin.Context, verifier authContract.PasswordVerifier) (string, secure.Secret, error) {
	fullName := c.PostForm("full_name")
	if fullName == "" {
		return "", "", FullNameIsEmpty
	}
	password := c.PostForm("password")
	if password == "" {
		return "", "", PasswordNameIsEmpty
	}
	passwordConfirm := c.PostForm("password_confirm")
	if password != passwordConfirm {
		return "", "", PasswordConfirmDoesNotMatch
	}

	securePassword := secure.Secret(password)
	score := verifier.Verify(secure.Secret(password))
	if score < 4 {
		return "", "", PasswordTooWeak
	}

	return fullName, securePassword, nil
}
