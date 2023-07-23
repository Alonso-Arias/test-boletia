package model

type Currency struct {
	ID            int    `json:"id"`
	Symbol        string `json:"symbol"`
	Name          string `json:"name"`
	SymbolNative  string `json:"symbolNative"`
	DecimalDigits int    `json:"decimalDigits"`
	Rounding      int    `json:"rounding"`
	Code          string `json:"code"`
	NamePlural    string `json:"namePlural"`
}

type CallsLog struct {
	CurrencyId     int `json:"currencyId"`
	ResponseTimeMs int `json:"responseTimeMs"`
}

func (CallsLog) TableName() string {
	return "calls_log"
}

func (Currency) TableName() string {
	return "currencies"
}
