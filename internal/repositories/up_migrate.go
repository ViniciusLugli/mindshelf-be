package repositories

import (
	"errors"
	"log"

	"github.com/ViniciusLugli/mindshelf/internal/utils/envutil"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunUpMigrations() {
	err := envutil.LoadDotEnvIfPresent()
	if err != nil {
		log.Printf("warning: failed to load .env file: %v", err)
	}

	dbURL := envutil.DatabaseDSN()
	if dbURL == "" {
		log.Fatal("DATABASE_URL or DSN is not configured")
	}

	m, err := migrate.New("file://migrations", dbURL)
	if err != nil {
		log.Fatal(err)
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatal(err)
	}

	log.Println("Migrations are up to date")
}
