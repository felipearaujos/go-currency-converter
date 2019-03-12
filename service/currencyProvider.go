package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type QuoteResponse struct {
	Quotes map[string]float64 `json:"quotes"`
}

func listAllCoinsAvaliableCoinsAndCurrency() QuoteResponse {
	url := "http://www.apilayer.net/api/live?access_key=5f6b1096c7a12c66227a659e438509cc"

	response, err := http.Get(url)

	if err != nil {
		fmt.Print(err.Error())
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var quoteRespose QuoteResponse
	err = json.Unmarshal(responseData, &quoteRespose)

	return quoteRespose
}
