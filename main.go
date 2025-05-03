package main

import (
	"fmt"
	"log"
	"restapi_go/config"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	cfg := config.LoadConfig()

	// Set up Fiber app

	app := fiber.New(fiber.Config{
		Prefork:           true,
		CaseSensitive:     true,
		StrictRouting:     true,
		ServerHeader:      "Fiber",
		AppName:           cfg.ApiName,
		BodyLimit:         3 * 1024 * 1024, // 3 MB
		RequestMethods:    []string{"GET", "HEAD", "POST", "PUT", "DELETE", "CONNECT", "OPTIONS", "TRACE", "PATCH"},
		EnablePrintRoutes: true,
	})

	// Set up CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE,PATCH,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	// Fiber logger formatter (On screen logger only)
	app.Use(logger.New(logger.Config{
		Format:     "${time} | ${method} | ${status} | ${path} | ${ip} | ${latency}\n",
		TimeFormat: "2006-01-02 15:04:05.000",
	}))

	// Set up routes
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "OK",
			"message": "Server is running",
			"date":    time.Now().Format("2006-01-02 15:04:05"),
		})
	})

	// Start the server
	address := fmt.Sprintf(":%d", cfg.ApiPort)
	log.Printf("Starting server on %s", address)
	log.Fatal(app.Listen(address))

}
