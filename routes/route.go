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
	app.Post("/api/post", controller.CreatePost)
	app.Get("/api/allpost", controller.AllPost)
	app.Get("/api/post/:id", controller.DetailPost)
	app.Put("/api/post/:id", controller.UpdatePost)
	app.Get("/api/uniquePost", controller.UniquePost)
	app.Delete("/api/deletepost/:id", controller.DeletePost)
	app.Post("/api/upload-image", controller.Upload)
	app.Static("/api/uploads", "./uploads")
}
