package models

type CurrencyAndAmount struct {
	Currency   string `json:"currency"`
	MinorUnits int64  `json:"minorUnits"`
}
