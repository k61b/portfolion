package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kayraberktuncer/portfolion/pkg/common/lib"
)

func (h *Handlers) AuthMiddleware(c *fiber.Ctx) error {
	token := c.Cookies("token")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid token",
		})
	}

	username, err := lib.ParseJWT(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid token",
		})
	}

	c.Locals("username", username)

	return c.Next()
}
