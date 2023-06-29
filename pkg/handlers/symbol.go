package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/kayraberktuncer/portfolion/pkg/common/lib"
	"github.com/kayraberktuncer/portfolion/pkg/common/models"
)

// CreateSymbol godoc
// @Summary Create a new symbol
// @Description Creates a new symbol
// @Tags symbols
// @Accept json
// @Produce json
// @Param body body models.Symbol true "Symbol object"
// @Success 200 {object} models.Symbol
// @Security ApiKeyAuth
// @Router /symbols [post]
func (h *Handlers) SearchSymbol(c *fiber.Ctx) error {
	symbol := c.Params("symbol")
	if symbol == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid symbol",
		})
	}

	url := fmt.Sprintf("https://www.alphavantage.co/query?function=SYMBOL_SEARCH&keywords=%s&apikey=%s", symbol, lib.GoDotEnvVariable("API_KEY"))

	resp, err := http.Get(url)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error making the request",
		})
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error reading the response body",
		})
	}

	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error parsing JSON",
		})
	}

	results, ok := data["bestMatches"].([]interface{})
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error retrieving search results",
		})
	}

	var searchResults []fiber.Map

	for _, result := range results {
		searchResult, ok := result.(map[string]interface{})
		if !ok {
			continue
		}

		searchResultMap := fiber.Map{
			"symbol":      searchResult["1. symbol"],
			"name":        searchResult["2. name"],
			"type":        searchResult["3. type"],
			"region":      searchResult["4. region"],
			"currency":    searchResult["8. currency"],
			"match_score": searchResult["9. matchScore"],
		}

		searchResults = append(searchResults, searchResultMap)
	}

	return c.JSON(searchResults)
}

const (
	symbolsPerBatch = 5
	batchInterval   = 2 * time.Minute
	retryInterval   = 1 * time.Hour
)

// UpdateSymbolValues godoc
// @Summary Update symbol values
// @Description Updates symbol values
// @Tags symbols
// @Accept json
// @Produce json
// @Success 200
// @Security ApiKeyAuth
// @Router /symbols [put]
func (h *Handlers) UpdateSymbolValues() {
	symbols, err := h.store.GetSymbols()

	if err != nil {
		fmt.Println("Error retrieving symbols:", err)
		return
	}

	numSymbols := len(symbols)
	if numSymbols == 0 {
		fmt.Println("No symbols found in the database.")
		return
	}

	batches := numSymbols / symbolsPerBatch
	lastBatchSize := numSymbols % symbolsPerBatch

	for i := 0; i < batches; i++ {
		startIndex := i * symbolsPerBatch
		endIndex := startIndex + symbolsPerBatch
		processSymbols(h, symbols[startIndex:endIndex])
	}

	if lastBatchSize > 0 {
		startIndex := batches * symbolsPerBatch
		endIndex := startIndex + lastBatchSize
		processSymbols(h, symbols[startIndex:endIndex])
	}
}

func processSymbols(h *Handlers, symbols []models.Symbol) {
	for _, symbol := range symbols {
		url := fmt.Sprintf("https://www.alphavantage.co/query?function=GLOBAL_QUOTE&symbol=%s&apikey=%s", symbol.Symbol, lib.GoDotEnvVariable("API_KEY"))

		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("Error making the request:", err)
			continue
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
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

		globalQuote, ok := data["Global Quote"]
		if !ok {
			fmt.Println("Error retrieving global quote")
			continue
		}

		currentPrice, _ := strconv.ParseFloat(globalQuote["05. price"].(string), 64)

		symbol.Price = currentPrice

		err = h.store.CreateOrUpdateSymbol(&symbol)
		if err != nil {
			fmt.Println("Error updating symbol data:", err)
			continue
		}
	}
}
