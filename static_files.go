package lumna

import (
	"embed"
	"io/fs"
	"log"
)

//go:embed all:resources/frontend_build
//go:embed all:resources/migrations
//go:embed all:resources/openapi
var resourcesFS embed.FS

// StaticFS returns the embedded static files rooted at p.
//
// Resources are always compiled into the binary, so nothing here depends on the
// working directory. The frontend under resources/assets is a build artifact and
// is not committed: run the frontend build before `go build`, otherwise the
// binary embeds only the .gitkeep placeholder and serves no UI.
func StaticFS(p string) fs.FS {
	sub, err := fs.Sub(resourcesFS, p)
	if err != nil {
		log.Printf("warning: failed to load embedded static files: %v", err)
		return nil
	}

	return sub
}
