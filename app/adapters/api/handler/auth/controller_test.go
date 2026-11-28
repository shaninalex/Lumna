package auth_test

import (
	"bytes"
	"encoding/json"
	"net/http"
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
		func(data map[string]string, status int) {
			b, _ := json.Marshal(data)
			res, err := http.Post(url, "application/json", bytes.NewReader(b))
			Expect(err).NotTo(HaveOccurred())
			DeferCleanup(res.Body.Close)
			Expect(res).To(HaveHTTPStatus(status))
		},
		Entry("empty credentials", map[string]string{"email": "", "password": ""}, http.StatusUnauthorized),
		Entry("invalid email", map[string]string{"email": "invalid email", "password": ""}, http.StatusUnauthorized),
		Entry("empty pwd", map[string]string{"email": "valid@email.com", "password": ""}, http.StatusUnauthorized),
		Entry("unknown user", map[string]string{"email": "valid@email.com", "password": "123"}, http.StatusUnauthorized),
	)
})
