package main

import (
	"os"

	"gitlab.com/shaninalex/lumna"
	"gitlab.com/shaninalex/lumna/app/bootstrap"
)

func main() {
	os.Exit(bootstrap.RunCLI(bootstrap.Assets{
		Migrations: lumna.StaticFS("resources/migrations"),
		OpenAPI:    lumna.StaticFS("resources/openapi"),
		SPA:        lumna.StaticFS("resources/assets"),
		Version:    lumna.Version,
	}))
}
