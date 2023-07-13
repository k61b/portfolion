package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/kayraberktuncer/portfolion/pkg/common/lib"
	"github.com/kayraberktuncer/portfolion/pkg/common/models"
)

// CreateBookmark godoc
// @Summary Create a new bookmark
// @Description Creates a new bookmark for the authenticated user
// @Tags Bookmarks
// @Accept json
// @Produce json
// @Param body body models.Bookmark true "Bookmark object"
// @Success 200 {object} models.Bookmark
// @Security ApiKeyAuth
// @Router /bookmarks [post]
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

	symbolData, err := h.CheckAndAddOrUpdateSymbol(b.Symbol)
	if err != nil {
		return err
	}

	b.Symbol = symbolData.Symbol

	if err := h.store.CreateBookmark(username, &b); err != nil {
		return err
	}

	return c.JSON(b)
}

// GetBookmarks godoc
// @Summary Get all bookmarks
// @Description Retrieves all bookmarks for the authenticated user
// @Tags Bookmarks
// @Accept json
// @Produce json
// @Success 200 {array} models.Bookmark
// @Security ApiKeyAuth
// @Router /bookmarks [get]
func (h *Handlers) GetBookmarks(c *fiber.Ctx) error {
	username := c.Locals("username").(string)
	if username == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid username",
		})
	}

	bookmarks, err := h.store.GetBookmarks(username)
	if err != nil {
		return err
	}

	var bookmarkResults []fiber.Map

	for _, bookmark := range bookmarks {
		url := fmt.Sprintf("https://www.alphavantage.co/query?function=GLOBAL_QUOTE&symbol=%s&apikey=%s", bookmark.Symbol, lib.GoDotEnvVariable("API_KEY"))

		symbolData, err := h.store.GetSymbolValue(bookmark.Symbol)
		if err != nil && err != mongo.ErrNoDocuments {
			continue
		}

		var currentPrice float64
		if symbolData != nil && symbolData.Price != 0 {
			currentPrice = symbolData.Price
		} else {
			resp, err := http.Get(url)
			if err != nil {
				continue
			}
			defer resp.Body.Close()

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				continue
			}

			var data map[string]map[string]interface{}
			err = json.Unmarshal(body, &data)
			if err != nil {
				continue
			}

			globalQuote := data["Global Quote"]

			currentPrice, _ = strconv.ParseFloat(globalQuote["05. price"].(string), 64)

			if symbolData != nil {
				symbolData.Price = currentPrice
				err := h.store.CreateOrUpdateSymbol(symbolData)
				if err != nil {
					continue
				}
			} else {
				newSymbol := &models.Symbol{
					Symbol: bookmark.Symbol,
					Price:  currentPrice,
				}

				err := h.store.CreateOrUpdateSymbol(newSymbol)
				if err != nil {
					continue
				}
			}
		}

		addedPrice := bookmark.Price
		pieces := bookmark.Pieces
		profitAndLoss := (currentPrice - addedPrice) * pieces

		bookmarkResult := fiber.Map{
			"symbol":          bookmark.Symbol,
			"added_price":     addedPrice,
			"current_price":   currentPrice,
			"pieces":          pieces,
			"profit_and_loss": profitAndLoss,
		}

		bookmarkResults = append(bookmarkResults, bookmarkResult)
	}

	return c.JSON(bookmarkResults)
}

// UpdateBookmark godoc
// @Summary Update a bookmark
// @Description Updates a bookmark for the authenticated user
// @Tags Bookmarks
// @Accept json
// @Produce json
// @Param symbol path string true "Symbol"
// @Param body body models.Bookmark true "Bookmark object"
// @Success 200 {object} models.Bookmark
// @Security ApiKeyAuth
// @Router /bookmarks/{symbol} [put]
func (h *Handlers) UpdateBookmark(c *fiber.Ctx) error {
	symbol := c.Params("symbol")
	if symbol == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid symbol",
		})
	}

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

	if err := h.store.UpdateBookmark(username, symbol, &b); err != nil {
		return err
	}

	return c.JSON(b)
}

// DeleteBookmark godoc
// @Summary Delete a bookmark
// @Description Deletes a bookmark for the authenticated user
// @Tags Bookmarks
// @Accept json
// @Produce json
// @Param symbol path string true "Symbol"
// @Success 200 {object} string
// @Security ApiKeyAuth
// @Router /bookmarks/{symbol} [delete]
func (h *Handlers) DeleteBookmark(c *fiber.Ctx) error {
	symbol := c.Params("symbol")
	if symbol == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid symbol",
		})
	}

	username := c.Locals("username").(string)
	if username == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid username",
		})
	}

	if err := h.store.DeleteBookmark(username, symbol); err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "success",
	})
}
