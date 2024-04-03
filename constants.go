package myna

import "fmt"

const (
	BASE_URL_PROD = "https://api.starlingbank.com/api/v2"
	BASE_URL_SB   = "https://api-sandbox.starlingbank.com/api/v2"
)

const (
	ACCOUNT_HOLDER_ENDPOINT        = "/account-holder"
	ACCOUNTS_ENDPOINT              = "/accounts"
	BALANCE_ENDPOINT               = "/balance"
	CONFIRMATION_OF_FUNDS_ENDPOINT = "/confirmation-of-funds"
	IDENTIFIERS_ENDPOINT           = "/identifiers"
)

var (
	ACCOUNTS_URL_PROD = fmt.Sprintf("%s%s", BASE_URL_PROD, ACCOUNTS_ENDPOINT)
)
