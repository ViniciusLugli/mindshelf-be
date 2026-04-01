// @title Mindshelf API
// @version 1.0
// @description API documentation for Mindshelf
// @host localhost:8080
// @BasePath /api
// @schemes http https
package main

import (
	"log"
	"net/http"

	"github.com/ViniciusLugli/mindshelf/internal/handlers"
	wsHandler "github.com/ViniciusLugli/mindshelf/internal/handlers/ws"
	"github.com/ViniciusLugli/mindshelf/internal/middlewares"
	"github.com/ViniciusLugli/mindshelf/internal/repositories"
	"github.com/ViniciusLugli/mindshelf/internal/services"
	util "github.com/ViniciusLugli/mindshelf/internal/utils/ws"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	_ "github.com/ViniciusLugli/mindshelf/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

	// API

	{
		userRoute := protected.Group("/user")
		userRoute.GET("/", userHandler.GetUser)
		userRoute.GET("/all", userHandler.GetAllUsers)
		userRoute.GET("/:name", userHandler.GetAllUsersByName)
		userRoute.PATCH("/update", userHandler.Update)
		userRoute.DELETE("/delete", userHandler.Delete)
	}

	{
		groupRoute := protected.Group("/group")
		groupRoute.GET("/", groupHandler.GetAllGroups)
		groupRoute.GET("/id/:id", groupHandler.GetGroupByID)
		groupRoute.GET("/name/:name", groupHandler.GetAllGroupsByName)
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

	// WebSocket

	chatRepository := repositories.NewMessageRepository(db)
	chatService := services.NewMessageService(chatRepository)

	hub := util.NewHub()
	wsRouter := util.NewRouter()

	wsHandler.NewFriendHandlers(userService).Register(wsRouter)
	wsHandler.NewChatHandler(chatService, hub)

	// Swagger
	router.GET("/swagger", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
	})
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// WebSocket docs endpoints (for Swagger UI only)
	wsHandler.RegisterWebsocketDocs(router)

	router.Run()
}
