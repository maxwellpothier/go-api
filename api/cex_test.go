package api_test

import (
	"testing"

	"maxpothier.com/go/api/api"
)

func TestAPICall(t *testing.T) {
	_, err := api.GetRate("")
	if err == nil {
		t.Error("Error was not returned when it should have been")
	}
}

func TestAPICallWithInvalidCurrency(t *testing.T) {
	_, err := api.GetRate("US")
	if err == nil {
		t.Error("Error was not returned when it should have been")
	}
}
