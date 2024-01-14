package controllers

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func Signup(c *fiber.Ctx, db *sql.DB) error  {

	log.Println("Received a request")

	form, err := c.MultipartForm()
	if err != nil {
		log.Println("Error parsing form:", err)
		return err
	}
	
	files := form.File["file-upload"]

	for _, file := range files {
		fmt.Printf("Received file: %s with size: %d bytes\n", file.Filename, file.Size)

		// Handle the file as needed, e.g., save it to disk or process it further.
	}

	return c.SendString("Signup handler")
}
