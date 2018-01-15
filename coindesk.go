package coindesk

import (
	"fmt"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// GetPrice fetches the latest Bitcoin price from coindesk API in given currency.
// s is the ISO code of the currency. Defaults to USD.
// Supported Currencies - https://api.coindesk.com/v1/bpi/supported-currencies.json
func GetPrice(s ...string) float64 {
	curr := "USD"
	if len(s) > 0 {
		curr = strings.ToUpper(s[0])
	}
	url := fmt.Sprintf("https://api.coindesk.com/v1/bpi/currentprice/%s.json", curr)

	response, err := http.Get(url)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	body := string(responseData)
	data := fmt.Sprintf("bpi.%s.rate_float", curr)
	price := gjson.Get(body, data).String()
	output, err := strconv.ParseFloat(price, 64)
	if err != nil {
		log.Fatal(err)
	}

	return output
}

// CurrentPrice fetches current price of Bitcoin from coindesk API.
// It returns the price in USD, GBP and EUR.
func CurrentPrice() (float64, float64, float64) {
	url := "https://api.coindesk.com/v1/bpi/currentprice.json"

	response, err := http.Get(url)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	body := string(responseData)
	priceUsd := gjson.Get(body, "bpi.USD.rate_float").String()
	USD, err := strconv.ParseFloat(priceUsd, 64)
	if err != nil {
		log.Fatal(err)
	}
	priceGbp := gjson.Get(body, "bpi.GBP.rate_float").String()
	GBP, err := strconv.ParseFloat(priceGbp, 64)
	if err != nil {
		log.Fatal(err)
	}
	priceEur := gjson.Get(body, "bpi.EUR.rate_float").String()
	EUR, err := strconv.ParseFloat(priceEur, 64)
	if err != nil {
		log.Fatal(err)
	}

	return USD, GBP, EUR
}
