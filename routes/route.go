package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/revacholiere-moralist/blogbackend/controller"
	"github.com/revacholiere-moralist/blogbackend/middleware"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controller.Register)
	app.Post("/api/login", controller.Login)

	app.Use(middleware.IsAuthenticate)
}
