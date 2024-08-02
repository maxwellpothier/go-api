package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"maxpothier.com/go/api/api"
	"maxpothier.com/go/api/model"
)

func getRate(w http.ResponseWriter, r *http.Request) {
	currency := r.URL.Query().Get("currency")
	upperCurrency := strings.ToUpper(currency)
	rate, err := api.GetRate(upperCurrency)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := model.RateResponse{
		Currency: rate.Currency,
		Price:	rate.Price,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	server := http.NewServeMux()
	server.HandleFunc("/rate", getRate)

	err := http.ListenAndServe(":8080", server)
	if err != nil {
		fmt.Println("Error starting server: ", err)
	}
}
