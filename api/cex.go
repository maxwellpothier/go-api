package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"maxpothier.com/go/api/model"
)

const apiURL = "https://cex.io/api/ticker/%s/USD"

func GetRate(currency string) (*model.Rate, error) {
	if len(currency) != 3 {
		return nil, fmt.Errorf("currency code must be 3 characters, recieved %s", currency)
	}
	upCurrency := strings.ToUpper(currency)
	res, err := http.Get(fmt.Sprintf(apiURL, upCurrency))
	var response CexResponse

	if err != nil {
		return nil, err
	}
	if res.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		
		err = json.Unmarshal(bodyBytes, &response)
		if err != nil {
			return nil, err
		}


	} else {
		return nil, fmt.Errorf("API returned status code %d", res.StatusCode)
	}

	rate := model.Rate{Currency: currency, Price: response.Ask}

	return &rate, nil
}
