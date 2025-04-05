package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/revacholiere-moralist/blogbackend/controller"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controller.Register)
}