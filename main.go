package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/lpernett/godotenv"
	"github.com/revacholiere-moralist/blogbackend/database"
	"github.com/revacholiere-moralist/blogbackend/routes"
)

func main() {
	database.Connect()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	app := fiber.New()
	routes.Setup(app)
	app.Listen(":" + port)
}
