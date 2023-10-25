package controller

import (
	"api-e-commerce/models"
	"github.com/gofiber/fiber/v2"
)

//Cadastrar novos produtos
func CreateProducts(c *fiber.Ctx) error {
	var products models.Products
	if err := c.BodyParser(&products); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Erro ao analisar o corpo da solicitação",
		})
	}
	if err := products.CreateProducts(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Erro ao cadastrar produto",
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Produto cadastrado com sucesso",
	})
}

//Lista produtos disponivéis
func ListProducts(c *fiber.Ctx) error {
	categories, err := models.Products{}.ProductsList()
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Erro ao listar produtos."})
	}
	return c.Status(200).JSON(categories)
}
