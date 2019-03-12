package main

import (
	"fmt"

	"log"
	"net/http"
	"strings"

	"github.com/felipearaujos/go.currency.convert/service"
	"github.com/labstack/echo"
)

func MakeHandlers() {
	e := echo.New()
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/", healthCheck)
	e.GET("/quotes", listAllCoinsAvaliableCoins)

	log.Fatal(e.Start(":8081"))
	fmt.Println("Listening:8081")
}

func healthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, "OK")
}

func listAllCoinsAvaliableCoins(c echo.Context) error {
	currencyAndCoinsResponse := services.ListAllCoinsAvaliableCoinsAndCurrency()

	keys := make([]string, 0, len(currencyAndCoinsResponse.Quotes))
	for k := range currencyAndCoinsResponse.Quotes {
		keys = append(keys, strings.Replace(k, "USD", "", -1))
	}

	return c.JSON(http.StatusOK, keys)
}

func main() {
	MakeHandlers()
}
