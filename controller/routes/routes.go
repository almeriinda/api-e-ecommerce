package routes

import (
	"api-e-commerce/controller"
	"github.com/gofiber/fiber/v2"
)

func InitRoutes() *fiber.App {
	app := fiber.New()
	app.Post("/user", controller.CreateUser)
	app.Get("/categories", controller.ListCategories)
	app.Get("/brands", controller.ListBrand)
	app.Post("/products", controller.CreateProducts)
	app.Get("/products", controller.ListProducts)
	return app
}
