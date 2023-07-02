package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"

	"github.com/kayraberktuncer/portfolion/pkg/common/lib"
	"github.com/kayraberktuncer/portfolion/pkg/common/models"
)

// Session godoc
// @Summary User session
// @Description Creates a new user session or retrieves an existing session
// @Tags auth
// @Accept json
// @Produce json
// @Param body body models.User true "User object"
// @Success 200 {object} models.User
// @Router /session [post]
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
		hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		u.Password = string(hash)
		u.Bookmarks = []models.Bookmark{}

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
		Expires: time.Now().Add(24 * time.Hour),
	}
	c.Cookie(&cookie)

	return c.JSON(user)
}

// Auth godoc
// @Summary User authentication
// @Description Retrieves the authenticated user
// @Tags auth
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} models.User
// @Router /auth [get]
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

	var value float64
	var profitAndLoss float64
	for _, bookmark := range user.Bookmarks {
		symbol, err := h.store.GetSymbolValue(bookmark.Symbol)
		if err != nil {
			return err
		}

		value += symbol.Price * bookmark.Pieces
		profitAndLoss += (symbol.Price - bookmark.Price) * bookmark.Pieces
	}

	userResult := fiber.Map{
		"username":        user.Username,
		"bookmarks":       user.Bookmarks,
		"value":           value,
		"profit_and_loss": profitAndLoss,
	}

	return c.JSON(userResult)
}

// Logout godoc
// @Summary User logout
// @Description Logs out the user
// @Tags auth
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} string
// @Router /logout [get]
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

func (h *Handlers) UserBalance(c *fiber.Ctx) error {
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

	var balance float64
	for _, bookmark := range user.Bookmarks {
		symbol, err := h.store.GetSymbolValue(bookmark.Symbol)
		if err != nil {
			return err
		}

		balance += symbol.Price * bookmark.Pieces
	}

	return c.JSON(fiber.Map{
		"balance": balance,
	})
}
