package main

import (
	"log"
	"os"
	"strconv"

	"github.com/golang-migrate/migrate/v4"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	if len(os.Args) < 2 {
		log.Fatal("use: go run ./cmd/down_migrate <number | all>")
	}

	dbURL := os.Getenv("DATABASE_URL")

	m, err := migrate.New("file://migrations", dbURL)
	if err != nil {
		log.Fatal(err)
	}

	arg := os.Args[1]

	if arg == "all" {
		if err := m.Down(); err != nil {
			log.Fatal(err)
		}
		log.Println("All migrations rolled back successfully")
		return
	}

	steps, err := strconv.Atoi(arg)
	if err != nil {
		log.Fatal("invalid number of steps")
	}

	if err := m.Down(); err != nil {
		log.Fatal(err)
	}
	log.Printf("Rolled back %d migrations successfully", steps)

}
