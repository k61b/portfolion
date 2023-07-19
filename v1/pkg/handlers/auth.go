package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"

	"github.com/kayraberktuncer/portfolion/pkg/common/lib"
	"github.com/kayraberktuncer/portfolion/pkg/common/models"

	log "github.com/sirupsen/logrus"
)

// Session godoc
// @Summary User session
// @Description Creates a new user session or retrieves an existing session
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body models.User true "User object"
// @Success 200 {object} models.User
// @Router /session [post]
func (h *Handlers) Session(c *fiber.Ctx) error {
	var u models.User
	if err := c.BodyParser(&u); err != nil {
		log.Error("Error parsing body:", err)
		return err
	}

	user, err := h.store.GetUserByUsername(u.Username)
	if err != nil {
		log.Error("Error retrieving user:", err)
		return err
	}

	if user == nil {
		hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			log.Error("Error hashing password:", err)
			return err
		}

		u.Password = string(hash)
		u.Bookmarks = []models.Bookmark{}
		u.Avatar = lib.GoDotEnvVariable("AVATAR_API") + u.Username + lib.GoDotEnvVariable("AVATAR_API_OPTIONS")

		if err := h.store.CreateUser(&u); err != nil {
			log.Error("Error creating user:", err)
			return err
		}

		user = &u
	} else {
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(u.Password)); err != nil {
			log.Error("Invalid username or password:", err)
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Invalid username or password",
			})
		}
	}

	token, err := lib.GenerateJWT(user.Username)
	if err != nil {
		log.Error("Error generating JWT:", err)
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
// @Tags Auth
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} models.User
// @Router /auth [get]
func (h *Handlers) Auth(c *fiber.Ctx) error {
	token := c.Cookies("token")
	if token == "" {
		log.Error("Missing token")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Missing token",
		})
	}

	username, err := lib.ParseJWT(token)
	if err != nil {
		log.Error("Invalid token:", err)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid token",
		})
	}

	user, err := h.store.GetUserByUsername(username)
	if err != nil {
		log.Error("Error retrieving user:", err)
		return err
	}

	var value float64
	var profitAndLoss float64
	for _, bookmark := range user.Bookmarks {
		symbol, err := h.store.GetSymbolValue(bookmark.Symbol)
		if err != nil {
			log.Error("Error retrieving symbol:", err)
			return err
		}

		value += symbol.Price * bookmark.Pieces
		profitAndLoss += (symbol.Price - bookmark.Price) * bookmark.Pieces
	}

	userResult := fiber.Map{
		"username":        user.Username,
		"avatar":          user.Avatar,
		"bookmarks":       user.Bookmarks,
		"value":           value,
		"profit_and_loss": profitAndLoss,
	}

	return c.JSON(userResult)
}

// Logout godoc
// @Summary User logout
// @Description Logs out the user
// @Tags Auth
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
		log.Error("Missing token")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Missing token",
		})
	}

	username, err := lib.ParseJWT(token)
	if err != nil {
		log.Error("Invalid token:", err)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid token",
		})
	}

	user, err := h.store.GetUserByUsername(username)
	if err != nil {
		log.Error("Error retrieving user:", err)
		return err
	}

	var balance float64
	for _, bookmark := range user.Bookmarks {
		symbol, err := h.store.GetSymbolValue(bookmark.Symbol)
		if err != nil {
			log.Error("Error retrieving symbol:", err)
			return err
		}

		balance += symbol.Price * bookmark.Pieces
	}

	return c.JSON(fiber.Map{
		"balance": balance,
	})
}
