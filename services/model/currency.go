package model

import "time"

type CurrencyData struct {
	Meta struct {
		LastUpdatedAt time.Time `json:"last_updated_at"`
	} `json:"meta"`
	Data map[string]CurrencyInfo `json:"data"`
}

type CurrencyResponse struct {
	Value float64 `json:"value"`
	Date  string  `json:"date"`
}

type CurrencyInfo struct {
	Code  string  `json:"code"`
	Value float64 `json:"value"`
}
