package main

import (
	"accounts-service/src/auth"
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

	app.Put("/api/v1/settings/update/username", auth.IsLogin, handlers.UpdateUsername)
	app.Put("/api/v1/settings/update/password", auth.IsLogin, handlers.UpdatePassword)

	app.Listen(":3000")
}
