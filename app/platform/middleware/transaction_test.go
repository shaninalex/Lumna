package middleware_test

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"testing"

	"gitlab.com/shaninalex/lumna/app/core/bus"
	"gitlab.com/shaninalex/lumna/app/platform/config"
	"gitlab.com/shaninalex/lumna/app/platform/database"
	pmw "gitlab.com/shaninalex/lumna/app/platform/middleware"
)

func openDB(t *testing.T) *database.DB {
	t.Helper()

	dir := t.TempDir()
	cfgPath := filepath.Join(dir, "config.yaml")
	body := "env: testing\ndatabase:\n  type: \"sqlite\"\n  url: \"" + filepath.Join(dir, "test.db") + "\"\n"
	if err := os.WriteFile(cfgPath, []byte(body), 0o600); err != nil {
		t.Fatal(err)
	}

	db := database.New(config.ReadConfig(cfgPath))
	if err := db.From(context.Background()).Exec("CREATE TABLE notes (id integer PRIMARY KEY AUTOINCREMENT, body text)").Error; err != nil {
		t.Fatal(err)
	}
	return db
}

func countNotes(t *testing.T, db *database.DB) int64 {
	t.Helper()

	var n int64
	if err := db.From(context.Background()).Table("notes").Count(&n).Error; err != nil {
		t.Fatal(err)
	}
	return n
}

type writeTwice struct{ fail bool }

// newBus wires a command bus that writes two rows through DB.From(ctx) — the
// same way a repository does — and optionally fails after the second one.
func newBus(t *testing.T, db *database.DB) *bus.CommandBus {
	t.Helper()

	b := bus.NewCommandBus(pmw.Transaction(db))
	err := bus.RegisterCommand(b, func(ctx context.Context, cmd writeTwice) (int, error) {
		for _, body := range []string{"first", "second"} {
			if err := db.From(ctx).Exec("INSERT INTO notes (body) VALUES (?)", body).Error; err != nil {
				return 0, err
			}
		}
		if cmd.fail {
			return 0, errors.New("handler failed after writing")
		}
		return 2, nil
	})
	if err != nil {
		t.Fatal(err)
	}
	return b
}

func TestTransactionCommits(t *testing.T) {
	db := openDB(t)

	written, err := bus.Execute[writeTwice, int](context.Background(), newBus(t, db), writeTwice{})
	if err != nil {
		t.Fatalf("execute: %v", err)
	}
	if written != 2 {
		t.Fatalf("handler returned %d, want 2", written)
	}
	if got := countNotes(t, db); got != 2 {
		t.Fatalf("rows after commit = %d, want 2", got)
	}
}

func TestTransactionRollsBackEveryWriteOfTheCommand(t *testing.T) {
	db := openDB(t)

	_, err := bus.Execute[writeTwice, int](context.Background(), newBus(t, db), writeTwice{fail: true})
	if err == nil {
		t.Fatal("expected the handler error to reach the caller")
	}
	// The point of the middleware: a failure halfway leaves nothing behind,
	// even though the handler wrote through two separate statements.
	if got := countNotes(t, db); got != 0 {
		t.Fatalf("rows after rollback = %d, want 0", got)
	}
}

func TestFromOutsideCommandAutocommits(t *testing.T) {
	db := openDB(t)
	ctx := context.Background()

	if err := db.From(ctx).Exec("INSERT INTO notes (body) VALUES (?)", "loose").Error; err != nil {
		t.Fatal(err)
	}
	if got := countNotes(t, db); got != 1 {
		t.Fatalf("rows = %d, want 1", got)
	}
}

func TestNestedTransactionIsRejected(t *testing.T) {
	db := openDB(t)

	_, err := db.InTx(context.Background(), func(ctx context.Context) (any, error) {
		return db.InTx(ctx, func(context.Context) (any, error) { return nil, nil })
	})
	if err == nil {
		t.Fatal("expected nested transaction to be rejected")
	}
}
