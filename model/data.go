package model

type Rate struct {
	Currency string
	Price	float64
}

type RateResponse struct {
    Currency 	string  `json:"currency"`
    Price     	float64 `json:"price"`
}
