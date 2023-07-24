package model

import "time"

type Currency struct {
	Id        int       `json:"id"`
	Currency  string    `json:"currency"`
	Value     float64   `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}

type CallsLog struct {
	CallTimestamp  time.Time `json:"callTimestamp"`
	ResponseTimeMs int64     `json:"responseTimeMs"`
	Status         string    `json:"status"`
}

func (CallsLog) TableName() string {
	return "calls_log"
}

func (Currency) TableName() string {
	return "currencies"
}
