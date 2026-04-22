package testutil

import (
	"context"
	"database/sql"
	"testing"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func TestStartPostgres(t *testing.T) {
	ctx := context.Background()

	_, connStr := StartPostgres(ctx, t)

	db, err := sql.Open("pgx", connStr)
	if err != nil {
		t.Fatalf("open db: %s", err)
	}
	defer db.Close()

	if err := db.PingContext(ctx); err != nil {
		t.Fatalf("ping db: %s", err)
	}
}

func TestPostgresSchema(t *testing.T) {

}
