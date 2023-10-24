package controller

import (
	"api-e-commerce/models"
	"github.com/gofiber/fiber/v2"
)

func ListCategories(c *fiber.Ctx) error {
	categories, err := models.Categories{}.CategoriesList()
	if err != nil {
		return c.Status(400).JSON("Erro ao listar as categorias.")
	}
	return c.Status(200).JSON(categories)
}
