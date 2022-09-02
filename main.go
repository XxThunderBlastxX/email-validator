package main

import (
	"github.com/XxThunderBlastxX/email-validator/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	//Loads variables from .env
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	//Init New Fiber App
	app := fiber.New()

	//Enable CORS
	app.Use(cors.New())

	//Application Route
	routes.Router(app)

	//Listen Application at desired port from .env
	log.Fatal(app.Listen(os.Getenv("PORT")))
}
