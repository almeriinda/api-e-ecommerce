package controller

import (
	"api-e-commerce/models"
	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Erro ao analisar o corpo da solicitação",
		})
	}
	if len(user.PasswordHash) < 5 || len(user.PasswordHash) > 10 {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": "Tamanho da senha inválido (deve ter entre 5 e 10 caracteres)",
		})
	}
	if err := user.CreateUser(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Erro ao criar o usuário",
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Usuário criado com sucesso",
	})
}
