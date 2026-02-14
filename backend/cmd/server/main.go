package main

import (
	"fmt"
	v1 "gin-tutorial/api/v1"
	"gin-tutorial/config"
	"gin-tutorial/internal/item"
	"gin-tutorial/internal/user"
	"gin-tutorial/pkg/logger"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// 1. Load Configuration, Logger
	cfg := config.LoadConfig()
	log := logger.New()

	log.Info("Starting application on port", cfg.Port)

	// 2. Initialize Repositories
	itemRepo := item.NewRepository()
	userRepo := user.NewRepository()

	// 3. Initialize Services (Business Logic)
	itemService := item.NewService(itemRepo)
	userService := user.NewService(userRepo)

	// 4. Initialize Handlers (API Layer)
	itemHandler := v1.NewItemHandler(itemService)
	userHandler := v1.NewUserHandler(userService)

	// 5. Setup Router
	r := gin.Default()

	// CORS Middleware
	r.Use(cors.Default())

	// Global Middleware
	r.Use(func(c *gin.Context) {
		log.Info("Request:", c.Request.Method, c.Request.URL.Path)
		c.Next()
	})

	// 6. Register Routes
	apiGroup := r.Group("/api/v1")
	itemHandler.RegisterRoutes(apiGroup)
	userHandler.RegisterRoutes(apiGroup)

	// 7. Start Server
	serverAddr := fmt.Sprintf(":%s", cfg.Port)
	if err := r.Run(serverAddr); err != nil {
		log.Error("Failed to start server:", err)
	}
}
