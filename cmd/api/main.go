package main

import (
	"log"

	"github.com/ViniciusLugli/mindshelf/internal/handlers"
	"github.com/ViniciusLugli/mindshelf/internal/middlewares"
	"github.com/ViniciusLugli/mindshelf/internal/repositories"
	"github.com/ViniciusLugli/mindshelf/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

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

	authService := services.NewAuthService(userRepository)
	authHandler := handlers.NewAuthHandler(authService)

	groupRepository := repositories.NewGroupRepository(db)
	groupService := services.NewGroupService(groupRepository)
	groupHandler := handlers.NewGroupHandler(groupService)

	taskRepository := repositories.NewTaskRepository(db)
	taskService := services.NewTaskService(taskRepository)
	taskHandler := handlers.NewTaskHandler(taskService)

	router := gin.Default()

	router.POST("/register", authHandler.Register)
	router.POST("/login", authHandler.Login)

	protected := router.Group("/api")
	protected.Use(middlewares.Auth())

	{
		userRoute := protected.Group("/user")
		userRoute.GET("/", userHandler.GetUser)
		userRoute.GET("/all", userHandler.GetAllUsers)
		userRoute.GET("/:name", userHandler.GetAllUsersByName)
		userRoute.PATCH("/update", userHandler.Update)
		userRoute.DELETE("/delete", userHandler.Delete)
	}

	{
		friendRoute := protected.Group("/friend")
		friendRoute.GET("/", userHandler.GetFriends)
		friendRoute.POST("/send", userHandler.SendFriendRequest)
		friendRoute.POST("/accept", userHandler.AcceptFriendRequest)
		friendRoute.POST("/reject", userHandler.RejectFriendRequest)
	}

	{
		groupRoute := protected.Group("/group")
		groupRoute.GET("/", groupHandler.GetAllGroups)
		groupRoute.GET("/:id", groupHandler.GetGroupByID)
		groupRoute.GET("/:name", groupHandler.GetAllGroupsByName)
		groupRoute.POST("/create", groupHandler.Create)
		groupRoute.PATCH("/update", groupHandler.Update)
		groupRoute.POST("/delete", groupHandler.Delete)
	}

	{
		taskRoute := protected.Group("/task")
		taskRoute.GET("/", taskHandler.GetTask)
		taskRoute.GET("/all", taskHandler.GetAllTasks)
		taskRoute.GET("/:title", taskHandler.GetAllTasksByTitle)
		taskRoute.POST("/create", taskHandler.Create)
		taskRoute.PATCH("/update", taskHandler.Update)
		taskRoute.DELETE("/delete", taskHandler.Delete)
	}

	router.Run()
}
