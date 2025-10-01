package main

import (
	"final/internal/config"
	"final/internal/database"
	"final/internal/handlers"
	"final/internal/middleware"
	"final/pkg/logger"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	// Initialize logger
	logger.Init()

	// Load configuration
	cfg := config.Load()

	// Connect to database
	if err := database.Connect(cfg); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer database.Close()

	// Create Fiber app
	app := fiber.New(fiber.Config{
		AppName: cfg.AppName,
	})

	// Middleware
	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: cfg.AllowedOrigins,
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	app.Use(middleware.Logger())

	// Routes
	setupRoutes(app)

	// Start server
	addr := fmt.Sprintf(":%s", cfg.Port)
	log.Printf("🚀 Server started at http://localhost:%s", cfg.Port)
	if err := app.Listen(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func setupRoutes(app *fiber.App) {
	// API v1 routes
	api := app.Group("/api/v1")

	// Health check
	api.Get("/health", handlers.HealthCheck)

	// Auth routes
	auth := api.Group("/auth")
	auth.Post("/signup", handlers.Signup)
	auth.Post("/login", handlers.Login)

	// Protected routes (add authentication middleware later)
	// protected := api.Group("/")
	// protected.Use(middleware.AuthRequired())
}
