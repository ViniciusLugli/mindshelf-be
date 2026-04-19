// @title Mindshelf API
// @version 1.0
// @description API documentation for Mindshelf
// @host localhost:8080
// @BasePath /api
// @schemes http https
package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/ViniciusLugli/mindshelf/internal/handlers"
	wsHandler "github.com/ViniciusLugli/mindshelf/internal/handlers/ws"
	"github.com/ViniciusLugli/mindshelf/internal/middlewares"
	"github.com/ViniciusLugli/mindshelf/internal/repositories"
	"github.com/ViniciusLugli/mindshelf/internal/services"
	"github.com/ViniciusLugli/mindshelf/internal/utils/logger"
	util "github.com/ViniciusLugli/mindshelf/internal/utils/ws"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	_ "github.com/ViniciusLugli/mindshelf/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	appLogger := logger.New("mindshelf-api")

	repositories.RunUpMigrations()

	err := godotenv.Load()
	if err != nil {
		appLogger.Warn("failed to load .env file", "error", err)
	}

	db, err := repositories.ConnectDB()
	if err != nil {
		appLogger.Error("failed to connect to database", "error", err)
		os.Exit(1)
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
	taskService := services.NewTaskService(taskRepository, groupRepository)
	taskHandler := handlers.NewTaskHandler(taskService)
	chatRepository := repositories.NewMessageRepository(db)
	chatService := services.NewMessageService(chatRepository, taskRepository, groupRepository)
	sharedTaskHandler := handlers.NewSharedTaskHandler(chatService)

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(middlewares.RequestID())
	router.Use(middlewares.RequestLogger(appLogger))
	router.Use(middlewares.Recovery(appLogger))

	router.POST("/register", authHandler.Register)
	router.POST("/login", authHandler.Login)

	protected := router.Group("/api")
	protected.Use(middlewares.Auth())

	// API

	{
		userRoute := protected.Group("/users")
		userRoute.GET("", userHandler.GetAllUsers)
		userRoute.GET("/me", userHandler.GetCurrentUser)
		userRoute.GET("/:id", userHandler.GetUserByID)
		userRoute.PATCH("/me", userHandler.Update)
		userRoute.DELETE("/me", userHandler.Delete)
	}

	{
		groupRoute := protected.Group("/groups")
		groupRoute.GET("", groupHandler.GetAllGroups)
		groupRoute.GET("/:id", groupHandler.GetGroupByID)
		groupRoute.POST("", groupHandler.Create)
		groupRoute.PATCH("/:id", groupHandler.Update)
		groupRoute.DELETE("/:id", groupHandler.Delete)
	}

	{
		taskRoute := protected.Group("/tasks")
		taskRoute.GET("", taskHandler.GetAllTasks)
		taskRoute.GET("/:id", taskHandler.GetTask)
		taskRoute.POST("", taskHandler.Create)
		taskRoute.PATCH("/:id", taskHandler.Update)
		taskRoute.DELETE("/:id", taskHandler.Delete)
	}

	{
		sharedTaskRoute := protected.Group("/shared-tasks")
		sharedTaskRoute.POST("/import", sharedTaskHandler.Import)
	}

	// WebSocket

	hub := util.NewHub()
	wsRouter := util.NewRouter()

	wsHandler.NewFriendHandlers(userService).Register(wsRouter)
	wsHandler.NewChatHandler(chatService, hub).Register(wsRouter)
	websocketHandler := handlers.NewWSHandler(hub, wsRouter)
	protected.GET("/ws", websocketHandler.Handle)

	// Swagger
	router.GET("/swagger", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
	})
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// WebSocket docs endpoints (for Swagger UI only)
	wsHandler.RegisterWebsocketDocs(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	appLogger.Info(
		"starting API server",
		"gin_mode", gin.Mode(),
		"port", port,
	)

	if err := router.Run(fmt.Sprintf(":%s", port)); err != nil {
		appLogger.Error("failed to run API server", "error", err)
		os.Exit(1)
	}
}
