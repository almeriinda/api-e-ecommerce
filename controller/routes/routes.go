package routes

import (
	"api-e-commerce/controller"
	"github.com/gofiber/fiber/v2"
)

func InitRoutes() *fiber.App {
	app := fiber.New()
	app.Post("/user", controller.CreateUser)
	app.Get("/categories", controller.ListCategories)
	return app
}
