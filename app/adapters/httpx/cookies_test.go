package httpx_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/adapters/httpx"
)

func sessionCookies(t *testing.T, cfg httpx.CookieConfig, clear bool) map[string]*http.Cookie {
	t.Helper()
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest(http.MethodPost, "/api/v1/auth/login", nil)
	if clear {
		httpx.ClearSession(c, cfg)
	} else {
		httpx.SetSession(c, cfg, "access", time.Minute, "refresh", time.Hour)
	}

	cookies := map[string]*http.Cookie{}
	for _, ck := range w.Result().Cookies() {
		cookies[ck.Name] = ck
	}
	return cookies
}

func TestSessionCookiesSameSite(t *testing.T) {
	tests := []struct {
		name        string
		cfg         httpx.CookieConfig
		wantAccess  http.SameSite
		wantRefresh http.SameSite
	}{
		{
			name:       "production over https (issue #16)",
			cfg:        httpx.CookieConfig{Secure: true},
			wantAccess: http.SameSiteLaxMode, wantRefresh: http.SameSiteStrictMode,
		},
		{
			name:       "local development over http",
			cfg:        httpx.CookieConfig{Secure: false},
			wantAccess: http.SameSiteLaxMode, wantRefresh: http.SameSiteStrictMode,
		},
	}

	for _, tt := range tests {
		for _, clear := range []bool{false, true} {
			t.Run(tt.name, func(t *testing.T) {
				cookies := sessionCookies(t, tt.cfg, clear)
				access, refresh := cookies[httpx.CookieAccessToken], cookies[httpx.CookieRefreshToken]
				if access == nil || refresh == nil {
					t.Fatalf("missing cookies: %v", cookies)
				}
				if access.SameSite != tt.wantAccess {
					t.Errorf("access_token SameSite = %v, want %v", access.SameSite, tt.wantAccess)
				}
				if refresh.SameSite != tt.wantRefresh {
					t.Errorf("refresh_token SameSite = %v, want %v", refresh.SameSite, tt.wantRefresh)
				}
				if access.Secure != tt.cfg.Secure || refresh.Secure != tt.cfg.Secure {
					t.Errorf("Secure flag does not follow config")
				}
			})
		}
	}
}
