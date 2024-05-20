package models

type SignedCurrencyAndAmount struct {
	Currency   string `json:"currency"`
	MinorUnits int64  `json:"minorUnits"`
}
