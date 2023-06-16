package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"

	"github.com/kayraberktuncer/portfolion/pkg/common/lib"
	"github.com/kayraberktuncer/portfolion/pkg/common/models"
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

func (h *Handlers) Session(c *fiber.Ctx) error {
	var u models.User
	if err := c.BodyParser(&u); err != nil {
		return err
	}

	user, err := h.store.GetUserByUsername(u.Username)
	if err != nil {
		return err
	}

	if user == nil {
		hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
		if err != nil {
			return err
		}

		u.Password = string(hash)

		if err := h.store.CreateUser(&u); err != nil {
			return err
		}

		user = &u
	} else {
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(u.Password)); err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Invalid username or password",
			})
		}
	}

	token, err := lib.GenerateJWT(user.Username)
	if err != nil {
		return err
	}

	cookie := fiber.Cookie{
		Name:    "token",
		Value:   token,
		Path:    "/",
		Expires: time.Now().Add(time.Hour * 24),
	}
	c.Cookie(&cookie)

	return c.JSON(user)
}

func (h *Handlers) Auth(c *fiber.Ctx) error {
	token := c.Cookies("token")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Missing token",
		})
	}

	username, err := lib.ParseJWT(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid token",
		})
	}

	user, err := h.store.GetUserByUsername(username)
	if err != nil {
		return err
	}

	return c.JSON(user)
}

func (h *Handlers) Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:    "token",
		Value:   "",
		Path:    "/",
		Expires: time.Now().Add(-time.Hour),
	}
	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}
