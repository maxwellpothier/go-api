package main

import (
	"fmt"
	"sync"

	"maxpothier.com/go/api/api"
)

func main() {
	currencies := []string{"BTC", "ETH", "LTC"}
	var wg sync.WaitGroup
	for _, currency := range currencies {
		wg.Add(1)
		go func (currencyCode string) {
			getCurrencyData(currencyCode)
			wg.Done()
		}(currency)
	}
	wg.Wait()

}

func getCurrencyData(currency string) {
	rate, err := api.GetRate(currency)

	if err != nil {
		panic(err)
	}

	fmt.Println("The rate for", rate.Currency, "is", rate.Price)
}