package auth_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"gitlab.com/shaninalex/lumna/app/adapters/api/handler/auth"
	"gitlab.com/shaninalex/lumna/app/adapters/httpx"
	"gitlab.com/shaninalex/lumna/app/test"
)

func TestControllers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Controller Auth Suite")
}

var _ = Describe("Login endpoint", func() {
	var url string

	BeforeEach(func() {
		app := test.NewInstance(GinkgoT())

		router := gin.Default()
		router.POST("/", auth.HandleLogin(test.Resolve(app.Core), httpx.CookieConfig{}))

		ts := httptest.NewServer(router)
		DeferCleanup(ts.Close)
		url = ts.URL
	})

	DescribeTable("rejects invalid credentials",
		func(data map[string]string, status int, substr string) {
			b, _ := json.Marshal(data)
			res, err := http.Post(url, "application/json", bytes.NewReader(b))
			Expect(err).NotTo(HaveOccurred())
			DeferCleanup(res.Body.Close)
			Expect(res).To(HaveHTTPStatus(status))
			var bodyBytes, _ = io.ReadAll(res.Body)
			Expect(string(bodyBytes)).Should(ContainSubstring(substr))
		},
		Entry("empty credentials", map[string]string{"email": "", "password": ""}, http.StatusUnauthorized, "\"email\":\"cannot be blank\""),
		Entry("invalid email", map[string]string{"email": "invalid email", "password": ""}, http.StatusUnauthorized, "\"email\":\"must be a valid email address\""),
		Entry("empty pwd", map[string]string{"email": "valid@email.com", "password": ""}, http.StatusUnauthorized, "\"password\":\"cannot be blank\""),
		Entry("unknown user", map[string]string{"email": "valid@email.com", "password": "123"}, http.StatusUnauthorized, "\"message\":\"invalid credentials\""),
	)

	DescribeTable("accept valid credentials",
		func(data map[string]string, status int, substr string) {
			b, _ := json.Marshal(data)
			res, err := http.Post(url, "application/json", bytes.NewReader(b))
			Expect(err).NotTo(HaveOccurred())
			DeferCleanup(res.Body.Close)
			var bodyBytes, _ = io.ReadAll(res.Body)
			Expect(string(bodyBytes)).Should(ContainSubstring(substr))
			cookies := res.Cookies()
			Expect(cookies).Should(HaveLen(2))

			authCookies := []string{httpx.CookieAccessToken, httpx.CookieRefreshToken}
			for _, cookie := range cookies {
				Expect(cookie.Name).Should(BeElementOf(authCookies))
			}
		},
		Entry("valid user", map[string]string{"email": "test@test.com", "password": "password"}, http.StatusOK, "Login successful"),
	)
})

var _ = Describe("Logout endpoint", func() {
	var (
		base   string
		client *http.Client
	)

	BeforeEach(func() {
		app := test.NewInstance(GinkgoT())
		router := gin.New()
		auth.Register(test.Resolve(app.Core), httpx.CookieConfig{}, router.Group("/api/v1/auth"))
		ts := httptest.NewServer(router)
		DeferCleanup(ts.Close)
		base = ts.URL + "/api/v1/auth"
		jar, _ := cookiejar.New(nil)
		client = &http.Client{Jar: jar}
	})

	login := func() *http.Cookie {
		b, _ := json.Marshal(map[string]string{"email": "test@test.com", "password": "password"})
		res, err := client.Post(base+"/login", "application/json", bytes.NewReader(b))
		Expect(err).NotTo(HaveOccurred())
		DeferCleanup(res.Body.Close)
		Expect(res).To(HaveHTTPStatus(http.StatusOK))
		for _, c := range res.Cookies() {
			if c.Name == httpx.CookieRefreshToken {
				return c
			}
		}
		Fail("No refresh token cookie found after login")
		return nil
	}

	It("clears cookies and revokes the refresh token", func() {
		refresh := login()

		res, err := client.Post(base+"/logout", "application/json", nil)
		Expect(err).NotTo(HaveOccurred())
		DeferCleanup(res.Body.Close)
		Expect(res).To(HaveHTTPStatus(http.StatusOK))

		// 1. both cookies expired
		cleared := map[string]bool{}
		for _, c := range res.Cookies() {
			Expect(c.MaxAge).To(BeNumerically("<", 0), c.Name)
			cleared[c.Name] = true
		}
		Expect(cleared).To(HaveKey(httpx.CookieAccessToken))
		Expect(cleared).To(HaveKey(httpx.CookieRefreshToken))

		// 2. old refresh token no longer works (jar is empty now, so send it by hand)
		req, _ := http.NewRequest(http.MethodPost, base+"/refresh", nil)
		req.AddCookie(&http.Cookie{Name: refresh.Name, Value: refresh.Value})
		res2, err := http.DefaultClient.Do(req)
		Expect(err).NotTo(HaveOccurred())
		DeferCleanup(res2.Body.Close)
		Expect(res2).To(HaveHTTPStatus(http.StatusUnauthorized))
	})

	It("is idempotent without a session", func() {
		res, err := client.Post(base+"/logout", "application/json", nil)
		Expect(err).NotTo(HaveOccurred())
		DeferCleanup(res.Body.Close)
		Expect(res).To(HaveHTTPStatus(http.StatusOK))
	})
})
