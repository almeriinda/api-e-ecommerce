package controller

import (
	"api-e-commerce/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func SignIn(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Erro ao analisar o corpo da solicitação",
		})
	}

	if user.Email == "" || user.PasswordHash == "" {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": "Email e senha são obrigatórios",
		})
	}

	existingUser, err := models.GetUserByEmail(user.Email)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Erro ao verificar as credenciais do usuário",
		})
	}

	if existingUser.ID > 0 && bcrypt.CompareHashAndPassword([]byte(existingUser.PasswordHash), []byte(user.PasswordHash)) == nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Login bem-sucedido",
			"user_id": existingUser.ID,
		})
	}

	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"message": "Email ou senha inválidos",
	})
}

func UpdateUser(c *fiber.Ctx) error {
	userID := c.Params("id")

	var updatedUser models.User
	if err := c.BodyParser(&updatedUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Erro ao analisar o corpo da solicitação",
		})
	}

	err := models.UpdateUserByID(userID, &updatedUser)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Erro ao atualizar informações do usuário",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Informações do usuário atualizadas com sucesso",
	})
}
