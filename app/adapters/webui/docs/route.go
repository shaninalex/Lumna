package docs

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna"
)

// RegisterDocsRoute serves the embedded OpenAPI spec and its Swagger UI under /docs.
//
// The UI has to be same-origin with the API: it calls the API with credentials so
// that the real session cookies are attached, and CORSMiddleware answers with
// Access-Control-Allow-Credentials: false, which rules out any other origin.
func RegisterDocsRoute(router *gin.Engine) {
	docsFS := lumna.StaticFS("resources/openapi")
	if docsFS == nil {
		return
	}

	router.StaticFS("/docs", http.FS(docsFS))
}
