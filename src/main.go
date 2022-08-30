package main

import (
	"accounts-service/src/db"
	"accounts-service/src/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	db.InitDataBase()
	defer db.CloseDataBase()
	db.Migrations()

	app := fiber.New()
	app.Use(cors.New())

	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	app.Post("/api/v1/login", handlers.Login)
	app.Post("/api/v1/signup", handlers.Signup)

	app.Listen(":3000")
}
