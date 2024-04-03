package myna

import "fmt"

const (
	BASE_URL_PROD = "https://api.starlingbank.com/api/v2"
	BASE_URL_SB   = "https://api-sandbox.starlingbank.com/api/v2"
)

const (
	ACCOUNTS_ENDPOINT = "/accounts"
	BALANCE_ENDPOINT  = "/balance"
)

var (
	ACCOUNTS_URL_PROD = fmt.Sprintf("%s%s", BASE_URL_PROD, ACCOUNTS_ENDPOINT)
)
