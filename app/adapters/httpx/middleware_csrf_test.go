package httpx_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/adapters/httpx"
)

func newProtectedEngine(t *testing.T) *gin.Engine {
	t.Helper()
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(httpx.CrossOriginMiddleware())
	ok := func(c *gin.Context) { c.Status(http.StatusOK) }
	r.GET("/api/v1/thing", ok)
	r.POST("/api/v1/thing", ok)
	r.DELETE("/api/v1/thing", ok)
	r.POST("/setup", ok)
	return r
}

func TestCrossOriginMiddleware(t *testing.T) {
	const host = "lumna.example.com"

	tests := []struct {
		name        string
		method      string
		path        string
		contentType string
		body        string
		headers     map[string]string
		wantStatus  int
	}{
		// Origin / Sec-Fetch-Site
		{
			name: "foreign Origin is rejected", method: http.MethodPost, path: "/api/v1/thing",
			contentType: "application/json", body: `{}`,
			headers: map[string]string{"Origin": "https://evil.com"}, wantStatus: http.StatusForbidden,
		},
		{
			name: "cross-site Sec-Fetch-Site is rejected", method: http.MethodPost, path: "/api/v1/thing",
			contentType: "application/json", body: `{}`,
			headers: map[string]string{"Sec-Fetch-Site": "cross-site", "Origin": "https://evil.com"}, wantStatus: http.StatusForbidden,
		},
		{
			name: "cross-site form to /setup is rejected", method: http.MethodPost, path: "/setup",
			contentType: "application/x-www-form-urlencoded", body: `a=b`,
			headers: map[string]string{"Sec-Fetch-Site": "cross-site", "Origin": "https://evil.com"}, wantStatus: http.StatusForbidden,
		},
		{
			name: "foreign Origin DELETE is rejected", method: http.MethodDelete, path: "/api/v1/thing",
			headers: map[string]string{"Origin": "https://evil.com"}, wantStatus: http.StatusForbidden,
		},
		{
			name: "same-origin form to /setup passes", method: http.MethodPost, path: "/setup",
			contentType: "application/x-www-form-urlencoded", body: `a=b`,
			headers: map[string]string{"Sec-Fetch-Site": "same-origin", "Origin": "https://" + host}, wantStatus: http.StatusOK,
		},
		{
			name: "same-origin Sec-Fetch-Site passes", method: http.MethodPost, path: "/api/v1/thing",
			contentType: "application/json", body: `{}`,
			headers: map[string]string{"Sec-Fetch-Site": "same-origin", "Origin": "https://" + host}, wantStatus: http.StatusOK,
		},
		{
			name: "Origin matching Host passes without Sec-Fetch-Site", method: http.MethodPost, path: "/api/v1/thing",
			contentType: "application/json", body: `{}`,
			headers: map[string]string{"Origin": "https://" + host}, wantStatus: http.StatusOK,
		},
		{
			name: "non-browser client without Origin passes", method: http.MethodPost, path: "/api/v1/thing",
			contentType: "application/json", body: `{}`, wantStatus: http.StatusOK,
		},
		{
			name: "same-site subdomain is rejected", method: http.MethodPost, path: "/api/v1/thing",
			contentType: "application/json", body: `{}`,
			headers: map[string]string{"Sec-Fetch-Site": "same-site", "Origin": "https://app.example.com"}, wantStatus: http.StatusForbidden,
		},
		{
			name: "cross-origin GET passes", method: http.MethodGet, path: "/api/v1/thing",
			headers: map[string]string{"Sec-Fetch-Site": "cross-site", "Origin": "https://evil.com"}, wantStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := newProtectedEngine(t)
			req := httptest.NewRequest(tt.method, tt.path, strings.NewReader(tt.body))
			req.Host = host
			if tt.contentType != "" {
				req.Header.Set("Content-Type", tt.contentType)
			}
			for k, v := range tt.headers {
				req.Header.Set(k, v)
			}
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)
			if w.Code != tt.wantStatus {
				t.Fatalf("status = %d, want %d; body: %s", w.Code, tt.wantStatus, w.Body.String())
			}
		})
	}
}
