package main

import (
	"log"

	"github.com/ViniciusLugli/mindshelf/internal/repositories"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Title string `gorm:"not null"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	db, err := repositories.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()

	router.Run()
}
