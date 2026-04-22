package testutil

import (
	"context"
	"log"
	"path/filepath" // <--- قم بإضافة هذا السطر هنا

	_ "github.com/lib/pq"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
)

func startPostgres() {
	ctx := context.Background()

	dbName := "termperature-db"
	dbUser := "user"
	dbPassword := "password"

	postgresContainer, err := postgres.Run(ctx,
		"postgres:16-alpine",
		postgres.WithInitScripts(filepath.Join("testdata", "init-user-db.sh")),

		postgres.WithDatabase(dbName),
		postgres.WithUsername(dbUser),
		postgres.WithPassword(dbPassword),
		postgres.BasicWaitStrategies(),
	)
	defer func() {
		if err := testcontainers.TerminateContainer(postgresContainer); err != nil {
			log.Printf("failed to terminate container: %s", err)
		}
	}()
	if err != nil {
		log.Printf("failed to start container: %s", err)
		return
	}
}
