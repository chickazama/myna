package models

type CurrencyFlag struct {
	Enabled  bool   `json:"enabled"`
	Currency string `json:"currency"`
}
