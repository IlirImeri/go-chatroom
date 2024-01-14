// internal/controllers/auth_controller.go

package controllers

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx, db *sql.DB) error {
	// Your login logic here
	return c.SendString("Login handler")
}
