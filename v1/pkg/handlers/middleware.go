package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kayraberktuncer/portfolion/pkg/common/lib"
	log "github.com/sirupsen/logrus"
)

func (h *Handlers) AuthMiddleware(c *fiber.Ctx) error {
	token := c.Cookies("token")
	if token == "" {
		log.Error("No token found")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid token",
		})
	}

	username, err := lib.ParseJWT(token)
	if err != nil {
		log.Error("Error parsing JWT:", err)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid token",
		})
	}

	c.Locals("username", username)

	return c.Next()
}
