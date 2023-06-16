package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
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

func (h *Handlers) CreateBookmark(c *fiber.Ctx) error {
	var b models.Bookmark
	if err := c.BodyParser(&b); err != nil {
		return err
	}

	username := c.Locals("username").(string)
	if username == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid username",
		})
	}

	if err := h.store.CreateBookmark(username, &b); err != nil {
		return err
	}

	return c.JSON(b)
}

func (h *Handlers) GetBookmarks(c *fiber.Ctx) error {
	username := c.Locals("username").(string)
	if username == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid username",
		})
	}

	bookmarks, err := h.store.GetBookmarks(username)
	if err != nil {
		fmt.Println("Error retrieving bookmarks:", err)
		return c.SendString("Error retrieving bookmarks")
	}

	var bookmarkResults []fiber.Map

	for _, bookmark := range bookmarks {
		url := fmt.Sprintf("https://www.alphavantage.co/query?function=GLOBAL_QUOTE&symbol=%s&apikey=%s", bookmark.Symbol, lib.GoDotEnvVariable("API_KEY"))

		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("Error making the request:", err)
			continue
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading the response body:", err)
			continue
		}

		var data map[string]map[string]interface{}
		err = json.Unmarshal(body, &data)
		if err != nil {
			fmt.Println("Error parsing JSON:", err)
			continue
		}

		globalQuote := data["Global Quote"]

		addedPrice := bookmark.Price
		currentPrice, _ := strconv.ParseFloat(globalQuote["05. price"].(string), 64)
		profitAndLoss := currentPrice - addedPrice

		bookmarkResult := fiber.Map{
			"symbol":          bookmark.Symbol,
			"added_price":     addedPrice,
			"current_price":   currentPrice,
			"profit_and_loss": profitAndLoss,
		}

		bookmarkResults = append(bookmarkResults, bookmarkResult)
	}

	return c.JSON(bookmarkResults)
}
