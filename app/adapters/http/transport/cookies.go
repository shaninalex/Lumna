package transport

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	CookieAccessToken  = "access_token"
	CookieRefreshToken = "refresh_token"

	// The refresh token is only ever presented to the auth endpoints, so the
	// browser has no reason to attach it to every other request.
	refreshCookiePath = "/http/v1/auth"
)

// CookieConfig carries the deployment-dependent half of cookie handling.
type CookieConfig struct {
	Secure bool
}

func (cfg CookieConfig) sameSite() http.SameSite {
	// A cross-origin SPA can only send cookies with SameSite=None, and
	// browsers accept None only together with Secure.
	if cfg.Secure {
		return http.SameSiteNoneMode
	}
	return http.SameSiteLaxMode
}

// SetSession writes the session cookies. HttpOnly throughout: no script ever
// needs to read these, and not exposing them removes a whole class of XSS
// token theft.
func SetSession(c *gin.Context, cfg CookieConfig, accessToken string, accessTTL time.Duration, refreshToken string, refreshTTL time.Duration) {
	c.SetSameSite(cfg.sameSite())
	c.SetCookie(CookieAccessToken, accessToken, int(accessTTL.Seconds()), "/", "", cfg.Secure, true)
	c.SetCookie(CookieRefreshToken, refreshToken, int(refreshTTL.Seconds()), refreshCookiePath, "", cfg.Secure, true)
}

// ClearSession expires both cookies. Paths must match SetSession exactly,
// otherwise the browser keeps the originals.
func ClearSession(c *gin.Context, cfg CookieConfig) {
	c.SetSameSite(cfg.sameSite())
	c.SetCookie(CookieAccessToken, "", -1, "/", "", cfg.Secure, true)
	c.SetCookie(CookieRefreshToken, "", -1, refreshCookiePath, "", cfg.Secure, true)
}
