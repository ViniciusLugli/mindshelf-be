package main

import (
	"github.com/ViniciusLugli/mindshelf/internal/repositories"
)

func main() {
	repositories.RunUpMigrations()
}
