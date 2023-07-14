package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"

	"github.com/kayraberktuncer/portfolion/pkg/common/lib"
	"github.com/kayraberktuncer/portfolion/pkg/common/models"
)

const (
	symbolsPerBatch = 5
	batchInterval   = 2 * time.Minute
	retryInterval   = 1 * time.Hour
)

func (h *Handlers) UpdateSymbolValues() {
	symbols, err := h.store.GetSymbols()

	if err != nil {
		log.Warning("Error getting symbols from the database:", err)
		return
	}

	numSymbols := len(symbols)
	if numSymbols == 0 {
		log.Warning("No symbols found in the database")
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
			log.Warning("Error making the request:", err)
			continue
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Warning("Error reading the response body:", err)
			continue
		}

		var data map[string]map[string]interface{}
		err = json.Unmarshal(body, &data)
		if err != nil {
			log.Warning("Error parsing JSON:", err)
			continue
		}

		globalQuote, ok := data["Global Quote"]
		if !ok {
			log.Warning("Error retrieving global quote")
			continue
		}

		currentPrice, _ := strconv.ParseFloat(globalQuote["05. price"].(string), 64)

		symbol.Price = currentPrice

		err = h.store.CreateOrUpdateSymbol(&symbol)
		if err != nil {
			log.Warning("Error updating symbol:", err)
			continue
		}
	}
}

func (h *Handlers) CheckAndAddOrUpdateSymbol(symbol string) (*models.Symbol, error) {
	url := fmt.Sprintf("https://www.alphavantage.co/query?function=GLOBAL_QUOTE&symbol=%s&apikey=%s", symbol, lib.GoDotEnvVariable("API_KEY"))

	resp, err := http.Get(url)
	if err != nil {
		log.Error("Error making the request:", err)
		return nil, fmt.Errorf("error making the request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error("Error reading the response body:", err)
		return nil, fmt.Errorf("error reading the response body: %v", err)
	}

	var data map[string]map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Error("Error parsing JSON:", err)
		return nil, fmt.Errorf("error parsing JSON: %v", err)
	}

	globalQuote, ok := data["Global Quote"]
	if !ok {
		log.Error("Error retrieving global quote")
		return nil, fmt.Errorf("error retrieving global quote")
	}

	if len(globalQuote) == 0 {
		log.Error("Invalid symbol")
		return nil, fmt.Errorf("invalid symbol")
	}

	currentPrice, _ := strconv.ParseFloat(globalQuote["05. price"].(string), 64)
	symbolName := globalQuote["01. symbol"].(string)

	symbolData := models.Symbol{
		Symbol: symbolName,
		Price:  currentPrice,
	}

	err = h.store.CreateOrUpdateSymbol(&symbolData)
	if err != nil {
		log.Error("Error updating symbol data:", err)
		return nil, fmt.Errorf("error updating symbol data: %v", err)
	}

	return &symbolData, nil
}

// SearchSymbol godoc
// @Summary Search symbol
// @Description Search symbol
// @Tags Symbols
// @Accept  json
// @Produce  json
// @Param symbol path string true "Symbol"
// @Success 200 {object} models.Symbol
// @Router /search/{symbol} [get]
func (h *Handlers) SearchSymbol(c *fiber.Ctx) error {
	symbol := c.Params("symbol")
	if symbol == "" {
		log.Error("Invalid symbol")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid symbol",
		})
	}

	url := fmt.Sprintf("https://www.alphavantage.co/query?function=SYMBOL_SEARCH&keywords=%s&apikey=%s", symbol, lib.GoDotEnvVariable("API_KEY"))

	resp, err := http.Get(url)
	if err != nil {
		log.Error("Error making the request:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error making the request",
		})
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error("Error reading the response body:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error reading the response body",
		})
	}

	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Error("Error parsing JSON:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error parsing JSON",
		})
	}

	results, ok := data["bestMatches"].([]interface{})
	if !ok {
		log.Error("Error retrieving search results")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error retrieving search results",
		})
	}

	var searchResults []fiber.Map

	for _, result := range results {
		searchResult, ok := result.(map[string]interface{})
		if !ok {
			log.Warning("Error retrieving search result")
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
