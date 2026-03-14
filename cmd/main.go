package main

import (
	"log"

	"github.com/ViniciusLugli/mindshelf/internal/repositories"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Title string `gorm:"not null"`
}

func main() {
	router := gin.Default()

	db, err := repositories.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	router.Run()
}
