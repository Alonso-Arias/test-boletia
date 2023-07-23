package model

import "time"

type Currency struct {
	Id        int       `json:"id"`
	Code      string    `json:"code"`
	Value     float64   `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}

type CallsLog struct {
	CurrencyId     int   `json:"currencyId"`
	ResponseTimeMs int64 `json:"responseTimeMs"`
}

func (CallsLog) TableName() string {
	return "calls_log"
}

func (Currency) TableName() string {
	return "currencies"
}
