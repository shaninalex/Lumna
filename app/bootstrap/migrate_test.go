package bootstrap_test

import (
	"context"
	"os"
	"path/filepath"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"gitlab.com/shaninalex/lumna/app/platform/config"
	"gitlab.com/shaninalex/lumna/app/platform/database"
)

const initVersion = 20261116005709

var migrations = os.DirFS("../../resources/migrations")

func openEmptyDB() *database.DB {
	dir := GinkgoT().TempDir()
	cfgPath := filepath.Join(dir, "config.yaml")
	body := "env: testing\ndatabase:\n  url: \"" + filepath.Join(dir, "lumna.db") + "\"\n"
	Expect(os.WriteFile(cfgPath, []byte(body), 0o600)).To(Succeed())

	db := database.New(config.ReadConfig(cfgPath))
	DeferCleanup(db.Close)
	return db
}

func tables(db *database.DB) []string {
	var names []string
	Expect(db.From(context.Background()).
		Raw("SELECT name FROM sqlite_master WHERE type = 'table' AND name NOT LIKE 'sqlite_%' ORDER BY name").
		Scan(&names).Error).To(Succeed())
	return names
}

var _ = Describe("Migrator", func() {
	var (
		db       *database.DB
		migrator *database.Migrator
	)

	BeforeEach(func() {
		db = openEmptyDB()
		var err error
		migrator, err = database.NewMigrator(db.From(context.Background()), migrations)
		Expect(err).NotTo(HaveOccurred())
	})

	expectVersion := func(version uint, dirty bool) {
		GinkgoHelper()
		v, d, err := migrator.Version()
		Expect(err).NotTo(HaveOccurred())
		Expect(v).To(Equal(version))
		Expect(d).To(Equal(dirty))
	}

	It("reports no version on an empty database", func() {
		expectVersion(0, false)
	})

	It("applies every migration to an empty database", func() {
		Expect(migrator.Up()).To(Succeed())

		expectVersion(initVersion, false)
		Expect(tables(db)).To(ContainElements(
			"identities", "credentials", "refresh_tokens",
			"workspaces", "identity_workspaces", "projects", "workspace_invitations",
			"scopes", "stages", "work_items", "work_items_assignees",
			"journal", "notifications", "notification_recipients", "email_queue",
		))
	})

	It("treats a second run as no change", func() {
		Expect(migrator.Up()).To(Succeed())
		Expect(migrator.Up()).To(Succeed())
		expectVersion(initVersion, false)
	})

	It("rolls everything back and can apply again", func() {
		Expect(migrator.Up()).To(Succeed())

		Expect(migrator.Down(0)).To(Succeed())
		expectVersion(0, false)
		Expect(tables(db)).To(ConsistOf("schema_migrations"), "down must drop every table up creates")

		Expect(migrator.Up()).To(Succeed())
		expectVersion(initVersion, false)
	})

	It("rolls back a single step", func() {
		Expect(migrator.Up()).To(Succeed())
		Expect(migrator.Down(1)).To(Succeed())
		expectVersion(0, false)
	})

	It("rolls back everything when asked for more steps than were applied", func() {
		Expect(migrator.Up()).To(Succeed())
		Expect(migrator.Down(5)).To(Succeed())
		expectVersion(0, false)
	})

	It("treats rolling back an empty database as no change", func() {
		Expect(migrator.Down(0)).To(Succeed())
		Expect(migrator.Down(1)).To(Succeed())
		expectVersion(0, false)
	})

	It("refuses to migrate a dirty database until the version is forced", func() {
		Expect(migrator.Up()).To(Succeed())
		Expect(db.From(context.Background()).Exec("UPDATE schema_migrations SET dirty = 1").Error).To(Succeed())
		expectVersion(initVersion, true)

		Expect(migrator.Up()).To(MatchError(ContainSubstring("Dirty database")))

		Expect(migrator.Force(initVersion)).To(Succeed())
		expectVersion(initVersion, false)
		Expect(migrator.Up()).To(Succeed())
	})
})
