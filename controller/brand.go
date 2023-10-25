package controller

import (
	"api-e-commerce/models"
	"github.com/gofiber/fiber/v2"
)

//Lista produtos disponivéis
func ListBrand(c *fiber.Ctx) error {
	brands, err := models.Brand{}.BrandsList()
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Erro ao listar marcas."})
	}
	return c.Status(200).JSON(brands)
}
