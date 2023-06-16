package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/kayraberktuncer/portfolion/pkg/common/lib"
)

func (h *Handlers) GetDataByName(c *fiber.Ctx) error {
	symbol := c.Params("symbol")
	if symbol == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid symbol",
		})
	}

	url := fmt.Sprintf("https://www.alphavantage.co/query?function=TIME_SERIES_INTRADAY&symbol=%s&interval=5min&apikey=%s", symbol, lib.GoDotEnvVariable("API_KEY"))

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making the request:", err)
		return c.SendString("Error making the request")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading the response body:", err)
		return c.SendString("Error reading the response body")
	}

	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return c.SendString("Error parsing JSON")
	}

	return c.JSON(data)
}
