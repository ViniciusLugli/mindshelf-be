package main

import (
	"database/sql"
	"log"
	"os"
	"strconv"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	if len(os.Args) < 2 {
		log.Fatal("use: go run ./cmd/down_migrate <number | all> <nil | force>")
	}

	dbURL := os.Getenv("DATABASE_URL")

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

	quantity := os.Args[1]

	var force string
	if len(os.Args) < 3 {
		force = ""
	} else {
		force = os.Args[2]
	}

	if force != "" {
		if quantity != "all" {
			force_steps_down(quantity, m)
			return
		}

		log.Fatal("use: go run ./cmd/down_migrate <number> force")
	}

	if quantity == "all" {
		all_down(m)
		return
	}

	steps_down(quantity, m)
}

func all_down(migrate *migrate.Migrate) {
	if err := migrate.Down(); err != nil {
		log.Fatal(err)
	}
	log.Println("All migrations rolled back successfully")
}

func steps_down(quantity string, migrate *migrate.Migrate) {
	steps, err := strconv.Atoi(quantity)
	if err != nil {
		log.Fatal("invalid number of steps")
	}

	if err := migrate.Down(); err != nil {
		log.Fatal(err)
	}
	log.Printf("Rolled back %d migrations successfully", steps)
}

func force_steps_down(quantity string, migrate *migrate.Migrate) {
	steps, err := strconv.Atoi(quantity)
	if err != nil {
		log.Fatal("invalid number of steps")
	}

	if err := migrate.Force(steps); err != nil {
		log.Fatal(err)
	}
	log.Printf("Forced rolled back %d migrations successfully", steps)
}
