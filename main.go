package main

import (
	"github.com/Vakishna/WebWalApiAuth/database"
	"github.com/Vakishna/WebWalApiAuth/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.Connect()
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	routes.Setup(app)
	app.Listen(":8080")
}
