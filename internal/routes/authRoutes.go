// internal/routes/auth_routes.go

package routes

import (
	"chatroom/internal/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupAuthRoutes(app *fiber.App) {
	auth := app.Group("/auth")

	auth.Post("/login", controllers.Login)

	auth.Post("/signup", controllers.Signup)
}
