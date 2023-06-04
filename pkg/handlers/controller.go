package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/kayraberktuncer/portfolion/pkg/common/models"
)

func (h *Handlers) GetHome(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func (h *Handlers) GetUsers(c *fiber.Ctx) error {
	cursor, err := h.store.GetUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	var users []models.User
	if err := cursor.All(c.Context(), &users); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(users)
}
