package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/felipearaujos/go.currency.convert/models"
)

func ListAllCoinsAvaliableCoinsAndCurrency() models.QuoteResponse {
	url := "http://www.apilayer.net/api/live?access_key="

	response, err := http.Get(url)

	if err != nil {
		fmt.Print(err.Error())
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var quoteRespose models.QuoteResponse
	err = json.Unmarshal(responseData, &quoteRespose)

	return quoteRespose
}
