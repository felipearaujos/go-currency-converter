package main

import (
	"fmt"

	"log"
	"net/http"

	"github.com/felipearaujos/go.currency.convert/repository"
	"github.com/labstack/echo"
)

func MakeHandlers() {
	e := echo.New()
	e.GET("/", healthCheck)
	e.GET("/quotes", listAllCoinsAvaliableCoins)
	e.POST("/convert", convertCurrency)

	log.Fatal(e.Start(":8081"))
	fmt.Println("Listening:8081")
}

func healthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, "OK")
}

type QuoteConverterRequest struct {
	From QuoteConverterValueRequest `json:"from"`
	To   QuoteConverterValueRequest `json:"to"`
}
type QuoteConverterValueRequest struct {
	Value    float64 `json:"Value"`
	Currency string  `json:"Currency"`
}

func convertCurrency(c echo.Context) error {
	var converter QuoteConverterRequest
	err := c.Bind(&converter)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	currencyAndCoinsResponse := repository.ListAllCoinsAvaliableCoinsAndCurrency()

	fromCurrency := currencyAndCoinsResponse.Quotes[converter.From.Currency]
	toCurrency := currencyAndCoinsResponse.Quotes[converter.To.Currency]

	dolarFrom := converter.From.Value / fromCurrency
	convertedValue := dolarFrom * toCurrency

	return c.JSON(http.StatusOK, convertedValue)

}

func listAllCoinsAvaliableCoins(c echo.Context) error {
	currencyAndCoinsResponse := repository.ListAllCoinsAvaliableCoinsAndCurrency()

	keys := make([]string, 0, len(currencyAndCoinsResponse.Quotes))
	for k := range currencyAndCoinsResponse.Quotes {
		//keys = append(keys, strings.Replace(k, "USD", "", -1))
		keys = append(keys, k)
	}

	return c.JSON(http.StatusOK, keys)
}

func main() {
	MakeHandlers()
}
