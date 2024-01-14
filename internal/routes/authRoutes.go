// internal/routes/auth_routes.go

package routes

import (
	"chatroom/internal/controllers"
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

func SetupAuthRoutes(app *fiber.App, db *sql.DB) {
	auth := app.Group("/auth")

	// Pass the *sql.DB reference to the controllers
	auth.Post("/login", func(c *fiber.Ctx) error {
		return controllers.Login(c, db)
	})

	auth.Post("/signup", func(c *fiber.Ctx) error {
		return controllers.Signup(c, db)
	})
}
