package main

import (
	"chatroom/internal/routes"
	"fmt"

	"github.com/joho/godotenv"

	"log"
	"os"

	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	_ "github.com/lib/pq"
)


func main() {

    err := godotenv.Load("../.env")

    if err != nil {
        log.Fatal("Error loading .env file")
    }

    psqlconn := fmt.Sprintf(
        "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", 
        os.Getenv("HOST"), os.Getenv("PORT"), os.Getenv("USER"), os.Getenv("PASSWORD"), os.Getenv("DBNAME"))

    db, err := sql.Open("postgres", psqlconn)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close() // Ensure the database connection is closed when the function exits

    app := fiber.New()

    app.Use(helmet.New())

    app.Static("/", "../public") 

    routes.SetupAuthRoutes(app, db)

    app.Listen(":3000")
}