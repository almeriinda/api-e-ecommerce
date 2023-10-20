package controller

import (
	"api-e-commerce/models"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {
	body := c.Body()
	user := models.User{}
	json.Unmarshal(body, &user)
	if len(user.PasswordHash) < 5 || len(user.PasswordHash) > 10 {
		return c.Status(401).JSON("Tamanho da senha inválido")
	}
	user.CreateUser()
	return c.Status(201).JSON("Usuário criado.")

}
