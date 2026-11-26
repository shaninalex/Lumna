package database_test

import (
	"context"
	"database/sql"
	"log"
	"testing"

	"gitlab.com/shaninalex/lumna/app/platform/database"
	"gitlab.com/shaninalex/lumna/app/test"
)

func TestDatabaseForeignKeysON(t *testing.T) {
	ctx := context.Background()
	c := test.NewMockConfig(map[string]any{"database.url": "file::memory:?cache=shared&_foreign_keys=on"})
	db := database.New(&c)
	t.Cleanup(func() { _ = db.Close() })

	sqlDB, err := db.From(ctx).DB()
	if err != nil {
		t.Fatal(err)
	}

	var n = 4
	sqlDB.SetMaxOpenConns(n)

	conns := make([]*sql.Conn, 0, n)
	t.Cleanup(func() {
		for _, c := range conns {
			_ = c.Close()
		}
	})

	for range n {
		conn, err := sqlDB.Conn(ctx)
		if err != nil {
			t.Fatal(err)
		}
		conns = append(conns, conn)
	}

	for i, conn := range conns {
		var on int
		if err := conn.QueryRowContext(ctx, "PRAGMA foreign_keys").Scan(&on); err != nil {
			t.Fatal(err)
		}
		if on != 1 {
			t.Errorf("conn %d: foreign_keys = %d, want 1", i, on)
		}
	}

	log.Println("breakpoint")
}
