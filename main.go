package main

import (
	"log"

	"github.com/arifinreinaldo/mine12/config"
	"github.com/arifinreinaldo/mine12/database"
	"github.com/arifinreinaldo/mine12/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Connect to database
	if err := database.Connect(cfg); err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Run migrations
	if err := database.Migrate(); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Initialize Fiber app
	app := fiber.New(fiber.Config{
		AppName: "Sales API v1.0.0",
	})

	// Middleware
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// Health check
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Sales API is running",
			"version": "1.0.0",
		})
	})

	// API routes
	api := app.Group("/api/v1")

	// User routes
	users := api.Group("/users")
	users.Get("/", handlers.GetAllUsers)
	users.Get("/:id", handlers.GetUserByID)
	users.Post("/", handlers.CreateUser)
	users.Put("/:id", handlers.UpdateUser)
	users.Delete("/:id", handlers.DeleteUser)

	// Item routes
	items := api.Group("/items")
	items.Get("/", handlers.GetAllItems)
	items.Get("/:id", handlers.GetItemByID)
	items.Post("/", handlers.CreateItem)
	items.Put("/:id", handlers.UpdateItem)
	items.Delete("/:id", handlers.DeleteItem)

	// Start server
	log.Printf("Server starting on port %s", cfg.AppPort)
	log.Fatal(app.Listen(":" + cfg.AppPort))
}
