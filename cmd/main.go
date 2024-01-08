package main

import (
	"chatroom/internal/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
)


func main() {
    app := fiber.New()

    app.Use(helmet.New())

    app.Static("/", "./public") 

    routes.SetupAuthRoutes(app)

    app.Listen(":3000")
}