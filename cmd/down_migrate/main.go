package main

import (
	"database/sql"
	"errors"
	"log"
	"os"
	"strconv"

	"github.com/ViniciusLugli/mindshelf/internal/utils/envutil"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {
	err := envutil.LoadDotEnvIfPresent()
	if err != nil {
		log.Printf("warning: failed to load .env file: %v", err)
	}

	if len(os.Args) < 2 {
		log.Fatal("use: go run ./cmd/down_migrate <number | all | wipe | force> [args]")
	}

	dbURL := envutil.DatabaseDSN()
	if dbURL == "" {
		log.Fatal("DATABASE_URL or DSN is not configured")
	}

	arg := os.Args[1]

	if arg == "wipe" {
		if len(os.Args) < 3 || (os.Args[2] != "yes" && os.Args[2] != "force") {
			log.Fatal("This will DROP the public schema and erase all data. Confirm by running: go run ./cmd/down_migrate wipe yes")
		}

		db, err := sql.Open("postgres", dbURL)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		_, err = db.Exec("DROP SCHEMA public CASCADE; CREATE SCHEMA public;")
		if err != nil {
			log.Fatal(err)
		}

		log.Println("Database schema 'public' dropped and recreated successfully")
		return
	}

	m, err := migrate.New("file://migrations", dbURL)
	if err != nil {
		log.Fatal(err)
	}

	if arg == "force" {
		if len(os.Args) < 3 {
			log.Fatal("use: go run ./cmd/down_migrate force <version>")
		}

		forceVersion(os.Args[2], m)
		return
	}

	quantity := os.Args[1]

	if quantity == "all" {
		all_down(m)
		return
	}

	steps_down(quantity, m)
}

func all_down(m *migrate.Migrate) {
	if err := m.Down(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatal(err)
	}
	log.Println("All migrations rolled back successfully")
}

func steps_down(quantity string, m *migrate.Migrate) {
	steps, err := strconv.Atoi(quantity)
	if err != nil || steps <= 0 {
		log.Fatal("invalid number of steps")
	}

	if err := m.Steps(-steps); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatal(err)
	}
	log.Printf("Rolled back %d migrations successfully", steps)
}

func forceVersion(version string, m *migrate.Migrate) {
	forceToVersion, err := strconv.Atoi(version)
	if err != nil || forceToVersion < 0 {
		log.Fatal("invalid migration version")
	}

	if err := m.Force(forceToVersion); err != nil {
		log.Fatal(err)
	}
	log.Printf("Forced migration state to version %d successfully", forceToVersion)
}
