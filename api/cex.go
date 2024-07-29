package api

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"maxpothier.com/go/api/model"
)

const apiURL = "https://cex.io/api/ticker/%s/USD"

func GetRate(currency string) (*model.Rate, error) {
	upCurrency := strings.ToUpper(currency)
	response, err := http.Get(fmt.Sprintf(apiURL, upCurrency))

	if err != nil {
		return nil, err
	}
	if response.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}
		jsonString := string(bodyBytes)
		fmt.Println(jsonString)
	} else {
		return nil, fmt.Errorf("API returned status code %d", response.StatusCode)
	}

	rate := model.Rate{Currency: currency, Price: 0.0}
	fmt.Println(rate)

	return &rate, nil
}
