package web

import (
	"io/fs"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna"
)

func RegisterEmbedRoute(router *gin.Engine) {
	static := lumna.StaticFS("resources/assets")
	if static == nil {
		return
	}

	// Angular build is under browser/ in the embedded FS
	browserFS, err := fs.Sub(static, "browser")
	if err != nil {
		return
	}

	fileServer := http.FileServer(http.FS(browserFS))
	router.NoRoute(func(c *gin.Context) {
		path := strings.TrimPrefix(c.Request.URL.Path, "/")
		if path == "" {
			path = "index.html"
		}

		// Try to serve the requested file
		if _, err := fs.Stat(browserFS, path); err == nil {
			fileServer.ServeHTTP(c.Writer, c.Request)
			return
		}

		// SPA fallback: serve index.html for Angular routing
		c.FileFromFS("index.html", http.FS(browserFS))
	})
}
