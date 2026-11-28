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
	"gitlab.com/shaninalex/lumna/app/bootstrap"
	"gitlab.com/shaninalex/lumna/app/core"
	"gitlab.com/shaninalex/lumna/app/test"
)

func TestControllers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Controller Auth Suite")
}

var _ = Describe("Login endpoint", func() {
	var resolve core.Resolve
	config := test.DefaultConfig()

	BeforeEach(func() {
		app, err := bootstrap.New(&config, bootstrap.Assets{})
		if err != nil {
			panic(err)
		}
		resolve = func() *core.App { return app.Core }
	})

	Describe("payload invalid", func() {
		It("returns 400", func() {
			// move to golang http router or chi router
			router := gin.Default()
			router.POST("/", auth.HandleLogin(resolve, httpx.CookieConfig{}))

			ts := httptest.NewServer(router)
			DeferCleanup(ts.Close)

			data := map[string]string{
				"email":    "",
				"password": "",
			}

			b, _ := json.Marshal(data)
			res, err := http.Post(ts.URL, "application/json", bytes.NewReader(b))
			Expect(err).NotTo(HaveOccurred())
			Expect(res).To(HaveHTTPStatus(http.StatusBadRequest))
		})
	})

})
