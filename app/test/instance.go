package test

import (
	"context"
	"path/filepath"

	"gitlab.com/shaninalex/lumna"
	"gitlab.com/shaninalex/lumna/app/bootstrap"
	"gitlab.com/shaninalex/lumna/app/platform/database"
)

// TB is the part of testing.TB that NewInstance needs. Both *testing.T and
// ginkgo's GinkgoT() satisfy it.
type TB interface {
	Helper()
	TempDir() string
	Cleanup(func())
	Fatalf(format string, args ...any)
}

func NewInstance(t TB, cfg ...MockConfig) *bootstrap.Instance {
	t.Helper()

	c := DefaultConfig()
	if len(cfg) > 0 {
		c = cfg[0]
	}
	c.SetKey("database.url", filepath.Join(t.TempDir(), "lumna.db")+"?_foreign_keys=on")

	migrations := lumna.StaticFS("resources/migrations")
	app, err := bootstrap.New(c, bootstrap.Assets{Migrations: migrations})
	if err != nil {
		t.Fatalf("bootstrap: %v", err)
	}
	t.Cleanup(func() { _ = app.Close(context.Background()) })

	m, err := database.NewMigrator(app.DB.From(context.Background()), migrations)
	if err != nil {
		t.Fatalf("migrator: %v", err)
	}
	if err := m.Up(); err != nil {
		t.Fatalf("migrate up: %v", err)
	}
	return app
}
