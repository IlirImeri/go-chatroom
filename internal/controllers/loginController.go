// internal/controllers/auth_controller.go

package controllers

import "github.com/gofiber/fiber/v2"

func Login(c *fiber.Ctx) error {
	// Your login logic here
	return c.SendString("Login handler")
}
