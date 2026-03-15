package main

import (
	"log"

	"github.com/ViniciusLugli/mindshelf/internal/handlers"
	"github.com/ViniciusLugli/mindshelf/internal/repositories"
	"github.com/ViniciusLugli/mindshelf/internal/services"
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

	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)

	router := gin.Default()

	{
		userRoute := router.Group("/user")
		userRoute.GET("/", userHandler.GetUser)
		userRoute.GET("/all", userHandler.GetAllUsers)
		userRoute.POST("/create", userHandler.Create)
		userRoute.PATCH("/update", userHandler.Update)
		userRoute.DELETE("/delete", userHandler.Delete)
	}

	router.Run()
}
